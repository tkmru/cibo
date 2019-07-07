package test

import (
	"fmt"
	"testing"
	"unsafe"

	"github.com/keystone-engine/keystone/bindings/go/keystone"
	"github.com/tkmru/cibo/core"
)

func TestAdd16(t *testing.T) {
	assembly := "add al, bl;" +
		"add al, 8;" +
		"add ax, bx;" +
		"add ax, 256"

	ks, err := keystone.New(keystone.ARCH_X86, keystone.MODE_16)
	if err != nil {
		panic(err)
	}
	defer ks.Close()

	if insn, _, ok := ks.Assemble(assembly, 0); !ok {
		panic(fmt.Errorf("Could not assemble instruction"))
	} else {
		insnBytes := (*(*[]byte)(unsafe.Pointer(&insn)))
		beginAddress := 0x7c00
		bitMode := 16
		emu := cibo.NewEmulator(bitMode, beginAddress, int64(len(insnBytes)), true)
		cpu := emu.CPU
		reg := &cpu.X86registers
		emu.RAM = insnBytes
		reg.Init()
		reg.EAX = 1
		reg.EBX = 2
		emu.Run()
		actual := reg.EAX
		expected := uint32(269)
		if actual != expected {
			t.Errorf("got AX: %v\nexpected AX: %v", actual, expected)
		}
	}
}

func TestOr16(t *testing.T) {
	assembly := "or al, bl;" +
		"or al, 8;" +
		"or ax, bx;" +
		"or ax, 256"

	ks, err := keystone.New(keystone.ARCH_X86, keystone.MODE_16)
	if err != nil {
		panic(err)
	}
	defer ks.Close()

	if insn, _, ok := ks.Assemble(assembly, 0); !ok {
		panic(fmt.Errorf("Could not assemble instruction"))
	} else {
		insnBytes := (*(*[]byte)(unsafe.Pointer(&insn)))
		beginAddress := 0x7c00
		bitMode := 16
		emu := cibo.NewEmulator(bitMode, beginAddress, int64(len(insnBytes)), true)
		cpu := emu.CPU
		reg := &cpu.X86registers
		emu.RAM = insnBytes
		reg.Init()
		reg.EAX = 3
		reg.EBX = 5
		emu.Run()

		actual := reg.EAX
		expected := uint32(271)
		if actual != expected {
			t.Errorf("got AX: %v\nexpected AX: %v", actual, expected)
		}
	}
}

func TestIncAndDec16(t *testing.T) {
	assembly := "" +
		"mov eax, 0;" +
		"mov ebx, 2;" +
		"inc eax;" +
		"dec ebx"
	ks, err := keystone.New(keystone.ARCH_X86, keystone.MODE_16)
	if err != nil {
		panic(err)
	}
	defer ks.Close()

	if insn, _, ok := ks.Assemble(assembly, 0); !ok {
		panic(fmt.Errorf("Could not assemble instruction"))
	} else {
		insnBytes := (*(*[]byte)(unsafe.Pointer(&insn)))
		beginAddress := 0x7c00
		bitMode := 16
		emu := cibo.NewEmulator(bitMode, beginAddress, int64(len(insnBytes)), true)
		cpu := emu.CPU
		reg := &cpu.X86registers
		emu.RAM = insnBytes
		reg.Init()
		emu.Run()

		actualAX := reg.EAX
		expectedAX := uint32(1)
		if actualAX != expectedAX {
			t.Errorf("got AX: %v\nexpected AX: %v", actualAX, expectedAX)
		}

		actualBX := reg.EBX
		expectedBX := uint32(1)
		if actualBX != expectedBX {
			t.Errorf("got AX: %v\nexpected AX: %v", actualBX, expectedBX)
		}
	}
}

func TestSubAX16(t *testing.T) {
	assembly := "mov ax, 0x40;" +
		"mov bx, 0x10;" +
		"sub ax, 0x10;" +
		"sub ax, bx;" 

	ks, err := keystone.New(keystone.ARCH_X86, keystone.MODE_16)
	if err != nil {
		panic(err)
	}
	defer ks.Close()

	if insn, _, ok := ks.Assemble(assembly, 0); !ok {
		panic(fmt.Errorf("Could not assemble instruction"))
	} else {	
		insnBytes := (*(*[]byte)(unsafe.Pointer(&insn)))
		fmt.Println(insnBytes)
		beginAddress := 0x7c00
		bitMode := 16
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