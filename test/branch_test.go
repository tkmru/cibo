package test

import (
	"fmt"
	"testing"
	"unsafe"

	"github.com/keystone-engine/keystone/bindings/go/keystone"
	"github.com/tkmru/cibo/core"
)

func TestBranchRel8WithZF16(t *testing.T) {
	beginAddress := 0x7c00
	bitMode := 16

	assembly := "" +
		"  mov ax, 0x1;" +
		"  cmp ax, 0x2;" +
		"  jnz not_equal;" +
		"equal:;" +
		"  jmp 0;" +
		"not_equal:;" +
		"  mov ax, 0x2;" +
		"  cmp ax, 0x2;" +
		"  jz equal"

	ks, err := keystone.New(keystone.ARCH_X86, keystone.MODE_16)
	if err != nil {
		panic(err)
	}
	defer ks.Close()

	var binary []byte
	if insn, _, ok := ks.Assemble(assembly, 0); !ok {
		panic(fmt.Errorf("Could not assemble instruction"))
	} else {
		binary = (*(*[]byte)(unsafe.Pointer(&insn)))
	}

	emu := cibo.NewEmulator(bitMode, beginAddress, int64(len(binary)), false, true)
	cpu := emu.CPU
	reg := &cpu.X86registers
	emu.RAM = binary

	reg.Init()
	emu.Run()

	actual := reg.EAX
	expected := uint32(2)
	if actual != expected {
		t.Errorf("got EAX: %v\nexpected EAX: %v", actual, expected)
	}

	if !reg.IsZF() {
		t.Errorf("not set ZF")
	}
}

func TestBranchRel8WithZF32(t *testing.T) {
	beginAddress := 0x7c00
	bitMode := 32

	assembly := "" +
		"  mov eax, 0x1;" +
		"  cmp eax, 0x2;" +
		"  jnz not_equal;" +
		"equal:;" +
		"  jmp 0;" +
		"not_equal:;" +
		"  mov eax, 0x2;" +
		"  cmp eax, 0x2;" +
		"  jz equal"

	ks, err := keystone.New(keystone.ARCH_X86, keystone.MODE_32)
	if err != nil {
		panic(err)
	}
	defer ks.Close()

	var binary []byte
	if insn, _, ok := ks.Assemble(assembly, 0); !ok {
		panic(fmt.Errorf("Could not assemble instruction"))
	} else {
		binary = (*(*[]byte)(unsafe.Pointer(&insn)))
	}

	emu := cibo.NewEmulator(bitMode, beginAddress, 29, false, true)
	cpu := emu.CPU
	reg := &cpu.X86registers
	emu.RAM = binary

	reg.Init()
	emu.Run()

	actual := reg.EAX
	expected := uint32(2)
	if actual != expected {
		t.Errorf("got EAX: %v\nexpected EAX: %v", actual, expected)
	}

	if !reg.IsZF() {
		t.Errorf("not set ZF")
	}
}
