package monitor

import (
	"context"
	"log"
	"sync"
)

// Direction 枚举由 codec/output 模块共享
type Direction int

const (
	ClientToServer Direction = iota
	ServerToClient
)

type Message interface {
	IsClient() bool
	TimestampMillis() int64
	Raw() []byte
	ToJSON() ([]byte, error)
	FromJSON([]byte) error
}

type Monitor interface {
	Start(ctx context.Context, filterBpf string, port int) (c2s chan []byte, s2c chan []byte, err error)
}

type Codec[T Message] interface {
	Parse(ctx context.Context, input <-chan []byte, direction Direction) <-chan T
}

type Output[T Message] interface {
	Write(ctx context.Context, msg T, dir Direction) error
	Close() error
}

type FlowManager[T Message] struct {
	ctx context.Context
	m   Monitor
	c   Codec[T]
	os  []Output[T]
}

// 构造函数
func NewFlowManager[T Message](ctx context.Context, m Monitor, c Codec[T], outputs ...Output[T]) *FlowManager[T] {
	return &FlowManager[T]{
		ctx: ctx,
		m:   m,
		c:   c,
		os:  outputs,
	}
}

func (f *FlowManager[T]) Run(filter string, port int) error {
	c2sRaw, s2cRaw, err := f.m.Start(f.ctx, filter, port)
	if err != nil {
		return err
	}

	c2sParsed := f.c.Parse(f.ctx, c2sRaw, ClientToServer)
	s2cParsed := f.c.Parse(f.ctx, s2cRaw, ServerToClient)

	var wg sync.WaitGroup
	wg.Add(2)

	go f.dispatch(c2sParsed, ClientToServer, &wg)
	go f.dispatch(s2cParsed, ServerToClient, &wg)

	wg.Wait()

	return nil
}

// 分发消息到所有 Output
func (f *FlowManager[T]) dispatch(in <-chan T, dir Direction, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-f.ctx.Done():
			return
		case msg, ok := <-in:
			if !ok {
				return
			}
			for _, o := range f.os {
				if err := o.Write(f.ctx, msg, dir); err != nil {
					log.Printf("output write error: %v", err)
				}
			}
		}
	}
}
