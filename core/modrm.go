package cibo

import (
	"fmt"
	"os"
)

type ModRM struct {
	Mod      uint8
	Rm       uint8
	Opcode   uint8
	RegIndex uint8
	Sib      uint8
	Disp8    int8
	Disp32   uint32
}

func (modrm *ModRM) parse(cpu *CPU) {
	mem := cpu.Memory
	reg := &cpu.X86registers
	code := uint8(mem.GetCode8(0))

	modrm.Mod = ((code & 0xc0) >> 6)
	modrm.Opcode = ((code & 0x38) >> 3)
	modrm.RegIndex = modrm.Opcode
	modrm.Rm = code & 0x7

	reg.EIP += 1

	if modrm.Mod != 3 && modrm.Rm == 4 {
		modrm.Sib = mem.GetCode8(0)
		reg.EIP += 1
	}

	if (modrm.Mod == 0 && modrm.Rm == 5) || modrm.Mod == 2 {
		modrm.Disp32 = mem.GetCode32(0)
		reg.EIP += 4

	} else if modrm.Mod == 1 {
		modrm.Disp8 = mem.GetSignCode8(0)
		modrm.Disp32 = uint32(modrm.Disp8)
		reg.EIP += 1
	}
}

func (modrm *ModRM) calcAddress(cpu *CPU) (result uint32) {
	reg := &cpu.X86registers

	if modrm.Mod == 0 {
		if modrm.Rm == 4 {
			fmt.Println("not implemented ModRM mod = 0, rm = 4")
			os.Exit(0)
		} else if modrm.Rm == 5 {
			result = modrm.Disp32
		} else {
			result = uint32(reg.GetByIndex(modrm.Rm))
		}
	} else if modrm.Mod == 1 {
		if modrm.Rm == 4 {
			fmt.Println("not implemented ModRM mod = 2, rm = 4")
			os.Exit(0)
		} else {
			result = uint32(reg.GetByIndex(modrm.Rm)) + modrm.Disp32
		}
	} else {
		fmt.Println("not implemented ModRM mod = 3")
		os.Exit(0)
	}
	return result
}

func (modrm *ModRM) setRM8(cpu *CPU, value uint8) {
	mem := cpu.Memory
	reg := &cpu.X86registers
	if modrm.Mod == 3 {
		reg.Set8ByIndex(modrm.Rm, value)
	} else {
		address := modrm.calcAddress(cpu)
		mem.Write8(address, value)
	}
}

func (modrm *ModRM) setRM16(cpu *CPU, value uint16) {
	mem := cpu.Memory
	reg := &cpu.X86registers
	if modrm.Mod == 3 {
		reg.Set16ByIndex(modrm.Rm, value)
	} else {
		address := modrm.calcAddress(cpu)
		mem.Write16(address, value)
	}
}

func (modrm *ModRM) setRM32(cpu *CPU, value uint32) {
	mem := cpu.Memory
	reg := &cpu.X86registers
	if modrm.Mod == 3 {
		reg.SetByIndex(modrm.Rm, value)
	} else {
		address := modrm.calcAddress(cpu)
		mem.Write32(address, value)
	}
}

func (modrm *ModRM) getRM8(cpu *CPU) (result uint8) {
	mem := cpu.Memory
	reg := &cpu.X86registers

	if modrm.Mod == 3 {
		result = reg.Get8ByIndex(modrm.Rm)
	} else {
		address := modrm.calcAddress(cpu)
		result = mem.Read8(address)
	}
	return result
}

func (modrm *ModRM) getRM16(cpu *CPU) (result uint16) {
	mem := cpu.Memory
	reg := &cpu.X86registers

	if modrm.Mod == 3 {
		result = reg.Get16ByIndex(modrm.Rm)
	} else {
		address := modrm.calcAddress(cpu)
		result = mem.Read16(address)
	}
	return result
}

func (modrm *ModRM) getRM32(cpu *CPU) (result uint32) {
	mem := cpu.Memory
	reg := &cpu.X86registers

	if modrm.Mod == 3 {
		result = reg.GetByIndex(modrm.Rm)
	} else {
		address := modrm.calcAddress(cpu)
		result = mem.Read32(address)
	}
	return result
}

func (modrm *ModRM) setR8(cpu *CPU, value uint8) {
	reg := &cpu.X86registers
	reg.Set8ByIndex(modrm.RegIndex, value)
}

func (modrm *ModRM) setR16(cpu *CPU, value uint16) {
	reg := &cpu.X86registers
	reg.Set16ByIndex(modrm.RegIndex, value)
}

func (modrm *ModRM) setR32(cpu *CPU, value uint32) {
	reg := &cpu.X86registers
	reg.SetByIndex(modrm.RegIndex, value)
}

func (modrm *ModRM) getR8(cpu *CPU) uint8 {
	reg := &cpu.X86registers
	return reg.Get8ByIndex(modrm.RegIndex)
}

func (modrm *ModRM) getR16(cpu *CPU) uint16 {
	reg := &cpu.X86registers
	return reg.Get16ByIndex(modrm.RegIndex)
}

func (modrm *ModRM) getR32(cpu *CPU) uint32 {
	reg := &cpu.X86registers
	return reg.GetByIndex(modrm.RegIndex)
}
