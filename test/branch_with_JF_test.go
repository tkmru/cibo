package test

import (
	"github.com/tkmru/cibo/core"
	"testing"
)

func TestHandlingZF(t *testing.T) {
	beginAddress := 0x7c00
	emu := cibo.NewEmulator(beginAddress, 29, true)
	cpu := emu.CPU
	reg := &cpu.X86registers

	emu.RAM = []byte{0xb8, 0x01, 0x00, 0x00, 0x00, 0x3d, 0x02, 0x00, 0x00, 0x00, 0x75, 0x05, 0xe9, 0xef, 0x83, 0xff,
		0xff, 0xb8, 0x02, 0x00, 0x00, 0x00, 0x3d, 0x02, 0x00, 0x00, 0x00, 0x74, 0xef}
	/*
	    mov eax, 0x1
	    cmp eax, 0x2
	    jnz not_equal

	  equal:
	    jmp 0

	  not_equal:
	    mov eax, 0x2
	    cmp eax, 0x2
	    jz equal
	*/

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
