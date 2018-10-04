package cibo

func (cpu *CPU) movRM8R8() {
	var modrm ModRM
	reg := &cpu.X86registers
	reg.EIP += 1
	modrm.parse(cpu)
	r8 := modrm.getR8(cpu)
	modrm.setRM8(cpu, r8)
}

func (cpu *CPU) movR8RM8() {
	var modrm ModRM
	reg := &cpu.X86registers
	reg.EIP += 1
	modrm.parse(cpu)
	rm8 := modrm.getRM8(cpu)
	modrm.setR8(cpu, rm8)
}

func (cpu *CPU) movRM16R16() {
	var modrm ModRM
	reg := &cpu.X86registers
	reg.EIP += 1
	modrm.parse(cpu)
	r16 := modrm.getR16(cpu)
	modrm.setRM16(cpu, r16)
}

func (cpu *CPU) movR16RM16() {
	var modrm ModRM
	reg := &cpu.X86registers
	reg.EIP += 1
	modrm.parse(cpu)
	rm16 := modrm.getRM16(cpu)
	modrm.setR16(cpu, rm16)
}

func (cpu *CPU) movRM32R32() {
	var modrm ModRM
	reg := &cpu.X86registers
	reg.EIP += 1
	modrm.parse(cpu)
	r32 := modrm.getR32(cpu)
	modrm.setRM32(cpu, r32)
}

func (cpu *CPU) movR32RM32() {
	var modrm ModRM
	reg := &cpu.X86registers
	reg.EIP += 1
	modrm.parse(cpu)
	rm32 := modrm.getRM32(cpu)
	modrm.setR32(cpu, rm32)
}

func (cpu *CPU) movR8Imm8() {
	mem := cpu.Memory
	regIndex := mem.GetCode8(0) - 0xb0
	imm8 := mem.GetCode8(1)
	reg := &cpu.X86registers
	reg.Set8ByIndex(regIndex, imm8)
	reg.EIP += 2
}

func (cpu *CPU) movR16Imm16() {
	mem := cpu.Memory
	regIndex := mem.GetCode8(0) - 0xb8
	imm16 := mem.GetCode16(1)
	reg := &cpu.X86registers
	reg.Set16ByIndex(regIndex, imm16)
	reg.EIP += 3
}

func (cpu *CPU) movRM16Imm16() {
	var modrm ModRM
	reg := &cpu.X86registers
	mem := cpu.Memory
	reg.EIP += 1
	modrm.parse(cpu)
	imm16 := mem.GetCode16(0)
	reg.EIP += 2
	modrm.setRM16(cpu, imm16)
}

func (cpu *CPU) movR32Imm32() {
	mem := cpu.Memory
	regIndex := mem.GetCode8(0) - 0xb8
	imm32 := mem.GetCode32(1)
	reg := &cpu.X86registers
	reg.SetByIndex(regIndex, imm32)
	reg.EIP += 5
}

func (cpu *CPU) movRM32Imm32() {
	var modrm ModRM
	reg := &cpu.X86registers
	mem := cpu.Memory
	reg.EIP += 1
	modrm.parse(cpu)
	imm32 := mem.GetCode32(0)
	reg.EIP += 4
	modrm.setRM32(cpu, imm32)
}

func (cpu *CPU) nop() {
	reg := &cpu.X86registers
	reg.EIP += 1
}
