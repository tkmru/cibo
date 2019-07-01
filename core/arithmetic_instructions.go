package cibo

func (cpu *CPU) addRM8R8() {
	reg := &cpu.X86registers
	reg.EIP += 1
	var modrm ModRM
	modrm.parse(cpu)
	r8 := modrm.getR8(cpu)
	rm8 := modrm.getRM8(cpu)
	modrm.setRM8(cpu, rm8+r8)
}

func (cpu *CPU) addRM16R16() {
	reg := &cpu.X86registers
	reg.EIP += 1
	var modrm ModRM
	modrm.parse(cpu)
	r16 := modrm.getR16(cpu)
	rm16 := modrm.getRM16(cpu)
	modrm.setRM16(cpu, rm16+r16)
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

func (cpu *CPU) addR8RM8() {
	reg := &cpu.X86registers
	reg.EIP += 1
	var modrm ModRM
	modrm.parse(cpu)
	r8 := modrm.getR8(cpu)
	rm8 := modrm.getRM8(cpu)
	modrm.setR8(cpu, rm8+r8)
}

func (cpu *CPU) addR16RM16() {
	reg := &cpu.X86registers
	reg.EIP += 1
	var modrm ModRM
	modrm.parse(cpu)
	r16 := modrm.getR16(cpu)
	rm16 := modrm.getRM16(cpu)
	modrm.setR16(cpu, rm16+r16)
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

func (cpu *CPU) addALImm8() {
	reg := &cpu.X86registers
	mem := cpu.Memory
	value := mem.GetCode8(1)
	reg.EAX += uint32(value)
	reg.EIP += 2
}

func (cpu *CPU) addAXImm16() {
	reg := &cpu.X86registers
	mem := cpu.Memory
	value := mem.GetCode16(1)
	reg.EAX += uint32(value)
	reg.EIP += 3
}

func (cpu *CPU) addEAXImm32() {
	reg := &cpu.X86registers
	mem := cpu.Memory
	value := mem.GetCode32(1)
	reg.EAX += uint32(value)
	reg.EIP += 5
}

func (cpu *CPU) orRM8R8() {
	reg := &cpu.X86registers
	reg.EIP += 1
	var modrm ModRM
	modrm.parse(cpu)
	r8 := modrm.getR8(cpu)
	rm8 := modrm.getRM8(cpu)
	modrm.setRM8(cpu, (rm8 | r8))
}

func (cpu *CPU) orRM16R16() {
	reg := &cpu.X86registers
	reg.EIP += 1
	var modrm ModRM
	modrm.parse(cpu)
	r16 := modrm.getR16(cpu)
	rm16 := modrm.getRM16(cpu)
	modrm.setRM16(cpu, (rm16 | r16))
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

func (cpu *CPU) orR8RM8() {
	reg := &cpu.X86registers
	reg.EIP += 1
	var modrm ModRM
	modrm.parse(cpu)
	r8 := modrm.getR8(cpu)
	rm8 := modrm.getRM8(cpu)
	modrm.setR8(cpu, (rm8 | r8))
}

func (cpu *CPU) orR16RM16() {
	reg := &cpu.X86registers
	reg.EIP += 1
	var modrm ModRM
	modrm.parse(cpu)
	r16 := modrm.getR16(cpu)
	rm16 := modrm.getRM16(cpu)
	modrm.setR16(cpu, (rm16 | r16))
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

func (cpu *CPU) orALImm8() {
	reg := &cpu.X86registers
	mem := cpu.Memory
	value := mem.GetCode8(1)
	reg.EAX = reg.EAX | uint32(value)
	reg.EIP += 2
}

func (cpu *CPU) orAXImm16() {
	reg := &cpu.X86registers
	mem := cpu.Memory
	value := mem.GetCode16(1)
	reg.EAX = reg.EAX | uint32(value)
	reg.EIP += 3
}

func (cpu *CPU) orEAXImm32() {
	reg := &cpu.X86registers
	mem := cpu.Memory
	value := mem.GetCode32(1)
	reg.EAX = (reg.EAX | uint32(value))
	reg.EIP += 5
}

func (cpu *CPU) cmpR16RM16() {
	reg := &cpu.X86registers
	reg.EIP += 1
	var modrm ModRM
	modrm.parse(cpu)
	r16 := modrm.getR16(cpu)
	rm16 := modrm.getRM16(cpu)
	result := uint32(r16) - uint32(rm16)
	reg.updateEflagsSub16(r16, rm16, result)
}

func (cpu *CPU) cmpR32RM32() {
	reg := &cpu.X86registers
	reg.EIP += 1
	var modrm ModRM
	modrm.parse(cpu)
	r32 := modrm.getR32(cpu)
	rm32 := modrm.getRM32(cpu)
	result := uint64(r32) - uint64(rm32)
	reg.updateEflagsSub32(r32, rm32, result)
}

func (cpu *CPU) cmpALImm8() {
	reg := &cpu.X86registers
	mem := cpu.Memory
	value := mem.GetCode8(1)
	al := reg.EAX & 0xff
	result := uint64(al) - uint64(value)
	reg.updateEflagsSub32(uint32(al), uint32(value), result)
	reg.EIP += 2
}

func (cpu *CPU) cmpAXImm16() {
	reg := &cpu.X86registers
	mem := cpu.Memory
	value := mem.GetCode16(1)
	ax := uint16(reg.EAX & 0xFF)
	result := uint32(ax) - uint32(value)
	reg.updateEflagsSub16(ax, value, result)
	reg.EIP += 3
}

func (cpu *CPU) cmpEAXImm32() {
	reg := &cpu.X86registers
	mem := cpu.Memory
	value := mem.GetCode32(1)
	eax := reg.EAX
	result := uint64(eax) - uint64(value)
	reg.updateEflagsSub32(eax, value, result)
	reg.EIP += 5
}

func (cpu *CPU) incR16() {
	reg := &cpu.X86registers
	mem := cpu.Memory
	index := mem.GetCode8(0) - 0x40
	value := reg.Get16ByIndex(index) + 1
	reg.Set16ByIndex(index, value)
	reg.EIP += 1
}

func (cpu *CPU) incR32() {
	reg := &cpu.X86registers
	mem := cpu.Memory
	index := mem.GetCode8(0) - 0x40
	value := reg.GetByIndex(index) + 1
	reg.SetByIndex(index, value)
	reg.EIP += 1
}

func (cpu *CPU) decR16() {
	reg := &cpu.X86registers
	mem := cpu.Memory
	index := mem.GetCode8(0) - 0x48
	value := reg.Get16ByIndex(index) - 1
	reg.Set16ByIndex(index, value)
	reg.EIP += 1
}

func (cpu *CPU) decR32() {
	reg := &cpu.X86registers
	mem := cpu.Memory
	index := mem.GetCode8(0) - 0x48
	value := reg.GetByIndex(index) - 1
	reg.SetByIndex(index, value)
	reg.EIP += 1
}

func (cpu *CPU) subAXImm16() {
	reg := &cpu.X86registers
	mem := cpu.Memory
	imm16 := mem.GetCode16(1)
	base := reg.EAX
	reg.EAX = base - uint32(imm16)
	reg.EIP += 3
	reg.updateEflagsSub16(uint16(base), imm16, reg.EAX)
}

func (cpu *CPU) subEAXImm32() {
	reg := &cpu.X86registers
	mem := cpu.Memory
	imm32 := mem.GetCode32(1)
	base := reg.EAX
	reg.EAX = reg.EAX - imm32
	reg.EIP += 5
	reg.updateEflagsSub32(base, imm32, uint64(reg.EAX))
}

func (cpu *CPU) subRM32Imm32(modrm *ModRM) {
	reg := &cpu.X86registers
	mem := cpu.Memory
	rm32 := int32(modrm.getRM32(cpu))
	imm32 := int32(mem.GetSignCode32(0))
	reg.EIP += 4
	modrm.setRM32(cpu, uint32(rm32-imm32))
}

func (cpu *CPU) subRM32Imm8(modrm *ModRM) {
	reg := &cpu.X86registers
	mem := cpu.Memory
	rm32 := modrm.getRM32(cpu)
	imm8 := mem.GetSignCode8(0)
	reg.EIP += 1
	result := uint64(rm32) - uint64(imm8)
	modrm.setRM32(cpu, uint32(result))
	reg.updateEflagsSub32(rm32, uint32(imm8), result)
}

func (cpu *CPU) cmpRM16Imm8(modrm *ModRM) {
	reg := &cpu.X86registers
	mem := cpu.Memory
	rm16 := modrm.getRM16(cpu)
	imm8 := mem.GetSignCode8(0)
	reg.EIP += 1
	result := uint32(rm16) - uint32(imm8)
	reg.updateEflagsSub16(rm16, uint16(imm8), result)
}

func (cpu *CPU) cmpRM32Imm8(modrm *ModRM) {
	reg := &cpu.X86registers
	mem := cpu.Memory
	rm32 := modrm.getRM32(cpu)
	imm8 := mem.GetSignCode8(0)
	reg.EIP += 1
	result := uint64(rm32) - uint64(imm8)
	reg.updateEflagsSub32(rm32, uint32(imm8), result)
}

func (cpu *CPU) incRM32(modrm *ModRM) {
	value := modrm.getRM32(cpu)
	modrm.setRM32(cpu, value+1)
}
