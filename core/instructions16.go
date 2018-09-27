package cibo

import "log"

func (cpu *CPU) createTable16() {
	cpu.Instr16[0x00] = cpu.addRM8R8
	cpu.Instr16[0x01] = cpu.addRM16R16
	cpu.Instr16[0x02] = cpu.addR8RM8
	cpu.Instr16[0x03] = cpu.addR16RM16
	cpu.Instr16[0x04] = cpu.addALImm8
	cpu.Instr16[0x05] = cpu.addAXImm16
	cpu.Instr16[0x06] = cpu.push16ES
	cpu.Instr16[0x07] = cpu.pop16ES
	cpu.Instr16[0x08] = cpu.orRM8R8
	cpu.Instr16[0x09] = cpu.orRM16R16
	cpu.Instr16[0x0a] = cpu.orR8RM8
	cpu.Instr16[0x0b] = cpu.orR16RM16
	cpu.Instr16[0x0c] = cpu.orALImm8
	cpu.Instr16[0x0d] = cpu.orAXImm16
	cpu.Instr16[0x16] = cpu.push16SS
	cpu.Instr16[0x17] = cpu.pop16SS
	cpu.Instr16[0x1e] = cpu.push16DS
	cpu.Instr16[0x1f] = cpu.pop16DS

	for i := 0; i < 8; i++ {
		cpu.Instr16[0x50+i] = cpu.push16Reg
	}

	for i := 0; i < 8; i++ {
		cpu.Instr16[0x58+i] = cpu.pop16Reg
	}

	cpu.Instr16[0x66] = cpu.overrideOperandTo32
	cpu.Instr16[0x68] = cpu.push16Imm16
	cpu.Instr16[0x6a] = cpu.push16Imm8
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

func (cpu *CPU) orRM8R8() {
	reg := &cpu.X86registers
	reg.EIP += 1
	var modrm ModRM
	modrm.parse(cpu)
	r8 := modrm.getR8(cpu)
	rm8 := modrm.getRM8(cpu)
	modrm.setRM8(cpu, (rm8 | r8))
}

func (cpu *CPU) orRM16R16() {
	reg := &cpu.X86registers
	reg.EIP += 1
	var modrm ModRM
	modrm.parse(cpu)
	r16 := modrm.getR16(cpu)
	rm16 := modrm.getRM16(cpu)
	modrm.setRM16(cpu, (rm16 | r16))
}

func (cpu *CPU) orR8RM8() {
	reg := &cpu.X86registers
	reg.EIP += 1
	var modrm ModRM
	modrm.parse(cpu)
	r8 := modrm.getR8(cpu)
	rm8 := modrm.getRM8(cpu)
	modrm.setR8(cpu, (rm8 | r8))
}

func (cpu *CPU) orR16RM16() {
	reg := &cpu.X86registers
	reg.EIP += 1
	var modrm ModRM
	modrm.parse(cpu)
	r16 := modrm.getR16(cpu)
	rm16 := modrm.getRM16(cpu)
	modrm.setR16(cpu, (rm16 | r16))
}

func (cpu *CPU) orALImm8() {
	reg := &cpu.X86registers
	mem := cpu.Memory
	value := mem.GetCode8(1)
	reg.EAX = reg.EAX | uint32(value)
	reg.EIP += 2
}

func (cpu *CPU) orAXImm16() {
	reg := &cpu.X86registers
	mem := cpu.Memory
	value := mem.GetCode16(1)
	reg.EAX = reg.EAX | uint32(value)
	reg.EIP += 3
}

func (cpu *CPU) push16ES() {
	reg := &cpu.X86registers
	mem := cpu.Memory
	mem.Push16(reg.ES)
	reg.EIP += 1
}

func (cpu *CPU) pop16ES() {
	reg := &cpu.X86registers
	mem := cpu.Memory
	reg.ES = mem.Pop16()
	reg.EIP += 1
}

func (cpu *CPU) push16SS() {
	reg := &cpu.X86registers
	mem := cpu.Memory
	mem.Push16(reg.SS)
	reg.EIP += 1
}

func (cpu *CPU) pop16SS() {
	reg := &cpu.X86registers
	mem := cpu.Memory
	reg.SS = mem.Pop16()
	reg.EIP += 1
}

func (cpu *CPU) push16DS() {
	reg := &cpu.X86registers
	mem := cpu.Memory
	mem.Push16(reg.DS)
	reg.EIP += 1
}

func (cpu *CPU) pop16DS() {
	reg := &cpu.X86registers
	mem := cpu.Memory
	reg.DS = mem.Pop16()
	reg.EIP += 1
}

func (cpu *CPU) push16Imm16() {
	reg := &cpu.X86registers
	mem := cpu.Memory
	value := mem.GetCode16(1)
	mem.Push16(value)
	reg.EIP += 3
}

func (cpu *CPU) push16Imm8() {
	reg := &cpu.X86registers
	mem := cpu.Memory
	value := mem.GetCode8(1)
	mem.Push16(uint16(value))
	reg.EIP += 2
}

func (cpu *CPU) push16Reg() {
	reg := &cpu.X86registers
	mem := cpu.Memory
	regIndex := mem.GetCode8(0) - 0x50
	mem.Push16(reg.Get16ByIndex(regIndex))
	reg.EIP += 1
}

func (cpu *CPU) pop16Reg() {
	reg := &cpu.X86registers
	mem := cpu.Memory
	regIndex := mem.GetCode8(0) - 0x58
	reg.Set16ByIndex(regIndex, mem.Pop16())
	reg.EIP += 1
}

func (cpu *CPU) overrideOperandTo32() {
	reg := &cpu.X86registers
	reg.EIP += 1
	mem := cpu.Memory
	code := uint8(mem.GetCode8(0))
	if cpu.Instr32[code] == nil {
		log.Fatalf("Not Implemented: 0x%x\n", code)
	}
	cpu.Instr32[code]()
}
