package cibo

import "fmt"

type CPU struct {
  Memory
  X86registers
  X64registers
}

func NewCPU(console *Console) *CPU {
	cpu := CPU{Memory: NewCPUMemory(console)}
  cpu.Write(0, 1)
  cpu.Write(1, 2)
  cpu.Write(2, 3)
  val := cpu.Read(0, 3)
  fmt.Print(val)
	return &cpu
}
