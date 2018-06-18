package cibo

type Console struct {
	CPU *CPU
	RAM []byte
	// TODO: add device
}

func NewConsole() (*Console, error) {
	ram := make([]byte, 2048)
	console := Console{nil, ram}
	console.CPU = NewCPU(&console)
	return &console, nil
}
