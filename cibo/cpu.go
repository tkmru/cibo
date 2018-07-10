package cibo

type CPU struct {
	Memory
	X86registers
	//X64registers
	InstTable [256]func()
}

func NewCPU(emulator *Emulator) *CPU {
	cpu := CPU{Memory: NewCPUMemory(emulator)}
	cpu.createTable()
	return &cpu
}
