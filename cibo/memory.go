package cibo

type Memory interface {
	Read(address uint64, index uint64) []byte
	Write(address uint64, value byte)
}

type cpuMemory struct {
	emulator *Emulator
}

func (mem *cpuMemory) Read(address uint64, index uint64) []byte {
	return mem.emulator.RAM[address:index]
}

func (mem *cpuMemory) Write(address uint64, value byte) {
	mem.emulator.RAM[address] = value
}

func NewCPUMemory(emulator *Emulator) Memory {
	return &cpuMemory{emulator}
}
