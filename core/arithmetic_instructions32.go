package cibo

import (
	"fmt"
	"os"
)

func (cpu *CPU) addRM32R32() {
	reg := &cpu.X86registers
	reg.EIP += 1
	var modrm ModRM
	modrm.parse(cpu)
	r32 := modrm.getR32(cpu)
	rm32 := modrm.getRM32(cpu)
	modrm.setRM32(cpu, rm32+r32)
}

func (cpu *CPU) addR32RM32() {
	reg := &cpu.X86registers
	reg.EIP += 1
	var modrm ModRM
	modrm.parse(cpu)
	r32 := modrm.getR32(cpu)
	rm32 := modrm.getRM32(cpu)
	modrm.setR32(cpu, rm32+r32)
}

func (cpu *CPU) addEAXImm32() {
	reg := &cpu.X86registers
	mem := cpu.Memory
	value := mem.GetCode32(1)
	reg.EAX += uint32(value)
	reg.EIP += 5
}

func (cpu *CPU) orRM32R32() {
	reg := &cpu.X86registers
	reg.EIP += 1
	var modrm ModRM
	modrm.parse(cpu)
	r32 := modrm.getR32(cpu)
	rm32 := modrm.getRM32(cpu)
	modrm.setRM32(cpu, (rm32 | r32))
}

func (cpu *CPU) orR32RM32() {
	reg := &cpu.X86registers
	reg.EIP += 1
	var modrm ModRM
	modrm.parse(cpu)
	r32 := modrm.getR32(cpu)
	rm32 := modrm.getRM32(cpu)
	modrm.setR32(cpu, (rm32 | r32))
}

func (cpu *CPU) orEAXImm32() {
	reg := &cpu.X86registers
	mem := cpu.Memory
	value := mem.GetCode32(1)
	reg.EAX = (reg.EAX | uint32(value))
	reg.EIP += 5
}

func (cpu *CPU) cmpR32RM32() {
	reg := &cpu.X86registers
	reg.EIP += 1
	var modrm ModRM
	modrm.parse(cpu)
	r32 := modrm.getR32(cpu)
	rm32 := modrm.getRM32(cpu)
	result := uint64(r32) - uint64(rm32)
	reg.updateEflagsSub(r32, rm32, result)
}

func (cpu *CPU) cmpALImm8() {
	reg := &cpu.X86registers
	mem := cpu.Memory
	value := mem.GetCode8(1)
	al := reg.EAX & 0xff
	result := uint64(al) - uint64(value)
	reg.updateEflagsSub(uint32(al), uint32(value), result)
	reg.EIP += 2
}

func (cpu *CPU) cmpEAXImm32() {
	reg := &cpu.X86registers
	mem := cpu.Memory
	value := mem.GetCode32(1)
	eax := reg.EAX
	result := uint64(eax) - uint64(value)
	reg.updateEflagsSub(eax, value, result)
	reg.EIP += 5
}

func (cpu *CPU) incR32() {
	reg := &cpu.X86registers
	mem := cpu.Memory
	index := mem.GetCode8(0) - 0x40
	value := reg.GetByIndex(index) + 1
	reg.SetByIndex(index, value)
	reg.EIP += 1
}

func (cpu *CPU) subRM32Imm32(modrm *ModRM) {
	reg := &cpu.X86registers
	mem := cpu.Memory
	rm32 := int32(modrm.getRM32(cpu))
	imm32 := int32(mem.GetSignCode32(0))
	reg.EIP += 4
	modrm.setRM32(cpu, uint32(rm32-imm32))
}

func (cpu *CPU) code81() {
	reg := &cpu.X86registers
	reg.EIP += 1
	var modrm ModRM
	modrm.parse(cpu)

	switch modrm.Opcode {
	case 0:
		//cpu.addRM32Imm32(&modrm)
	case 1:
		//cpu.orRM32Imm32(&modrm)
	case 2:
		//cpu.adcRM32Imm32(&modrm)
	case 3:
		//cpu.sbbRM32Imm32(&modrm)
	case 4:
		//cpu.andRM32Imm32(&modrm)
	case 5:
		cpu.subRM32Imm32(&modrm)
	case 6:
		//cpu.xorRM32Imm32(&modrm)
	case 7:
		//cpu.cmpRM32Imm32(&modrm)
	default:
		fmt.Printf("not implemented: 0x81 /%d\n", modrm.Opcode)
		os.Exit(1)
	}
}

func (cpu *CPU) subRM32Imm8(modrm *ModRM) {
	reg := &cpu.X86registers
	mem := cpu.Memory
	rm32 := modrm.getRM32(cpu)
	imm8 := mem.GetSignCode8(0)
	reg.EIP += 1
	result := uint64(rm32) - uint64(imm8)
	modrm.setRM32(cpu, uint32(result))
	reg.updateEflagsSub(rm32, uint32(imm8), result)
}

func (cpu *CPU) cmpRM32Imm8(modrm *ModRM) {
	reg := &cpu.X86registers
	mem := cpu.Memory
	rm32 := modrm.getRM32(cpu)
	imm8 := mem.GetSignCode8(0)
	reg.EIP += 1
	result := uint64(rm32) - uint64(imm8)
	reg.updateEflagsSub(rm32, uint32(imm8), result)
}

func (cpu *CPU) code83() {
	reg := &cpu.X86registers
	reg.EIP += 1
	var modrm ModRM
	modrm.parse(cpu)

	switch modrm.Opcode {
	case 0:
		//cpu.addRM32Imm8(&modrm)
	case 1:
		//cpu.orRM32Imm8(&modrm)
	case 2:
		//cpu.adcRM32Imm8(&modrm)
	case 3:
		//cpu.sbbRM32Imm8(&modrm)
	case 4:
		//cpu.andRM32Imm8(&modrm)
	case 5:
		cpu.subRM32Imm8(&modrm)
	case 6:
		//cpu.xorRM32Imm8(&modrm)
	case 7:
		cpu.cmpRM32Imm8(&modrm)
	default:
		fmt.Printf("not implemented: 0x83 /%d\n", modrm.Opcode)
		os.Exit(1)
	}
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
