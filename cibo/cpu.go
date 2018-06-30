package cibo

//import "fmt"

type CPU struct {
	Memory
	X86registers
	X64registers
}

func NewCPU(emulator *Emulator) *CPU {
	cpu := CPU{Memory: NewCPUMemory(emulator)}
	//cpu.Write(0, 1)
	//cpu.Write(1, 2)
	//cpu.Write(2, 3)
	//val := cpu.Read(0, 3)
	//fmt.Print(val)
	return &cpu
}
