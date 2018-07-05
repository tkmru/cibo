package cibo

type Emulator struct {
	CPU *CPU
	RAM []byte
	beginAddress int
	// TODO: add device
}

func NewEmulator(beginAddress int, memSize int64) *Emulator {
	ram := make([]byte, memSize)
	emu := Emulator{nil, ram, beginAddress}
	emu.CPU = NewCPU(&emu)
	return &emu
}
