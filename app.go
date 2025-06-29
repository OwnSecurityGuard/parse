package main

// App struct

type Playback struct {
	Req  []*Tmp
	Resp []*Tmp
}
type Tmp struct {
	Timestamp int64
	Uri       string
	Msg       string
}

type Param struct {
	Index   int
	SrcPath string
	TarPath string
	TarUri  string
}

// Greet returns a greeting for the given name
//func (a *App) Greet(name string) string {
//	return fmt.Sprintf("Hello %s, It's show time!", name)
//}
//func (a *App) A() *Playback {
//
//	return &Playback{
//		Req: []*Tmp{
//			&Tmp{
//				Timestamp: time.Now().Unix(),
//				Uri:       "AAAA",
//				Msg:       `{"uri" :"AAAA", "DADA": "SDASDASDAS"}`,
//			},
//			{
//				Timestamp: time.Now().Unix(),
//				Uri:       "AAAA2",
//				Msg:       `{"uri" :"AAAA", "DADA": "SDASDASDAS"}`,
//			},
//		},
//		Resp: []*Tmp{
//			&Tmp{
//				Timestamp: time.Now().Unix() - 5,
//				Uri:       "AAAA",
//				Msg:       `{"uri" :"AAAA", "DADA": "SDASDASDAS"}`,
//			},
//		},
//	}
//}
