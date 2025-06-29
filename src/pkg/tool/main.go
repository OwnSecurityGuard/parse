package main

import (
	"encoding/json"
	"fmt"
	"os"
	"parse/src/pkg/monitor"
	"time"
)

type A struct {
	m monitor.Monitor
	//c monitor.Codec[monitor.BaseMessage]
	//o monitor.Output[monitor.BaseMessage]
}

func main() {
	f, _ := os.Create("a.txt")
	defer f.Close()
	a := []byte("白金卡SD卡收不到卡视角你打算、")
	data := monitor.NewBaseMessage(a, true)
	yy := monitor.GenericCodec[*monitor.BaseMessage]{
		Factory: monitor.NewBaseMessage,
	}

	data.Name = "巴萨大"
	data.Time = time.Now().UnixMilli()
	fmt.Println(data)
	da, err := json.Marshal(data)
	fmt.Println(err)
	fmt.Println(string(da))

	aw := &monitor.BaseMessage{}
	fmt.Println(json.Unmarshal(da, aw))
	fmt.Println(aw)
	//a := &A{}
	//ctx := context.Background()
	//c2s, s2c, _ := a.m.Start(ctx, "", 1)
	//
	//c2sCh := a.c.ParseC2S(ctx, c2s)
	//s2cCh := a.c.ParseS2C(ctx, s2c)
	//a.o.Write(ctx, c2sCh, s2cCh)

}
