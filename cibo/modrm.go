package cibo

import (
  "os"
  "fmt"
)

type ModRM struct {
  Mod uint8
  Rm uint8
  Opcode uint8
  RegIndex uint8
  Sib uint8
  Disp8 int8
  Disp32 uint32
};

func (modrm *ModRM) parse(cpu *CPU) {
  mem := cpu.Memory
  reg := cpu.X86registers
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
  reg := cpu.X86registers

  if modrm.Mod == 0 {
    if modrm.Rm == 4 {
      fmt.Println("not implemented ModRM mod = 0, rm = 4")
      os.Exit(0)
    } else if modrm.Rm == 5 {
      result = modrm.Disp32
    } else {
      result = uint32(reg.GetRegister32(modrm.Rm))
    }
  } else if modrm.Mod == 1 {
    if modrm.Rm == 4 {
      fmt.Println("not implemented ModRM mod = 2, rm = 4")
      os.Exit(0)
    } else {
      result = uint32(reg.GetRegister32(modrm.Rm)) + modrm.Disp32
    }
  } else {
    fmt.Println("not implemented ModRM mod = 3")
    os.Exit(0)
  }
  return result
}

func (modrm *ModRM) setRM32(cpu *CPU, value uint32) {
  mem := cpu.Memory
  reg := cpu.X86registers

  if modrm.Mod == 3 {
    reg.SetRegister32(modrm.Rm, value)
  } else {
    address := modrm.calcAddress(cpu)
    mem.WriteMemory32(address, value)
  }
}

func (modrm *ModRM) getRM32(cpu *CPU) (result uint32) {
  mem := cpu.Memory
  reg := cpu.X86registers

  if modrm.Mod == 3 {
    result = reg.GetRegister32(modrm.Rm)
  } else {
    address := modrm.calcAddress(cpu)
    result = mem.ReadMemory32(address)
  }
  return result
}

func (modrm *ModRM) setR32(cpu *CPU, value uint32) {
  reg := cpu.X86registers
  reg.SetRegister32(modrm.RegIndex, value)
}

func (modrm *ModRM) getR32(cpu *CPU) uint32 {
  reg := cpu.X86registers
  return reg.GetRegister32(modrm.RegIndex)
}
