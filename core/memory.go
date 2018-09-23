package cibo

// import "fmt"

type Memory interface {
	Read(address uint32) byte
	Read8(address uint32) uint8
	Read16(address uint32) uint16
	Read32(address uint32) uint32
	Write(address uint32, value byte)
	Write8(address uint32, value uint8)
	Write16(address uint32, value uint16)
	Write32(address uint32, value uint32)
	GetCode8(index int) uint8
	GetSignCode8(index int) int8
	GetCode32(index int) uint32
	GetSignCode32(index int) int32
	Push(value uint32)
	Pop() uint32
}

type cpuMemory struct {
	emulator *Emulator
}

func NewCPUMemory(emulator *Emulator) Memory {
	return &cpuMemory{emulator}
}

func (mem *cpuMemory) Read(address uint32) byte {
	emu := mem.emulator
	index := address - uint32(emu.baseAddress)
	return mem.emulator.RAM[index]
}

func (mem *cpuMemory) Read8(address uint32) uint8 {
	return uint8(mem.Read(address))
}

func (mem *cpuMemory) Read16(address uint32) uint16 {
	var ret uint16
	for i := 0; i < 2; i++ {
		ret |= uint16(mem.Read(address+uint32(i))) << (8 * uint32(i))
	}
	return ret
}

func (mem *cpuMemory) Read32(address uint32) uint32 {
	var ret uint32
	for i := 0; i < 4; i++ {
		ret |= uint32(mem.Read(address+uint32(i))) << (8 * uint32(i))
	}
	return ret
}

func (mem *cpuMemory) Write(address uint32, value byte) {
	emu := mem.emulator
	index := address - uint32(emu.baseAddress)
	mem.emulator.RAM[index] = value
}

func (mem *cpuMemory) Write8(address uint32, value uint8) {
	mem.Write(address, byte(value))
}

func (mem *cpuMemory) Write16(address uint32, value uint16) {
	for i := 0; i < 2; i++ {
		mem.Write(address+uint32(i), byte(value>>(uint(i)*8)))
	}
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
	return uint8(mem.Read(reg.EIP + uint32(index)))
}

func (mem *cpuMemory) GetSignCode8(index int) int8 {
	emu := mem.emulator
	cpu := emu.CPU
	reg := &cpu.X86registers
	return int8(mem.Read(reg.EIP + uint32(index)))
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

func (mem *cpuMemory) Push(value uint32) {
	emu := mem.emulator
	cpu := emu.CPU
	reg := &cpu.X86registers
	address := reg.ESP - 4
	reg.ESP = address
	mem.Write32(address, value)
}

func (mem *cpuMemory) Pop() (ret uint32) {
	emu := mem.emulator
	cpu := emu.CPU
	reg := &cpu.X86registers
	address := reg.ESP
	value := mem.Read32(address)
	reg.ESP = address + 4
	return value
}
