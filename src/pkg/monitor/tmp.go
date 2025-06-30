package monitor

type Rule struct {
	SrcPath string
	DstKey  string
	DstPath string
	Fs      []*Filter
}
type Filter struct {
	Path      string
	ExpectKey string
	ExpectVal string
	Op        string
}
