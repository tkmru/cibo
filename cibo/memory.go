package cibo

type Memory interface {
	Read(address uint64) byte
	Write(address uint64, value byte)
}

type cpuMemory struct {
	console *Console
}

func (mem *cpuMemory) Read(address uint64) byte {
	return mem.console.RAM[address]
}

func (mem *cpuMemory) Write(address uint64, value byte) {
	mem.console.RAM[address] = value
}

func NewCPUMemory(console *Console) Memory {
	return &cpuMemory{console}
}
