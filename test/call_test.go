package test

import (
	"fmt"
	"testing"
	"unsafe"

	"github.com/keystone-engine/keystone/bindings/go/keystone"
	"github.com/tkmru/cibo/core"
)

func TestCallRel16(t *testing.T) {
	beginAddress := 0x7c00
	bitMode := 16

	assembly := "" +
		"  mov ax, 0xf1;" +
		"  mov bx, 0x29;" +
		"  call swap;" +
		"  jmp 0;" +
		"swap:;" +
		"  mov cx, bx;" +
		"  mov bx, ax;" +
		"  mov ax, cx;" +
		"  ret"

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

	emu := cibo.NewEmulator(bitMode, beginAddress, int64(len(binary)), true)
	cpu := emu.CPU
	reg := &cpu.X86registers
	emu.RAM = binary

	reg.Init()
	emu.Run()

	actualAX := reg.EAX
	expectedAX := uint32(0x29)
	if actualAX != expectedAX {
		t.Errorf("got AX: %v\nexpected AX: %v", actualAX, expectedAX)
	}

	actualBX := reg.EBX
	expectedBX := uint32(0xf1)
	if actualBX != expectedBX {
		t.Errorf("got BX: %v\nexpected BX: %v", actualBX, expectedBX)
	}
}

func TestCallRel32(t *testing.T) {
	beginAddress := 0x7c00
	bitMode := 32

	assembly := "" +
		"  mov eax, 0xf1;" +
		"  mov ebx, 0x29;" +
		"  call swap;" +
		"  jmp 0;" +
		"swap:;" +
		"  mov ecx, ebx;" +
		"  mov ebx, eax;" +
		"  mov eax, ecx;" +
		"  ret"

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

	emu := cibo.NewEmulator(bitMode, beginAddress, int64(len(binary)), true)
	cpu := emu.CPU
	reg := &cpu.X86registers
	emu.RAM = binary

	reg.Init()
	emu.Run()

	actualEAX := reg.EAX
	expectedEAX := uint32(0x29)
	if actualEAX != expectedEAX {
		t.Errorf("got EAX: %v\nexpected EAX: %v", actualEAX, expectedEAX)
	}

	actualEBX := reg.EBX
	expectedEBX := uint32(0xf1)
	if actualEBX != expectedEBX {
		t.Errorf("got EBX: %v\nexpected EBX: %v", actualEBX, expectedEBX)
	}
}
