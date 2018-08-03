package cibo

func (cpu *CPU) createTable() {
	cpu.InstTable[0x01] = cpu.addRM32R32
	cpu.InstTable[0x05] = cpu.addEAXImm32
	cpu.InstTable[0x3b] = cpu.cmpR32RM32
	cpu.InstTable[0x3c] = cpu.cmpALImm8
	cpu.InstTable[0x3d] = cpu.cmpEAXImm32

	for i := 0; i < 8; i++ {
		cpu.InstTable[0x50+i] = cpu.pushReg
	}

	for i := 0; i < 8; i++ {
		cpu.InstTable[0x58+i] = cpu.popReg
	}

	cpu.InstTable[0x68] = cpu.pushImm32
	cpu.InstTable[0x6a] = cpu.pushImm8
	cpu.InstTable[0x70] = cpu.jo
	cpu.InstTable[0x71] = cpu.jno
	cpu.InstTable[0x72] = cpu.jc
	cpu.InstTable[0x73] = cpu.jnc
	cpu.InstTable[0x74] = cpu.jz
	cpu.InstTable[0x75] = cpu.jnz
	cpu.InstTable[0x78] = cpu.js
	cpu.InstTable[0x79] = cpu.jns
	cpu.InstTable[0x7c] = cpu.jl
	cpu.InstTable[0x7e] = cpu.jle
	cpu.InstTable[0x81] = cpu.code81
	cpu.InstTable[0x83] = cpu.code83
	cpu.InstTable[0x89] = cpu.movRM32R32
	cpu.InstTable[0x8b] = cpu.movR32RM32
	cpu.InstTable[0x90] = cpu.nop

	for i := 0; i < 8; i++ {
		cpu.InstTable[0xb8+i] = cpu.movR32Imm32
	}

	cpu.InstTable[0xc3] = cpu.ret
	cpu.InstTable[0xc7] = cpu.movRM32Imm32
	cpu.InstTable[0xc9] = cpu.leave
	cpu.InstTable[0xe8] = cpu.callRelative
	cpu.InstTable[0xe9] = cpu.nearJump
	cpu.InstTable[0xeb] = cpu.shortJump
	cpu.InstTable[0xff] = cpu.codeFF
}

func (cpu *CPU) pushReg() {
	reg := &cpu.X86registers
	mem := cpu.Memory
	regIndex := mem.GetCode8(0) - 0x50
	mem.Push(reg.GetByIndex(regIndex))
	reg.EIP += 1
}

func (cpu *CPU) pushImm32() {
	reg := &cpu.X86registers
	mem := cpu.Memory
	value := mem.GetCode32(1)
	mem.Push(value)
	reg.EIP += 5
}

func (cpu *CPU) pushImm8() {
	reg := &cpu.X86registers
	mem := cpu.Memory
	value := mem.GetCode8(1)
	mem.Push(uint32(value))
	reg.EIP += 2
}

func (cpu *CPU) popReg() {
	reg := &cpu.X86registers
	mem := cpu.Memory
	regIndex := mem.GetCode8(0) - 0x58
	reg.SetByIndex(regIndex, mem.Pop())
	reg.EIP += 1
}

func (cpu *CPU) ret() {
	reg := &cpu.X86registers
	mem := cpu.Memory
	reg.EIP = mem.Pop()
}

func (cpu *CPU) leave() {
	reg := &cpu.X86registers
	mem := cpu.Memory
	ebp := reg.EBP
	reg.ESP = ebp
	reg.EBP = mem.Pop()
	reg.EIP += 1
}

func (cpu *CPU) callRelative() {
	reg := &cpu.X86registers
	mem := cpu.Memory
	diff := mem.GetSignCode32(1)
	mem.Push(reg.EIP + 5)
	reg.EIP += uint32(diff + 5)
}
