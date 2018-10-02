package cibo

func (cpu *CPU) joRel8() {
	reg := &cpu.X86registers
	mem := cpu.Memory
	var diff uint32 = 2
	if reg.IsOF() {
		diff += uint32(mem.GetSignCode8(1))
	}
	reg.EIP += diff
}

func (cpu *CPU) jnoRel8() {
	reg := &cpu.X86registers
	mem := cpu.Memory
	var diff uint32 = 2
	if !reg.IsOF() {
		diff += uint32(mem.GetSignCode8(1))
	}
	reg.EIP += diff
}

func (cpu *CPU) jcRel8() {
	reg := &cpu.X86registers
	mem := cpu.Memory
	var diff uint32 = 2
	if reg.IsCF() {
		diff += uint32(mem.GetSignCode8(1))
	}
	reg.EIP += diff
}

func (cpu *CPU) jncRel8() {
	reg := &cpu.X86registers
	mem := cpu.Memory
	var diff uint32 = 2
	if !reg.IsCF() {
		diff += uint32(mem.GetSignCode8(1))
	}
	reg.EIP += diff
}

func (cpu *CPU) jzRel8() {
	reg := &cpu.X86registers
	mem := cpu.Memory
	var diff uint32 = 2
	if reg.IsZF() {
		diff += uint32(mem.GetSignCode8(1))
	}
	reg.EIP += diff
}

func (cpu *CPU) jnzRel8() {
	reg := &cpu.X86registers
	mem := cpu.Memory
	var diff uint32 = 2
	if !reg.IsZF() {
		diff += uint32(mem.GetSignCode8(1))
	}
	reg.EIP += diff
}

func (cpu *CPU) jsRel8() {
	reg := &cpu.X86registers
	mem := cpu.Memory
	var diff uint32 = 2
	if reg.IsSF() {
		diff += uint32(mem.GetSignCode8(1))
	}
	reg.EIP += diff
}

func (cpu *CPU) jnsRel8() {
	reg := &cpu.X86registers
	mem := cpu.Memory
	var diff uint32 = 2
	if !reg.IsSF() {
		diff += uint32(mem.GetSignCode8(1))
	}
	reg.EIP += diff
}

func (cpu *CPU) jlRel8() {
	reg := &cpu.X86registers
	mem := cpu.Memory
	var diff uint32 = 2
	if reg.IsSF() != reg.IsOF() {
		diff += uint32(mem.GetSignCode8(1))
	}
	reg.EIP += diff
}

func (cpu *CPU) jleRel8() {
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
