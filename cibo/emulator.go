package cibo

type Emulator struct {
	CPU *CPU
	RAM []byte
	baseAddress int
	// TODO: add device
}

func NewEmulator(beginAddress int, memSize int64) *Emulator {
	ram := make([]byte, memSize)
	emu := Emulator{nil, ram, beginAddress}
	emu.CPU = NewCPU(&emu)
	reg := &emu.CPU.X86registers
	reg.EIP = uint32(beginAddress)
	return &emu
}
