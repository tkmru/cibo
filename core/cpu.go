package cibo

type CPU struct {
	Memory
	X86registers
	//X64registers
	Instr16 [256]func()
	Instr32 [256]func()
}

func NewCPU(emulator *Emulator) *CPU {
	cpu := CPU{Memory: NewCPUMemory(emulator)}
	cpu.createTable32()
	return &cpu
}
