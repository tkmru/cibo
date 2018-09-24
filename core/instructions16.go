package cibo

func (cpu *CPU) createTable16() {
	cpu.Instr16[0x00] = cpu.addRM8R8
	cpu.Instr16[0x01] = cpu.addRM16R16
	cpu.Instr16[0x02] = cpu.addR8RM8
	cpu.Instr16[0x03] = cpu.addR16RM16
	cpu.Instr16[0x04] = cpu.addALImm8
	cpu.Instr16[0x05] = cpu.addAXImm16
	cpu.Instr16[0x06] = cpu.pushES
	cpu.Instr16[0x07] = cpu.popES
	cpu.Instr16[0x16] = cpu.pushSS
	cpu.Instr16[0x17] = cpu.popSS
	cpu.Instr16[0x1e] = cpu.pushDS
	cpu.Instr16[0x1f] = cpu.popDS
}

func (cpu *CPU) addRM8R8() {
	reg := &cpu.X86registers
	reg.EIP += 1
	var modrm ModRM
	modrm.parse(cpu)
	r8 := modrm.getR8(cpu)
	rm8 := modrm.getRM8(cpu)
	modrm.setRM8(cpu, rm8+r8)
}

func (cpu *CPU) addRM16R16() {
	reg := &cpu.X86registers
	reg.EIP += 1
	var modrm ModRM
	modrm.parse(cpu)
	r16 := modrm.getR16(cpu)
	rm16 := modrm.getRM16(cpu)
	modrm.setRM16(cpu, rm16+r16)
}

func (cpu *CPU) addR8RM8() {
	reg := &cpu.X86registers
	reg.EIP += 1
	var modrm ModRM
	modrm.parse(cpu)
	r8 := modrm.getR8(cpu)
	rm8 := modrm.getRM8(cpu)
	modrm.setR8(cpu, rm8+r8)
}

func (cpu *CPU) addR16RM16() {
	reg := &cpu.X86registers
	reg.EIP += 1
	var modrm ModRM
	modrm.parse(cpu)
	r16 := modrm.getR16(cpu)
	rm16 := modrm.getRM16(cpu)
	modrm.setR16(cpu, rm16+r16)
}

func (cpu *CPU) addALImm8() {
	reg := &cpu.X86registers
	mem := cpu.Memory
	value := mem.GetCode8(1)
	reg.EAX += uint32(value)
	reg.EIP += 2
}

func (cpu *CPU) addAXImm16() {
	reg := &cpu.X86registers
	mem := cpu.Memory
	value := mem.GetCode16(1)
	reg.EAX += uint32(value)
	reg.EIP += 3
}

func (cpu *CPU) pushES() {
	reg := &cpu.X86registers
	mem := cpu.Memory
	mem.Push16(reg.ES)
	reg.EIP += 1
}

func (cpu *CPU) popES() {
	reg := &cpu.X86registers
	mem := cpu.Memory
	reg.ES = mem.Pop16()
	reg.EIP += 1
}

func (cpu *CPU) pushSS() {
	reg := &cpu.X86registers
	mem := cpu.Memory
	mem.Push16(reg.SS)
	reg.EIP += 1
}

func (cpu *CPU) popSS() {
	reg := &cpu.X86registers
	mem := cpu.Memory
	reg.SS = mem.Pop16()
	reg.EIP += 1
}

func (cpu *CPU) pushDS() {
	reg := &cpu.X86registers
	mem := cpu.Memory
	mem.Push16(reg.DS)
	reg.EIP += 1
}

func (cpu *CPU) popDS() {
	reg := &cpu.X86registers
	mem := cpu.Memory
	reg.DS = mem.Pop16()
	reg.EIP += 1
}
