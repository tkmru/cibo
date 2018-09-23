package cibo

func (cpu *CPU) jo() {
	reg := &cpu.X86registers
	mem := cpu.Memory
	var diff uint32 = 2
	if reg.IsOF() {
		diff += uint32(mem.GetSignCode8(1))
	}
	reg.EIP += diff
}

func (cpu *CPU) jno() {
	reg := &cpu.X86registers
	mem := cpu.Memory
	var diff uint32 = 2
	if !reg.IsOF() {
		diff += uint32(mem.GetSignCode8(1))
	}
	reg.EIP += diff
}

func (cpu *CPU) jc() {
	reg := &cpu.X86registers
	mem := cpu.Memory
	var diff uint32 = 2
	if reg.IsCF() {
		diff += uint32(mem.GetSignCode8(1))
	}
	reg.EIP += diff
}

func (cpu *CPU) jnc() {
	reg := &cpu.X86registers
	mem := cpu.Memory
	var diff uint32 = 2
	if !reg.IsCF() {
		diff += uint32(mem.GetSignCode8(1))
	}
	reg.EIP += diff
}

func (cpu *CPU) jz() {
	reg := &cpu.X86registers
	mem := cpu.Memory
	var diff uint32 = 2
	if reg.IsZF() {
		diff += uint32(mem.GetSignCode8(1))
	}
	reg.EIP += diff
}

func (cpu *CPU) jnz() {
	reg := &cpu.X86registers
	mem := cpu.Memory
	var diff uint32 = 2
	if !reg.IsZF() {
		diff += uint32(mem.GetSignCode8(1))
	}
	reg.EIP += diff
}

func (cpu *CPU) js() {
	reg := &cpu.X86registers
	mem := cpu.Memory
	var diff uint32 = 2
	if reg.IsSF() {
		diff += uint32(mem.GetSignCode8(1))
	}
	reg.EIP += diff
}

func (cpu *CPU) jns() {
	reg := &cpu.X86registers
	mem := cpu.Memory
	var diff uint32 = 2
	if !reg.IsSF() {
		diff += uint32(mem.GetSignCode8(1))
	}
	reg.EIP += diff
}

func (cpu *CPU) jl() {
	reg := &cpu.X86registers
	mem := cpu.Memory
	var diff uint32 = 2
	if reg.IsSF() != reg.IsOF() {
		diff += uint32(mem.GetSignCode8(1))
	}
	reg.EIP += diff
}

// jump if less or equal
func (cpu *CPU) jle() {
	reg := &cpu.X86registers
	mem := cpu.Memory
	var diff uint32 = 2
	if reg.IsZF() || (reg.IsSF() != reg.IsOF()) {
		diff += uint32(mem.GetSignCode8(1))
	}
	reg.EIP += diff
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
