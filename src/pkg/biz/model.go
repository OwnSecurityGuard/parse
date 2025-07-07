package biz

import (
	"context"
	"parse/src/pkg/monitor"
	tmp1 "parse/src/pkg/tmp"
	"time"
)

type Column struct {
	Prop          string `json:"prop"`
	Label         string `json:"label"`
	ExtractorPath string `json:"extractor_path"`
}

type Msg struct {
	//Prop   string   `json:"prop"`
	//Label  string   `json:"label"`
	//Values []string `json:"values"` // 表达式可能有需要提取成数组，一个包可能也有多条协议？
	IsClient  bool                `json:"is_client"`
	KeyValues map[string][]string `json:"key_values"`
}

type Filter struct {
	SkipKeyValue []*Skip // 当path的值命中Values的任意一个均跳过
}

type Skip struct {
	Path   string
	Values map[string]bool
}

type FrontOutput[T monitor.Message] struct {
	ch chan T
}

func NewFrontOutput[T monitor.Message]() *FrontOutput[T] {
	return &FrontOutput[T]{
		ch: make(chan T, 100),
	}
}
func (f *FrontOutput[t]) Write(ctx context.Context, msg t, dir tmp1.Direction) error {
	f.ch <- msg
	return nil
}

func (f *FrontOutput[t]) Close() error {
	//TODO implement me
	panic("implement me")
}

func (f *FrontOutput[t]) Pull(ctx context.Context, timeout time.Duration) (data []t) {
	tick := time.Tick(timeout)
	for {
		select {
		case <-tick:
			return data
		case <-ctx.Done():
			return data
		case msg := <-f.ch:
			data = append(data, msg)
		}
	}
}

var (
	tmp = []*Column{{
		Prop:          "timestamp",
		Label:         "时间",
		ExtractorPath: "a",
	}, {
		Prop:          "name",
		Label:         "协议名",
		ExtractorPath: "b",
	}, {
		Prop:          "data",
		Label:         "数据",
		ExtractorPath: "a",
	}}
)
