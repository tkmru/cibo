package cibo

type Memory interface {
	Read(address uint32) byte
	Read32(address uint32) uint32
	Write(address uint32, value byte)
	Write32(address uint32, value uint32)
	GetCode8(index int) uint8
	GetSignCode8(index int) int8
	GetCode32(index int) uint32
	GetSignCode32(index int) int32
}

type cpuMemory struct {
	emulator *Emulator
}

func NewCPUMemory(emulator *Emulator) Memory {
	return &cpuMemory{emulator}
}

func (mem *cpuMemory) Read(address uint32) byte {
	return mem.emulator.RAM[address]
}

func (mem *cpuMemory) Read32(address uint32) uint32 {
	var ret uint32
	for i := 0; i < 4; i++ {
		ret |= uint32(mem.Read(address+uint32(i)) << (8 * uint32(i)))
	}
	return ret
}

func (mem *cpuMemory) Write(address uint32, value byte) {
	mem.emulator.RAM[address] = value
}

func (mem *cpuMemory) Write32(address uint32, value uint32) {
	for i := 0; i < 4; i++ {
		mem.Write(address+uint32(i), byte(value>>(uint(i)*8)))
	}
}

func (mem *cpuMemory) GetCode8(index int) uint8 {
	emu := mem.emulator
	cpu := emu.CPU
	reg := &cpu.X86registers
	address := reg.EIP - uint32(emu.BaseAddress)
	return uint8(mem.Read(address + uint32(index)))
}

func (mem *cpuMemory) GetSignCode8(index int) int8 {
	emu := mem.emulator
	cpu := emu.CPU
	reg := &cpu.X86registers
	address := reg.EIP - uint32(emu.BaseAddress)
	return int8(mem.Read(address + uint32(index)))
}

func (mem *cpuMemory) GetCode32(index int) uint32 {
	var i uint
	var ret uint32
	for i = 0; i < 4; i++ {
		ret |= uint32(mem.GetCode8(index+int(i))) << (i * 8)
	}
	return ret
}

func (mem *cpuMemory) GetSignCode32(index int) int32 {
	return int32(mem.GetCode32(index))
}
