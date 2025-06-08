package monitor

type Monitor struct {
	Port string
}

func NewMonitor(port string) *Monitor {
	return &Monitor{
		Port: port,
	}
}

func (m *Monitor) Start() {

}
