package cibo

type Memory struct {
	RAM []byte
}

func (m *Memory) Read(address uint32) byte {
}

func (m *Memory) Write(address uint32, value byte) {
}
