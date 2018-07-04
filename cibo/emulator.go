package cibo

type Emulator struct {
	CPU *CPU
	RAM []byte
	// TODO: add device
}

func NewEmulator() *Emulator {
	ram := make([]byte, 2048)
	emu := Emulator{nil, ram}
	emu.CPU = NewCPU(&emu)
	return &emu
}
