package cibo

import (
  "os"
  "fmt"
)

func (cpu *CPU) createTable() {
	cpu.table[0x01] = cpu.addRM32R32
	cpu.table[0x83] = cpu.code83
	cpu.table[0x89] = cpu.movRM32R32
	cpu.table[0x8b] = cpu.movR32RM32

	for i := 0; i < 8; i++ {
		cpu.table[0xb8+i] = cpu.movR32Imm32
	}

	cpu.table[0xc7] = cpu.movRM32Imm32
	cpu.table[0xe9] = cpu.nearJump
	cpu.table[0xeb] = cpu.shortJump
	cpu.table[0xff] = cpu.codeFF
}

func (cpu *CPU) addRM32R32() {
  reg := cpu.X86registers
	reg.EIP += 1
	var modrm ModRM
	modrm.parse(cpu)
	r32 := modrm.getR32(cpu)
	rm32 := modrm.getRM32(cpu)
	modrm.setRM32(cpu, rm32+r32)
}

func (cpu *CPU) code83() {
  reg := cpu.X86registers
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
  reg := cpu.X86registers
  mem := cpu.Memory
  rm32 := int32(modrm.getRM32(cpu))
  imm8 := int32(mem.GetSignCode8(0))
  reg.EIP += 1
  modrm.setRM32(cpu, uint32(rm32 - imm8))
}

func (cpu *CPU) movRM32R32() {
  var modrm ModRM
  reg := cpu.X86registers
	reg.EIP += 1
	modrm.parse(cpu)
	r32 := modrm.getR32(cpu)
	modrm.setRM32(cpu, r32)
}

func (cpu *CPU) movR32RM32() {
  var modrm ModRM
  reg := cpu.X86registers
	reg.EIP += 1
	modrm.parse(cpu)
	rm32 := modrm.getRM32(cpu)
	modrm.setR32(cpu, rm32)
}

func (cpu *CPU) movR32Imm32() {
  mem := cpu.Memory
	regIndex := mem.GetCode8(0) - 0xb8
	value := mem.GetCode32(1)
  reg := cpu.X86registers
  reg.SetRegister32(regIndex, value)
	reg.EIP += 5
}

func (cpu *CPU) movRM32Imm32() {
  var modrm ModRM
  reg := cpu.X86registers
  mem := cpu.Memory
	reg.EIP += 5
	modrm.parse(cpu)
	value := mem.GetCode32(0)
	modrm.setRM32(cpu, value)
}

func (cpu *CPU) nearJump() {
  reg := cpu.X86registers
  mem := cpu.Memory
	diff := mem.GetSignCode32(1)
	reg.EIP += uint32(diff + 5)
}

func (cpu *CPU) shortJump() {
  reg := cpu.X86registers
  mem := cpu.Memory
	diff := mem.GetCode8(1)
	reg.EIP += uint32(diff + 2)
}

func (cpu *CPU) incRM32(modrm *ModRM) {
  value := modrm.getRM32(cpu)
  modrm.setRM32(cpu, value + 1)
}

func (cpu *CPU) codeFF() {
  var modrm ModRM
  reg := cpu.X86registers
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
