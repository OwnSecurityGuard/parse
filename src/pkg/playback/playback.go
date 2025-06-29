package playback

//type Playback struct {
//	ctx     context.Context
//	M       monitor.Monitor
//	C       monitor.Codec
//	Outputs []monitor.Output
//}
//
//func NewPlayback(m monitor.Monitor, c monitor.Codec, outputs []monitor.Output) *Playback {
//	a := &Playback{
//		M:       m,
//		C:       c,
//		Outputs: outputs,
//	}
//
//	return a
//}
//
//func (p *Playback) AddOutput(o ...monitor.Output) {
//	p.Outputs = append(p.Outputs, o...)
//}
//func (p *Playback) Start(ctx context.Context, filterBpf string, port int) error {
//	p.ctx = ctx
//
//	c2s, s2c, err := p.M.Start(p.ctx, filterBpf, port)
//	if err != nil {
//		return err
//	}
//
//	c2sMsg := p.C.ParseC2S(ctx, c2s)
//	s2cMsg := p.C.ParseS2C(ctx, s2c)
//
//	for _, o := range p.Outputs {
//		o.Write(ctx, c2sMsg, s2cMsg)
//	}
//	return nil
//}
