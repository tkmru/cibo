package test

import (
	"fmt"
	"testing"
	"unsafe"

	"github.com/keystone-engine/keystone/bindings/go/keystone"
	"github.com/tkmru/cibo/core"
)

func TestAddRM16R16(t *testing.T) {
	beginAddress := 0x7c00
	bitMode := 16
	emu := cibo.NewEmulator(bitMode, beginAddress, 2, true)
	cpu := emu.CPU
	reg := &cpu.X86registers

	assembly := "add ax, bx"

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
	reg.EAX = 1
	reg.EBX = 2
	emu.Run()

	actual := reg.EAX
	expected := uint32(3)
	if actual != expected {
		t.Errorf("got AX: %v\nexpected AX: %v", actual, expected)
	}
}
