package cibo

import (
	"fmt"
	"os"
)

func (cpu *CPU) createTable32() {
	cpu.Instr32[0x00] = cpu.addRM8R8
	cpu.Instr32[0x01] = cpu.addRM32R32
	cpu.Instr32[0x02] = cpu.addR8RM8
	cpu.Instr32[0x03] = cpu.addR32RM32
	cpu.Instr32[0x04] = cpu.addALImm8
	cpu.Instr32[0x05] = cpu.addEAXImm32
	cpu.Instr32[0x06] = cpu.push32ES
	cpu.Instr32[0x07] = cpu.pop32ES
	cpu.Instr32[0x08] = cpu.orRM8R8
	cpu.Instr32[0x09] = cpu.orRM32R32
	cpu.Instr32[0x0a] = cpu.orR8RM8
	cpu.Instr32[0x0b] = cpu.orR32RM32
	cpu.Instr32[0x0c] = cpu.orALImm8
	cpu.Instr32[0x0d] = cpu.orEAXImm32
	cpu.Instr32[0x16] = cpu.push32SS
	cpu.Instr32[0x17] = cpu.pop32SS
	cpu.Instr32[0x1e] = cpu.push32DS
	cpu.Instr32[0x1f] = cpu.pop32DS
	cpu.Instr32[0x3b] = cpu.cmpR32RM32
	cpu.Instr32[0x3c] = cpu.cmpALImm8
	cpu.Instr32[0x3d] = cpu.cmpEAXImm32

	for i := 0; i < 8; i++ {
		cpu.Instr32[0x40+i] = cpu.incR32
	}

	for i := 0; i < 8; i++ {
		cpu.Instr32[0x50+i] = cpu.pushReg
	}

	for i := 0; i < 8; i++ {
		cpu.Instr32[0x58+i] = cpu.popReg
	}

	cpu.Instr32[0x68] = cpu.push32Imm32
	cpu.Instr32[0x6a] = cpu.push32Imm8
	cpu.Instr32[0x70] = cpu.jo
	cpu.Instr32[0x71] = cpu.jno
	cpu.Instr32[0x72] = cpu.jc
	cpu.Instr32[0x73] = cpu.jnc
	cpu.Instr32[0x74] = cpu.jz
	cpu.Instr32[0x75] = cpu.jnz
	cpu.Instr32[0x78] = cpu.js
	cpu.Instr32[0x79] = cpu.jns
	cpu.Instr32[0x7c] = cpu.jl
	cpu.Instr32[0x7e] = cpu.jle
	cpu.Instr32[0x81] = cpu.code81
	cpu.Instr32[0x83] = cpu.code83
	cpu.Instr32[0x88] = cpu.movRM8R8
	cpu.Instr32[0x89] = cpu.movRM32R32
	cpu.Instr32[0x8a] = cpu.movR8RM8
	cpu.Instr32[0x8b] = cpu.movR32RM32
	cpu.Instr32[0x90] = cpu.nop

	for i := 0; i < 8; i++ {
		cpu.Instr32[0xb0+i] = cpu.movR8Imm8
	}

	for i := 0; i < 8; i++ {
		cpu.Instr32[0xb8+i] = cpu.movR32Imm32
	}

	cpu.Instr32[0xc3] = cpu.ret
	cpu.Instr32[0xc7] = cpu.movRM32Imm32
	cpu.Instr32[0xc9] = cpu.leave
	/*
		0xd8 - 0xdf: x87 FPU Instructions
	*/
	cpu.Instr32[0xe8] = cpu.callRelative
	cpu.Instr32[0xe9] = cpu.nearJump
	cpu.Instr32[0xeb] = cpu.shortJump
	cpu.Instr32[0xec] = cpu.inALDX
	cpu.Instr32[0xee] = cpu.outDXAL
	cpu.Instr32[0xff] = cpu.codeFF
}

func (cpu *CPU) pushReg() {
	reg := &cpu.X86registers
	mem := cpu.Memory
	regIndex := mem.GetCode8(0) - 0x50
	mem.Push32(reg.GetByIndex(regIndex))
	reg.EIP += 1
}

func (cpu *CPU) push32Imm32() {
	reg := &cpu.X86registers
	mem := cpu.Memory
	value := mem.GetCode32(1)
	mem.Push32(value)
	reg.EIP += 5
}

func (cpu *CPU) push32Imm8() {
	reg := &cpu.X86registers
	mem := cpu.Memory
	value := mem.GetCode8(1)
	mem.Push32(uint32(value))
	reg.EIP += 2
}

func (cpu *CPU) popReg() {
	reg := &cpu.X86registers
	mem := cpu.Memory
	regIndex := mem.GetCode8(0) - 0x58
	reg.SetByIndex(regIndex, mem.Pop32())
	reg.EIP += 1
}

func (cpu *CPU) push32ES() {
	reg := &cpu.X86registers
	mem := cpu.Memory
	mem.Push32(uint32(reg.ES))
	reg.EIP += 1
}

func (cpu *CPU) pop32ES() {
	reg := &cpu.X86registers
	mem := cpu.Memory
	reg.ES = uint16(mem.Pop32())
	reg.EIP += 1
}

func (cpu *CPU) push32SS() {
	reg := &cpu.X86registers
	mem := cpu.Memory
	mem.Push32(uint32(reg.SS))
	reg.EIP += 1
}

func (cpu *CPU) pop32SS() {
	reg := &cpu.X86registers
	mem := cpu.Memory
	reg.SS = uint16(mem.Pop32())
	reg.EIP += 1
}

func (cpu *CPU) push32DS() {
	reg := &cpu.X86registers
	mem := cpu.Memory
	mem.Push32(uint32(reg.DS))
	reg.EIP += 1
}

func (cpu *CPU) pop32DS() {
	reg := &cpu.X86registers
	mem := cpu.Memory
	reg.DS = uint16(mem.Pop32())
	reg.EIP += 1
}

func (cpu *CPU) ret() {
	reg := &cpu.X86registers
	mem := cpu.Memory
	reg.EIP = mem.Pop32()
}

func (cpu *CPU) leave() {
	reg := &cpu.X86registers
	mem := cpu.Memory
	ebp := reg.EBP
	reg.ESP = ebp
	reg.EBP = mem.Pop32()
	reg.EIP += 1
}

func (cpu *CPU) callRelative() {
	reg := &cpu.X86registers
	mem := cpu.Memory
	diff := mem.GetSignCode32(1)
	mem.Push32(reg.EIP + 5)
	reg.EIP += uint32(diff + 5)
}

func (cpu *CPU) inALDX() {
	reg := &cpu.X86registers
	var address uint16 = uint16(reg.EDX & 0xffff)
	var value uint8 = ioIn8(address)
	AH := reg.EAX & 0xff00
	reg.EAX = (AH + uint32(value))
	reg.EIP += 1
}

func (cpu *CPU) outDXAL() {
	reg := &cpu.X86registers
	var address uint16 = uint16(reg.EDX & 0xffff)
	AL := uint8(reg.EAX & 0xff)
	ioOut8(address, AL)
	reg.EIP += 1
}

func ioIn8(address uint16) uint8 {
	fmt.Println("[cibo] asking for input:")
	switch address {
	case 0x03f8:
		var input []byte = make([]byte, 1)
		os.Stdin.Read(input)
		return uint8(input[0])
		break
	}
	return 0
}

func ioOut8(address uint16, ascii uint8) {
	switch address {
	case 0x03f8:
		fmt.Println(string(ascii))
		break
	}
}
