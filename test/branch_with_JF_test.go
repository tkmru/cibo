package test

import (
	"fmt"
	"testing"
	"unsafe"

	"github.com/keystone-engine/keystone/bindings/go/keystone"
	"github.com/tkmru/cibo/core"
)

func TestHandlingZF(t *testing.T) {
	beginAddress := 0x7c00
	emu := cibo.NewEmulator(beginAddress, 29, true)
	cpu := emu.CPU
	reg := &cpu.X86registers

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

	if insn, _, ok := ks.Assemble(assembly, 0); !ok {
		panic(fmt.Errorf("Could not assemble instruction"))
	} else {
		emu.RAM = (*(*[]byte)(unsafe.Pointer(&insn)))
	}

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
