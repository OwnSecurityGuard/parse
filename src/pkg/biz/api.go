package biz

import (
	"github.com/labstack/gommon/log"
	"parse/src/pkg/monitor"
	"time"
)

func (a *App) GetColumns() []*Column {
	// todo 自定义Column
	return tmp
}
func (a *App) Pull() (dat []*Msg) {
	data := a.out.Pull(a.ctx, 3*time.Second)
	var paths []string
	for _, t := range tmp {
		//m := &Msg{map[string][]string{}}
		paths = append(paths, t.ExtractorPath)
		//m.KeyValues[t.Prop] = nil
	}

	for _, d := range data {
		jd, _ := d.ToJSON()
		msg := &Msg{IsClient: d.IsClient(), KeyValues: map[string][]string{}}
		values := monitor.ExtractorPath(string(jd), paths)
		for i := range values {
			msg.KeyValues[tmp[i].Prop] = values[i]
		}
		dat = append(dat, msg)
	}

	return
}

// todo 增加返回值，有可能失败
func (a *App) Start(port int) {
	go func() { // todo 同一时间只能开启一个监控，需要加互斥
		if err := a.fm.Run("", port); err != nil {
			log.Errorf(err.Error())
		}
	}()

}

func (a *App) Start1() *monitor.Rule {
	return nil
}

func (a *App) Start3() *monitor.Filter {
	return nil
}
