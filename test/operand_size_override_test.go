package test

import (
	"fmt"
	"testing"
	"unsafe"

	"github.com/keystone-engine/keystone/bindings/go/keystone"
	"github.com/tkmru/cibo/core"
)

func overrideOperandTo16(t *testing.T) {
	beginAddress := 0x7c00
	bitMode := 32
	emu := cibo.NewEmulator(bitMode, beginAddress, 2, false, true)
	cpu := emu.CPU
	reg := &cpu.X86registers

	assembly := "add ax, 16"

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

	actual := reg.EAX
	expected := uint32(32)
	if actual != expected {
		t.Errorf("got EAX: %v\nexpected EAX: %v", actual, expected)
	}
}

func overrideOperandTo32(t *testing.T) {
	beginAddress := 0x7c00
	bitMode := 16
	emu := cibo.NewEmulator(bitMode, beginAddress, 2, false, true)
	cpu := emu.CPU
	reg := &cpu.X86registers

	assembly := "add eax, 32"

	ks, err := keystone.New(keystone.ARCH_X86, keystone.MODE_16)
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

	actual := reg.EAX
	expected := uint32(32)
	if actual != expected {
		t.Errorf("got AX: %v\nexpected AX: %v", actual, expected)
	}
}
