package test

import (
	"github.com/tkmru/cibo/core"
	"testing"
)

func TestHandlingLowBit(t *testing.T) {
	beginAddress := 0x7c00
	emu := cibo.NewEmulator(beginAddress, 29)
	cpu := emu.CPU
	reg := &cpu.X86registers
	reg.Init()
	reg.Set16ByIndex(0, uint16(0xffff))
	actualAX := reg.Get16ByIndex(0)
	expectedAX := uint16(0xffff)
	if actualAX != expectedAX {
		t.Errorf("got AX: %v\nexpected AX: %v", actualAX, expectedAX)
	}
	reg.Set8ByIndex(0, uint8(0xee))
	actualAL := reg.Get8ByIndex(0)
	expectedAL := uint8(0xee)
	if actualAL != expectedAL {
		t.Errorf("got AL: %v\nexpected AL: %v", actualAL, expectedAL)
	}
}
