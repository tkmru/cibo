package cibo

import (
	"fmt"
	"os"
)

func (cpu *CPU) createTable() {
	cpu.InstTable[0x01] = cpu.addRM32R32
	cpu.InstTable[0x05] = cpu.addEAXImm32
	for i := 0; i < 8; i++ {
		cpu.InstTable[0x50+i] = cpu.PushR32
	}
	for i := 0; i < 8; i++ {
		cpu.InstTable[0x58+i] = cpu.PopR32
	}

	cpu.InstTable[0x68] = cpu.PushImm32
	cpu.InstTable[0x6a] = cpu.PushImm8
	cpu.InstTable[0x83] = cpu.code83
	cpu.InstTable[0x89] = cpu.movRM32R32
	cpu.InstTable[0x8b] = cpu.movR32RM32

	for i := 0; i < 8; i++ {
		cpu.InstTable[0xb8+i] = cpu.movR32Imm32
	}

	cpu.InstTable[0xc7] = cpu.movRM32Imm32
	cpu.InstTable[0xe9] = cpu.nearJump
	cpu.InstTable[0xeb] = cpu.shortJump
	cpu.InstTable[0xff] = cpu.codeFF
}

func (cpu *CPU) addRM32R32() {
	reg := &cpu.X86registers
	reg.EIP += 1
	var modrm ModRM
	modrm.parse(cpu)
	r32 := modrm.getR32(cpu)
	rm32 := modrm.getRM32(cpu)
	modrm.setRM32(cpu, rm32+r32)
}

func (cpu *CPU) addEAXImm32() {
	reg := &cpu.X86registers
	mem := cpu.Memory
	value := mem.GetCode32(1)
	reg.EAX += uint32(value)
	reg.EIP += 5
}

func (cpu *CPU) code83() {
	reg := &cpu.X86registers
	reg.EIP += 1
	var modrm ModRM
	modrm.parse(cpu)

	switch modrm.Opcode {
	case 5:
		cpu.subRM32Imm8(&modrm)
	default:
		fmt.Printf("not implemented: 0x83 /%d\n", modrm.Opcode)
		os.Exit(1)
	}
}

func (cpu *CPU) subRM32Imm8(modrm *ModRM) {
	reg := &cpu.X86registers
	mem := cpu.Memory
	rm32 := int32(modrm.getRM32(cpu))
	imm8 := int32(mem.GetSignCode8(0))
	reg.EIP += 1
	modrm.setRM32(cpu, uint32(rm32-imm8))
}

func (cpu *CPU) movRM32R32() {
	var modrm ModRM
	reg := &cpu.X86registers
	modrm.parse(cpu)
	r32 := modrm.getR32(cpu)
	modrm.setRM32(cpu, r32)
	reg.EIP += 1
}

func (cpu *CPU) movR32RM32() {
	var modrm ModRM
	reg := &cpu.X86registers
	modrm.parse(cpu)
	rm32 := modrm.getRM32(cpu)
	modrm.setR32(cpu, rm32)
	reg.EIP += 5
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
	modrm.parse(cpu)
	value := mem.GetCode32(0)
	modrm.setRM32(cpu, value)
	reg.EIP += 5
}

func (cpu *CPU) nearJump() {
	reg := &cpu.X86registers
	mem := cpu.Memory
	diff := mem.GetSignCode32(1)
	reg.EIP += uint32(diff + 5)
}

func (cpu *CPU) shortJump() {
	reg := &cpu.X86registers
	mem := cpu.Memory
	diff := mem.GetCode8(1)
	reg.EIP += uint32(diff + 2)
}

func (cpu *CPU) incRM32(modrm *ModRM) {
	value := modrm.getRM32(cpu)
	modrm.setRM32(cpu, value+1)
}

func (cpu *CPU) codeFF() {
	var modrm ModRM
	reg := &cpu.X86registers
	reg.EIP += 1
	modrm.parse(cpu)
	switch modrm.Opcode {
	case 0:
		cpu.incRM32(&modrm)
	default:
		fmt.Printf("not implemented: FF /%d\n", modrm.Opcode)
		os.Exit(1)
	}
}

func (cpu *CPU) PushR32() {
	reg := &cpu.X86registers
	mem := cpu.Memory
	regIndex := mem.GetCode8(0) - 0x50
	mem.Push(reg.GetByIndex(regIndex))
	reg.EIP += 1
}

func (cpu *CPU) PushImm32() {
	reg := &cpu.X86registers
	mem := cpu.Memory
	value := mem.GetCode32(1)
	mem.Push(value)
	reg.EIP += 5
}

func (cpu *CPU) PushImm8() {
	reg := &cpu.X86registers
	mem := cpu.Memory
	value := mem.GetCode8(1)
	mem.Push(uint32(value))
	reg.EIP += 2
}

func (cpu *CPU) PopR32() {
	reg := &cpu.X86registers
	mem := cpu.Memory
	regIndex := mem.GetCode8(0) - 0x58
	reg.SetByIndex(regIndex, mem.Pop())
	reg.EIP += 1
}
