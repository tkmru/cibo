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
	value := mem.GetCode8(1)
	reg := &cpu.X86registers
	reg.Set8ByIndex(regIndex, value)
	reg.EIP += 2
}

func (cpu *CPU) movR32Imm32() {
	mem := cpu.Memory
	regIndex := mem.GetCode8(0) - 0xb8
	value := mem.GetCode32(1)
	reg := &cpu.X86registers
	reg.SetByIndex(regIndex, value)
	reg.EIP += 5
}

func (cpu *CPU) movRM32Imm32() {
	var modrm ModRM
	reg := &cpu.X86registers
	mem := cpu.Memory
	reg.EIP += 1
	modrm.parse(cpu)
	value := mem.GetCode32(0)
	reg.EIP += 4
	modrm.setRM32(cpu, value)
}

func (cpu *CPU) nop() {
	reg := &cpu.X86registers
	reg.EIP += 1
}
