package cibo

type Memory interface {
	Read(address uint64, index uint64) []byte
	Write(address uint64, value byte)
	GetCode8(index int) uint8
	GetSignCode8(index int) int8
	GetCode32(index int) uint32
	GetSignCode32(index int) int32
	WriteMemory8(address uint32, value uint32)
	WriteMemory32(address uint32, value uint32)
	ReadMemory8(address uint32) uint32
	ReadMemory32(address uint32) uint32
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

func (mem *cpuMemory) GetCode8(index int) uint8 {
	emu := mem.emulator
	cpu := emu.CPU
	reg := &cpu.X86registers
	address := reg.EIP - uint32(emu.baseAddress)
	return uint8(mem.emulator.RAM[address+uint32(index)])
}

func (mem *cpuMemory) GetSignCode8(index int) int8 {
	emu := mem.emulator
	cpu := emu.CPU
	reg := &cpu.X86registers
	address := reg.EIP - uint32(emu.baseAddress)
	return int8(mem.emulator.RAM[address+uint32(index)])
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

func (mem *cpuMemory) WriteMemory8(address uint32, value uint32) {
	mem.emulator.RAM[address] = byte(value & 0xff)
}

func (mem *cpuMemory) WriteMemory32(address uint32, value uint32) {
	for i := 0; i < 4; i++ {
		mem.WriteMemory8(address+uint32(i), value>>(uint(i)*8))
	}
}

func (mem *cpuMemory) ReadMemory8(address uint32) uint32 {
	return uint32(mem.emulator.RAM[address])
}

func (mem *cpuMemory) ReadMemory32(address uint32) uint32 {
	var ret uint32
	for i := 0; i < 4; i++ {
		ret |= mem.ReadMemory8(address+uint32(i)) << (8 * uint32(i))
	}
	return ret
}

func NewCPUMemory(emulator *Emulator) Memory {
	return &cpuMemory{emulator}
}
