package cibo

import "fmt"

type CPU struct {
  Memory
  Registers
}

func NewCPU(console *Console) *CPU {
	cpu := CPU{Memory: NewCPUMemory(console)}
  cpu.Write(0, 1)
  val := cpu.Read(0)
  fmt.Print(val)
	return &cpu
}
