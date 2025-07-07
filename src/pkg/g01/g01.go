package g01

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	"log"
	"parse/src/pkg/tmp"
	"time"
)

type TcpMonitor struct {
	c2s chan []byte
	s2c chan []byte
	ctx context.Context
}

func (t *TcpMonitor) Start(ctx context.Context, filterBpf string, port int) (c2s chan []byte, s2c chan []byte, err error) {
	t.ctx = ctx
	t.c2s, t.s2c = make(chan []byte, 100), make(chan []byte, 100)
	devs, _ := pcap.FindAllDevs()

	for _, dev := range devs {
		go func(device string, port string) {
			t.monitor(device, "udp", port)
		}(dev.Name, fmt.Sprintf("%d", port))
	}

	t.MockData() // todo 测试使用
	return t.c2s, t.s2c, nil
}
func (t *TcpMonitor) Stop() {
	t.ctx.Done()
}

func (t *TcpMonitor) monitor(device string, filterBpf string, port string) {

	handle, err := pcap.OpenLive(device, 1600, true, pcap.BlockForever)
	if err != nil {
		log.Panic(err)
	}
	err = handle.SetBPFFilter(fmt.Sprintf("%s and port %s", filterBpf, port))
	if err != nil {
		log.Panic(err)
		return
	}
	defer handle.Close()
	packets := gopacket.NewPacketSource(handle, handle.LinkType()).Packets()
	for {
		select {
		case packet := <-packets:

			//isClient := packet.TransportLayer().TransportFlow().Dst().String() == port
			data := packet.Layer(layers.LayerTypeUDP).LayerPayload()
			if true {
				fmt.Println(data)
				t.c2s <- data
			} else {
				t.s2c <- data
			}
		case <-t.ctx.Done():
			return
		}
	}

}

type Req struct {
	A int64  `json:"a"`
	B string `json:"b"`
}

func (t *TcpMonitor) MockData() {
	go func() {
		for {
			a := &Req{
				A: time.Now().UnixMilli(),
				B: fmt.Sprintf("req %d", time.Now().UnixMilli()),
			}
			data, _ := json.Marshal(a)
			t.c2s <- data
			time.Sleep(1 * time.Second)
		}

	}()
	go func() {
		for {
			a := &Req{
				A: time.Now().UnixMilli(),
				B: fmt.Sprintf("resp %d", time.Now().UnixMilli()),
			}
			data, _ := json.Marshal(a)
			t.s2c <- data
			time.Sleep(1 * time.Second)
		}

	}()

}

type G01Codec struct {
}

func (g G01Codec) Parse(ctx context.Context, input <-chan []byte, direction tmp.Direction) <-chan G01Msg {

	if direction == tmp.ClientToServer {
		return g.ParseC2S(ctx, input, direction)
	} else {
		return g.ParseS2C(ctx, input, direction)
	}

}
func (g G01Codec) ParseC2S(ctx context.Context, input <-chan []byte, direction tmp.Direction) <-chan G01Msg {
	ch := make(chan G01Msg, 100)
	go func() {
		for d := range input {
			msg := NewG01Msg() //
			msg.RawData = d
			msg.IsC = direction == tmp.ClientToServer
			ch <- msg
		} // todo ctx

	}()
	return ch
}

func (g G01Codec) ParseS2C(ctx context.Context, input <-chan []byte, direction tmp.Direction) <-chan G01Msg {
	ch := make(chan G01Msg, 100)
	go func() {
		for d := range input {
			msg := NewG01Msg()
			msg.RawData = d
			msg.IsC = direction == tmp.ClientToServer
			ch <- msg
		} // todo ctx

	}()
	return ch
}

type G01Msg struct {
	RawData []byte
	IsC     bool
	time    int64
}

func NewG01Msg() G01Msg {
	return G01Msg{}
}
func (g G01Msg) IsClient() bool {
	return g.IsC
}

func (g G01Msg) TimestampMillis() int64 {
	return g.time
}

func (g G01Msg) Raw() []byte {
	return g.RawData
}

func (g G01Msg) ToJSON() ([]byte, error) {
	return json.Marshal(g)
	//return g.RawData, nil
}

func (g G01Msg) FromJSON(bytes []byte) error {
	return json.Unmarshal(bytes, &g)

}
