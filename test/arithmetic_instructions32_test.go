package test

import (
	"fmt"
	"testing"
	"unsafe"

	"github.com/keystone-engine/keystone/bindings/go/keystone"
	"github.com/tkmru/cibo/core"
)

func TestSub32(t *testing.T) {
	assembly := "mov eax, 0x40;" +
		"mov ebx, 0x10;" +
		"sub eax, 0x10;" +
		"sub eax, ebx;" 

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
		expectedEAX := uint32(0x20)
		if actualEAX != expectedEAX {
			t.Errorf("got EAX: %v\nexpected EAX: %v", actualEAX, expectedEAX)
		}
	}
}