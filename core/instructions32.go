package cibo

import (
	"fmt"
	"log"
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
	cpu.Instr32[0x0e] = cpu.push32CS
	cpu.Instr32[0x0f] = cpu.code0Fb32

	cpu.Instr32[0x16] = cpu.push32SS
	cpu.Instr32[0x17] = cpu.pop32SS
	cpu.Instr32[0x1e] = cpu.push32DS
	cpu.Instr32[0x1f] = cpu.pop32DS

	// cpu.Instr32[0x21] = cpu.andRM32R32
	// cpu.Instr32[0x22] = cpu.andR8RM8
	// cpu.Instr32[0x23] = cpu.andR32RM32
	// cpu.Instr32[0x24] = cpu.andALImm8
	// cpu.Instr32[0x25] = cpu.andEAXImm32

	// cpu.Instr32[0x28] = cpu.subRM8R8
	// cpu.Instr32[0x29] = cpu.subRM32R32
	// cpu.Instr32[0x2a] = cpu.subR8RM8
	// cpu.Instr32[0x2b] = cpu.subR32RM32
	// cpu.Instr32[0x2c] = cpu.subALImm8
	cpu.Instr32[0x2d] = cpu.subEAXImm32

	// cpu.Instr32[0x31] = cpu.xorRM32R32
	// cpu.Instr32[0x32] = cpu.xorR8RM8
	// cpu.Instr32[0x33] = cpu.xorR32RM32
	// cpu.Instr32[0x34] = cpu.xorALImm8
	// cpu.Instr32[0x35] = cpu.xorEAXImm32

	cpu.Instr32[0x3b] = cpu.cmpR32RM32
	cpu.Instr32[0x3c] = cpu.cmpALImm8
	cpu.Instr32[0x3d] = cpu.cmpEAXImm32

	for i := 0; i < 8; i++ {
		cpu.Instr32[0x40+i] = cpu.incR32
	}

	for i := 0; i < 8; i++ {
		cpu.Instr32[0x48+i] = cpu.decR32
	}

	for i := 0; i < 8; i++ {
		cpu.Instr32[0x50+i] = cpu.pushR32
	}

	for i := 0; i < 8; i++ {
		cpu.Instr32[0x58+i] = cpu.popR32
	}

	cpu.Instr32[0x66] = cpu.overrideOperandTo16
	cpu.Instr32[0x68] = cpu.push32Imm32
	cpu.Instr32[0x6a] = cpu.push32Imm8

	cpu.Instr32[0x70] = cpu.joRel8
	cpu.Instr32[0x71] = cpu.jnoRel8
	cpu.Instr32[0x72] = cpu.jcRel8
	cpu.Instr32[0x73] = cpu.jncRel8
	cpu.Instr32[0x74] = cpu.jzRel8
	cpu.Instr32[0x75] = cpu.jnzRel8
	cpu.Instr32[0x78] = cpu.jsRel8
	cpu.Instr32[0x79] = cpu.jnsRel8
	cpu.Instr32[0x7c] = cpu.jlRel8
	cpu.Instr32[0x7e] = cpu.jleRel8
	cpu.Instr32[0x81] = cpu.code81b32
	cpu.Instr32[0x83] = cpu.code83b32

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

	cpu.Instr32[0xc3] = cpu.ret32
	cpu.Instr32[0xc7] = cpu.movRM32Imm32
	cpu.Instr32[0xc9] = cpu.leave32
	/*
		0xd8 - 0xdf: x87 FPU Instructions
	*/
	cpu.Instr32[0xe8] = cpu.callRel32
	cpu.Instr32[0xe9] = cpu.jmpRel32
	cpu.Instr32[0xeb] = cpu.jmpRel8
	cpu.Instr32[0xec] = cpu.inALDX
	cpu.Instr32[0xed] = cpu.inEAXDX
	cpu.Instr32[0xee] = cpu.outDXAL
	cpu.Instr32[0xef] = cpu.outDXEAX
	cpu.Instr32[0xff] = cpu.codeFFb32
}

func (cpu *CPU) pushR32() {
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

func (cpu *CPU) popR32() {
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

func (cpu *CPU) push32CS() {
	reg := &cpu.X86registers
	mem := cpu.Memory
	mem.Push32(uint32(reg.CS))
	reg.EIP += 1
}

func (cpu *CPU) pop32CS() {
	reg := &cpu.X86registers
	mem := cpu.Memory
	reg.CS = uint16(mem.Pop32())
	reg.EIP += 1
}

func (cpu *CPU) code0Fb32() {
	cpu.pop32CS()
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

func (cpu *CPU) overrideOperandTo16() {
	reg := &cpu.X86registers
	reg.EIP += 1
	mem := cpu.Memory
	code := uint8(mem.GetCode8(0))
	if cpu.Instr16[code] == nil {
		log.Fatalf("Not Implemented: 0x%x\n", code)
	}
	cpu.Instr16[code]()
}

func (cpu *CPU) ret32() {
	reg := &cpu.X86registers
	mem := cpu.Memory
	reg.EIP = mem.Pop32()
}

func (cpu *CPU) leave32() {
	reg := &cpu.X86registers
	mem := cpu.Memory
	ebp := reg.EBP
	reg.ESP = ebp
	reg.EBP = mem.Pop32()
	reg.EIP += 1
}

func (cpu *CPU) callRel32() {
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

func (cpu *CPU) inEAXDX() {
	reg := &cpu.X86registers
	var address uint16 = uint16(reg.EDX & 0xffff)
	var value uint32 = ioIn32(address)
	reg.EAX = value
	reg.EIP += 1
}

func (cpu *CPU) outDXAL() {
	reg := &cpu.X86registers
	var address uint16 = uint16(reg.EDX & 0xffff)
	AL := uint8(reg.EAX & 0xff)
	ioOut8(address, AL)
	reg.EIP += 1
}

func (cpu *CPU) outDXEAX() {
	reg := &cpu.X86registers
	var address uint16 = uint16(reg.EDX & 0xffff)
	ioOut32(address, reg.EAX)
	reg.EIP += 1
}

func (cpu *CPU) code81b32() {
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

func (cpu *CPU) code83b32() {
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

func (cpu *CPU) codeFFb32() {
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
