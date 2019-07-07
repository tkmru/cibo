package cibo

import (
	"fmt"
	"log"
	"os"
)

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
	cpu.Instr16[0x0e] = cpu.push16CS
	cpu.Instr16[0x0f] = cpu.code0F16

	cpu.Instr16[0x16] = cpu.push16SS
	cpu.Instr16[0x17] = cpu.pop16SS
	cpu.Instr16[0x1e] = cpu.push16DS
	cpu.Instr16[0x1f] = cpu.pop16DS

	// cpu.Instr16[0x21] = cpu.andRM16R16
	// cpu.Instr16[0x22] = cpu.andR8RM8
	// cpu.Instr16[0x23] = cpu.andR16RM16
	// cpu.Instr16[0x24] = cpu.andALImm8
	// cpu.Instr16[0x25] = cpu.andAXImm16

	// cpu.Instr16[0x28] = cpu.subRM8R8
	cpu.Instr16[0x29] = cpu.subRM16R16
	// cpu.Instr16[0x2a] = cpu.subR8RM8
	cpu.Instr16[0x2b] = cpu.subR16RM16
	// cpu.Instr16[0x2c] = cpu.subALImm8
	cpu.Instr16[0x2d] = cpu.subAXImm16

	// cpu.Instr16[0x31] = cpu.xorRM16R16
	// cpu.Instr16[0x32] = cpu.xorR8RM8
	// cpu.Instr16[0x33] = cpu.xorR16RM16
	// cpu.Instr16[0x34] = cpu.xorALImm8
	// cpu.Instr16[0x35] = cpu.xorAXImm16

	cpu.Instr16[0x3b] = cpu.cmpR16RM16
	cpu.Instr16[0x3c] = cpu.cmpALImm8
	cpu.Instr16[0x3d] = cpu.cmpAXImm16

	for i := 0; i < 8; i++ {
		cpu.Instr16[0x40+i] = cpu.incR16
	}

	for i := 0; i < 8; i++ {
		cpu.Instr16[0x48+i] = cpu.decR16
	}

	for i := 0; i < 8; i++ {
		cpu.Instr16[0x50+i] = cpu.pushR16
	}

	for i := 0; i < 8; i++ {
		cpu.Instr16[0x58+i] = cpu.popR16
	}

	cpu.Instr16[0x66] = cpu.overrideOperandTo32
	cpu.Instr16[0x68] = cpu.push16Imm16
	cpu.Instr16[0x6a] = cpu.push16Imm8

	cpu.Instr16[0x70] = cpu.joRel8
	cpu.Instr16[0x71] = cpu.jnoRel8
	cpu.Instr16[0x72] = cpu.jcRel8
	cpu.Instr16[0x73] = cpu.jncRel8
	cpu.Instr16[0x74] = cpu.jzRel8
	cpu.Instr16[0x75] = cpu.jnzRel8
	cpu.Instr16[0x78] = cpu.jsRel8
	cpu.Instr16[0x79] = cpu.jnsRel8
	cpu.Instr16[0x7c] = cpu.jlRel8
	cpu.Instr16[0x7e] = cpu.jleRel8
	cpu.Instr16[0x83] = cpu.code83b16

	cpu.Instr16[0x88] = cpu.movRM8R8
	cpu.Instr16[0x89] = cpu.movRM16R16
	cpu.Instr16[0x8a] = cpu.movR8RM8
	cpu.Instr16[0x8b] = cpu.movR16RM16

	cpu.Instr16[0x90] = cpu.nop

	for i := 0; i < 8; i++ {
		cpu.Instr16[0xb0+i] = cpu.movR8Imm8
	}

	for i := 0; i < 8; i++ {
		cpu.Instr16[0xb8+i] = cpu.movR16Imm16
	}

	cpu.Instr16[0xc3] = cpu.ret16
	cpu.Instr16[0xc7] = cpu.movRM16Imm16
	cpu.Instr16[0xc9] = cpu.leave16
	/*
		0xd8 - 0xdf: x87 FPU Instructions
	*/
	cpu.Instr16[0xe8] = cpu.callRel16
	cpu.Instr16[0xe9] = cpu.jmpRel16
	cpu.Instr16[0xeb] = cpu.jmpRel8
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

func (cpu *CPU) push16CS() {
	reg := &cpu.X86registers
	mem := cpu.Memory
	mem.Push16(reg.CS)
	reg.EIP += 1
}

func (cpu *CPU) pop16CS() {
	reg := &cpu.X86registers
	mem := cpu.Memory
	reg.CS = mem.Pop16()
	reg.EIP += 1
}

func (cpu *CPU) code0F16() {
	cpu.pop16CS()
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

func (cpu *CPU) pushR16() {
	reg := &cpu.X86registers
	mem := cpu.Memory
	regIndex := mem.GetCode8(0) - 0x50
	mem.Push16(reg.Get16ByIndex(regIndex))
	reg.EIP += 1
}

func (cpu *CPU) popR16() {
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

func (cpu *CPU) code83b16() {
	reg := &cpu.X86registers
	reg.EIP += 1
	var modrm ModRM
	modrm.parse(cpu)

	switch modrm.Opcode {
	case 0:
		//cpu.addRM16Imm8(&modrm)
	case 1:
		//cpu.orRM16Imm8(&modrm)
	case 2:
		//cpu.adcRM16Imm8(&modrm)
	case 3:
		//cpu.sbbRM16Imm8(&modrm)
	case 4:
		//cpu.andRM16Imm8(&modrm)
	case 5:
		cpu.subRM16Imm8(&modrm)
	case 6:
		//cpu.xorRM16Imm8(&modrm)
	case 7:
		cpu.cmpRM16Imm8(&modrm)
	default:
		fmt.Printf("not implemented: 0x83 /%d\n", modrm.Opcode)
		os.Exit(1)
	}
}

func (cpu *CPU) ret16() {
	reg := &cpu.X86registers
	mem := cpu.Memory
	reg.EIP = uint32(mem.Pop16())
}

func (cpu *CPU) leave16() {
	reg := &cpu.X86registers
	mem := cpu.Memory
	ebp := reg.EBP
	reg.ESP = ebp
	reg.EBP = uint32(mem.Pop16())
	reg.EIP += 1
}

func (cpu *CPU) callRel16() {
	reg := &cpu.X86registers
	mem := cpu.Memory
	diff := mem.GetSignCode16(1)
	mem.Push16(uint16(reg.EIP + 3))
	reg.EIP += uint32(uint32(diff) + 3)
}
