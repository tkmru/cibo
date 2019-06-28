package test

import (
	"fmt"
	"testing"
	"unsafe"

	"github.com/keystone-engine/keystone/bindings/go/keystone"
	"github.com/tkmru/cibo/core"
)


func TestMovLowerR8Imm8(t *testing.T) {
	beginAddress := 0x7c00
	bitMode := 32
	emu := cibo.NewEmulator(bitMode, beginAddress, 2, true)
	cpu := emu.CPU
	reg := &cpu.X86registers

	assembly := "mov al, 0x1;" +
		"mov bl, 0x4;"

	ks, err := keystone.New(keystone.ARCH_X86, keystone.MODE_32)
	if err != nil {
		panic(err)
	}
	defer ks.Close()

	if insn, _, ok := ks.Assemble(assembly, 0); !ok {
		panic(fmt.Errorf("Could not assemble instruction"))
	} else {
		emu.RAM = (*(*[]byte)(unsafe.Pointer(&insn)))
	}

	reg.Init()
	emu.Run()

	actualAL := reg.EAX
	expectedAL := uint32(0x1)
	if actualAL != expectedAL {
		t.Errorf("got AL: %v\nexpected AL: %v", actualAL, expectedAL)
	}
	actualBL := reg.EBX
	expectedBL := uint32(0x4)
	if actualBL != expectedBL {
		t.Errorf("got BL: %v\nexpected BL: %v", actualBL, expectedBL)
	}
}

func TestMovHigherR8Imm8(t *testing.T) {
	beginAddress := 0x7c00
	bitMode := 32
	emu := cibo.NewEmulator(bitMode, beginAddress, 2, true)
	cpu := emu.CPU
	reg := &cpu.X86registers

	assembly := "mov ah, 0x1;" +
		"mov bh, 0x4;"

	ks, err := keystone.New(keystone.ARCH_X86, keystone.MODE_32)
	if err != nil {
		panic(err)
	}
	defer ks.Close()

	if insn, _, ok := ks.Assemble(assembly, 0); !ok {
		panic(fmt.Errorf("Could not assemble instruction"))
	} else {
		emu.RAM = (*(*[]byte)(unsafe.Pointer(&insn)))
	}

	reg.Init()
	emu.Run()

	actualAH := reg.EAX >> 8
	expectedAH := uint32(0x1)
	if actualAH != expectedAH {
		t.Errorf("got AH: %v\nexpected AH: %v", actualAH, expectedAH)
	}

	actualBH := reg.EBX >> 8
	expectedBH := uint32(0x4)
	if actualAH != expectedAH {
		t.Errorf("got BH: %v\nexpected BH: %v", actualBH, expectedBH)
	}
}

func TestMovR16Imm16(t *testing.T) {
	assembly := "mov ax, 0x1234;" +
		"mov bx, 0x4321;"

	ks, err := keystone.New(keystone.ARCH_X86, keystone.MODE_32)
	if err != nil {
		panic(err)
	}
	defer ks.Close()

	if insn, _, ok := ks.Assemble(assembly, 0); !ok {
		panic(fmt.Errorf("Could not assemble instruction"))
	} else {
		insnBytes := (*(*[]byte)(unsafe.Pointer(&insn)))
		beginAddress := 0x7c00
		bitMode := 32
		emu := cibo.NewEmulator(bitMode, beginAddress, int64(len(insnBytes)), true)
		cpu := emu.CPU
		reg := &cpu.X86registers
		emu.RAM = insnBytes

		reg.Init()
		emu.Run()

		actualAX := reg.EAX
		expectedAX := uint32(0x1234)
		if actualAX != expectedAX {
			t.Errorf("got AX: %v\nexpected AX: %v", actualAX, expectedAX)
		}

		actualBX := reg.EBX
		expectedBX := uint32(0x4321)
		if actualAX != expectedAX {
			t.Errorf("got BX: %v\nexpected BX: %v", actualBX, expectedBX)
		}
	}
}

func TestMovR32Imm32(t *testing.T) {
	assembly := "mov eax, 0x12345678;" +
		"mov ebx, 0x87654321;"

	ks, err := keystone.New(keystone.ARCH_X86, keystone.MODE_32)
	if err != nil {
		panic(err)
	}
	defer ks.Close()

	if insn, _, ok := ks.Assemble(assembly, 0); !ok {
		panic(fmt.Errorf("Could not assemble instruction"))
	} else {	
		insnBytes := (*(*[]byte)(unsafe.Pointer(&insn)))
		beginAddress := 0x7c00
		bitMode := 32
		emu := cibo.NewEmulator(bitMode, beginAddress, int64(len(insnBytes)), true)
		cpu := emu.CPU
		reg := &cpu.X86registers
		emu.RAM = insnBytes

		reg.Init()
		emu.Run()

		actualEAX := reg.EAX
		expectedEAX := uint32(0x12345678)
		if actualEAX != expectedEAX {
			t.Errorf("got EAX: %v\nexpected EAX: %v", actualEAX, expectedEAX)
		}

		actualEBX := reg.EBX
		expectedEBX := uint32(0x87654321)
		if actualEBX != expectedEBX {
			t.Errorf("got EBX: %v\nexpected EBX: %v", actualEBX, expectedEBX)
		}
	}
}
