package monitor

import (
	"fmt"
	"github.com/labstack/gommon/log"
	"sync"
)

type Playback struct {
	extractUtil *ExtractUtil
}

type CtxStore[T Message] struct {
	lock  sync.Mutex
	index int
	C2S   []T

	Resp map[string][]string
	cc   *CtxConfig

	extractUtil ExtractUtil
}
type CtxConfig struct {
	KeyPath string // 提取消息的key来作为唯一值 例如 { "seq":1, "name":"loginReq", "data":{ "s":2}}, 若为name, 则使用name来进行分组
	// 所有规则
	KeyRules map[string][]Rule // key对应的规则,如果没有配置,则默认按原有的数据进行回放即可
}

func NewCtxStore[T Message](path string, cc *CtxConfig) *CtxStore[T] {
	c := NewJSONCoder[T](path)
	for {
		msg, err := c.Decode()
		if err != nil {
			log.Error("decode err ")
			break
		}

		jsonData, err := msg.ToJSON()
		if err != nil {
			log.Error("to json err")
			break
		}
		fmt.Println(jsonData)
	}

	return &CtxStore[T]{

		index:       0,
		C2S:         nil,
		Resp:        make(map[string][]string),
		cc:          nil,
		extractUtil: nil,
	}
}

func (cs *CtxStore[T]) AddResp(a T) {
	cs.lock.Lock()
	defer cs.lock.Unlock()

	data, err := a.ToJSON()
	if err != nil {
		log.Error(err)
	}

	keys := cs.extractUtil.GetValueByBytes(data, cs.cc.KeyPath)
	if len(keys) != 1 {
		log.Errorf("keys length != 1: %v", keys)
	}
	cs.Resp[keys[0]] = append(cs.Resp[keys[0]], string(data))
}

func (cs *CtxStore[T]) Push() (t T, b bool) {
	// 理论上这里不会有并发
	defer func() {
		cs.index++
	}()
	if cs.index >= len(cs.C2S) {
		return t, false
	}
	t = cs.C2S[cs.index]

	jsonData, err := t.ToJSON()
	if err != nil {
		log.Error(err)
		return t, false
	}
	data := cs.Deal(jsonData)
	if len(data) == 0 {
		log.Errorf("data is empty: %v", data)
		return t, false
	}

	if err = t.FromJSON(data); err != nil {
		log.Error("maral json err ", err)
		return t, false
	}

	return t, true
}

// Todo 理论上这个方法出现任何为空的情况，均会导致协议数据异常,  故基本就代表回放失败了！！！
func (cs *CtxStore[T]) Deal(jsonData []byte) []byte {
	cs.lock.Lock()
	defer cs.lock.Unlock() // 主要是对Resp的操作

	key := cs.extractUtil.GetValueByBytes(jsonData, cs.cc.KeyPath)
	var (
		rules       []Rule
		ok          bool
		tmpJsonData = string(jsonData)
		err         error
	)
	if rules, ok = cs.cc.KeyRules[key[0]]; !ok {
		return jsonData
	}
	// 一个req 的json data  如果需要动态替换参数,则进行处理，处理后返回处理后的jsonData

	for _, rule := range rules {
		if datas, ok1 := cs.Resp[rule.DstKey]; ok1 {
			filterData := cs.extractUtil.Filter(tmpJsonData, datas, rule.Fs)
			if len(filterData) > 0 {
				vals := cs.extractUtil.GetValue(filterData[0], rule.DstPath)

				tmpJsonData, err = cs.extractUtil.SetValue(tmpJsonData, rule.SrcPath, vals[0])
				if err != nil {
					log.Error(err)
					return nil
				}

			} else {
				// todo
				log.Error("filter error")
			}

		} else {
			log.Error("key not found", rule.DstKey)

		}
	}
	return []byte(tmpJsonData)
}

func NewCtxConfig() *CtxConfig {
	return &CtxConfig{}
}

func (cc *CtxConfig) Dt(c2s []Message) {

	cc.KeyRules = make(map[string][]Rule)

}
