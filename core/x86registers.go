package cibo

import (
	"fmt"
	"reflect"
)

var registerIndex = [8]string{"EAX", "ECX", "EDX", "EBX", "ESP", "EBP", "ESI", "EDI"}

type X86registers struct {
	// GPR
	EAX uint32
	ECX uint32
	EDX uint32
	EBX uint32
	ESP uint32
	EBP uint32
	ESI uint32
	EDI uint32
	// Instruction Register
	EIP uint32
	// Segment Registers
	CS uint16
	DS uint16
	SS uint16
	ES uint16
	FS uint16
	GS uint16
	// FLAGS Register
	EFLAGS uint32
	// MMX registers (MM0 through MM7)
	MM0 uint64
	MM1 uint64
	MM2 uint64
	MM3 uint64
	MM4 uint64
	MM5 uint64
	MM6 uint64
	MM7 uint64
	// TODO: XMM registers (XMM0 through XMM15) and the MXCSR register
	// uint128 doesn't existed

	// Control Registers
	CR0 uint64
	CR1 uint64
	CR2 uint64
	CR3 uint64
	CR4 uint64
	CR5 uint64
	CR6 uint64
	CR7 uint64
}

func (r *X86registers) Init() {
	v := reflect.ValueOf(r).Elem()
	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		f := v.Field(i)
		switch f.Kind() {
		case reflect.Uint16:
			f.Set(reflect.ValueOf(uint16(0)))

		case reflect.Uint32:
			if t.Field(i).Name == "EFLAGS" {
				// Reserved 1st bit, it's always 1 in EFLAGS.
				f.Set(reflect.ValueOf(uint32(2))) // 0b10
			} else if t.Field(i).Name == "EIP" {
				continue
			} else if t.Field(i).Name == "ESP" {
				f.Set(reflect.ValueOf(uint32(0x7c04)))
			} else {
				f.Set(reflect.ValueOf(uint32(0)))
			}
		}
	}
}

func (r X86registers) Dump() {
	v := reflect.ValueOf(&r).Elem()
	t := v.Type()

	fmt.Println("==================== registers ====================")
	for i := 0; i < 24; i++ {
		registerName := t.Field(i).Name
		registerValue := v.Field(i).Interface()

		switch registerName {
		case "EFLAGS":
			fmt.Printf("%02d: %s = 0x%X (%032b)\n",
				i+1, registerName, registerValue, registerValue)
		default:
			fmt.Printf("%02d: %s = 0x%X\n",
				i+1, registerName, registerValue)
		}
	}
}

func (r *X86registers) GetByIndex(index uint8) uint32 {
	registerName := registerIndex[index]
	registersPointer := reflect.ValueOf(r) // pointer to struct - addressable
	registers := registersPointer.Elem()   // struct
	targetRegister := registers.FieldByName(registerName)
	return targetRegister.Interface().(uint32)
}

func (r *X86registers) Get8ByIndex(index uint8) uint8 {
	registerName := registerIndex[index]
	registersPointer := reflect.ValueOf(r) // pointer to struct - addressable
	registers := registersPointer.Elem()   // struct
	targetRegister := registers.FieldByName(registerName)
	return uint8(targetRegister.Interface().(uint32) & 0xFF)
}

func (r *X86registers) SetByIndex(index uint8, value uint32) {
	registerName := registerIndex[index]
	registersPointer := reflect.ValueOf(r) // pointer to struct - addressable
	registers := registersPointer.Elem()   // struct
	targetRegister := registers.FieldByName(registerName)
	targetRegister.Set(reflect.ValueOf(value))
}

func (r *X86registers) Set8ByIndex(index uint8, value uint8) {
	registerName := registerIndex[index]
	registersPointer := reflect.ValueOf(r) // pointer to struct - addressable
	registers := registersPointer.Elem()   // struct
	targetRegister := registers.FieldByName(registerName)
	// Rewrite only lower 8 bits
	value32 := (targetRegister.Interface().(uint32) & 0xFFFFFF00) + uint32(value)
	targetRegister.Set(reflect.ValueOf(value32))
}

func (r *X86registers) updateEflagsSub(v1 uint32, v2 uint32, result uint64) {
	var sign1 int = int(v1 >> 31)
	var sign2 int = int(v2 >> 31)
	var signr int = int((result >> 31) & 1)

	if (result >> 32) == 1 {
		r.SetCF()
	} else {
		r.RemoveCF()
	}

	if result == 0 {
		r.SetZF()
	} else {
		r.RemoveZF()
	}

	if signr == 1 {
		r.SetSF()
	} else {
		r.RemoveSF()
	}

	if (sign1^sign2 == 1) && (sign1^signr == 1) {
		r.SetOF()
	} else {
		r.RemoveOF()
	}
}

func (r *X86registers) updateEflagsAdd(v1 uint32, v2 uint32, result uint64) {
	var sign1 int = int(v1 >> 31)
	var sign2 int = int(v2 >> 31)
	var signr int = int((result >> 31) & 1)

	if (result >> 32) == 1 {
		r.SetCF()
	} else {
		r.RemoveCF()
	}

	if result == 0 {
		r.SetZF()
	} else {
		r.RemoveZF()
	}

	if signr == 1 {
		r.SetSF()
	} else {
		r.RemoveSF()
	}

	if (sign1^sign2 == 0) && (sign1^signr == 1) {
		r.SetOF()
	} else {
		r.RemoveOF()
	}
}

// FLAGS Register
// Carry Flag (0 bit)
func (r *X86registers) IsCF() bool {
	return (r.EFLAGS & 1) != 0
}

func (r *X86registers) SetCF() {
	r.EFLAGS = r.EFLAGS | 1
}

func (r *X86registers) RemoveCF() {
	mask := ^1
	r.EFLAGS &= uint32(mask)
}

// Parity Flag (2bit)
func (r *X86registers) IsPF() bool {
	return (r.EFLAGS & 4) != 0
}

func (r *X86registers) SetPF() {
	r.EFLAGS = r.EFLAGS | 4
}

func (r *X86registers) RemovePF() {
	mask := ^4
	r.EFLAGS &= uint32(mask)
}

// Adjust Flag (4bit)
func (r *X86registers) IsAF() bool {
	return (r.EFLAGS & 16) != 0
}

func (r *X86registers) SetAF() {
	r.EFLAGS = r.EFLAGS | 16
}

func (r *X86registers) RemoveAF() {
	mask := ^16
	r.EFLAGS &= uint32(mask)
}

// Zero Flag (6bit)
func (r *X86registers) IsZF() bool {
	return (r.EFLAGS & 64) != 0
}

func (r *X86registers) SetZF() {
	r.EFLAGS = r.EFLAGS | 64
}

func (r *X86registers) RemoveZF() {
	mask := ^64
	r.EFLAGS &= uint32(mask)
}

// Sign Flag (7bit)
func (r *X86registers) IsSF() bool {
	return (r.EFLAGS & 128) != 0
}

func (r *X86registers) SetSF() {
	r.EFLAGS = r.EFLAGS | 128
}

func (r *X86registers) RemoveSF() {
	mask := ^128
	r.EFLAGS &= uint32(mask)
}

// Trap Flag (8bit)
func (r *X86registers) IsTF() bool {
	return (r.EFLAGS & 256) != 0
}

func (r *X86registers) SetTF() {
	r.EFLAGS = r.EFLAGS | 256
}

func (r *X86registers) RemoveTF() {
	mask := ^256
	r.EFLAGS &= uint32(mask)

}

// Interrupt Enable Flag (9bit)
func (r *X86registers) IsEF() bool {
	return (r.EFLAGS & 512) != 0
}

func (r *X86registers) SetIF() {
	r.EFLAGS = r.EFLAGS | 512
}

func (r *X86registers) RemoveIF() {
	mask := ^512
	r.EFLAGS &= uint32(mask)
}

// Direction Flag (10bit)
func (r *X86registers) IsDF() bool {
	return (r.EFLAGS & 1024) != 0
}

func (r *X86registers) SetDF() {
	r.EFLAGS = r.EFLAGS | 1024
}

func (r *X86registers) RemoveDF() {
	mask := ^1024
	r.EFLAGS &= uint32(mask)
}

// Overflow Flag (11bit)
func (r *X86registers) IsOF() bool {
	return (r.EFLAGS & 2048) != 0
}

func (r *X86registers) SetOF() {
	r.EFLAGS = r.EFLAGS | 2048
}

func (r *X86registers) RemoveOF() {
	mask := ^2048
	r.EFLAGS &= uint32(mask)
}

// I/O Privilege Level Field (12-13bit)
func (r *X86registers) IsIOPL() bool {
	return (r.EFLAGS & 4096) != 0
}

func (r *X86registers) SetIOPL() {
	r.EFLAGS = r.EFLAGS | 4096 // TODO: fix later
}

func (r *X86registers) RemoveIOPL() {
	mask := ^4096 // TODO: fix later
	r.EFLAGS &= uint32(mask)
}

// Nested Task Flag (14bit)
func (r *X86registers) IsNT() bool {
	return (r.EFLAGS & 16384) != 0
}

func (r *X86registers) SetNT() {
	r.EFLAGS = r.EFLAGS | 16384
}

func (r *X86registers) RemoveNT() {
	mask := ^16384
	r.EFLAGS &= uint32(mask)
}

// Resume Flag (16bit)
func (r *X86registers) IsRF() bool {
	return (r.EFLAGS & 65536) != 0
}

func (r *X86registers) SetRF() {
	r.EFLAGS = r.EFLAGS | 65536
}

func (r *X86registers) RemoveRF() {
	mask := ^65536
	r.EFLAGS &= uint32(mask)
}

// Virtual x86 Mode Flag (17bit)
func (r *X86registers) IsVM() bool {
	return (r.EFLAGS & 131072) != 0
}

func (r *X86registers) SetVM() {
	r.EFLAGS = r.EFLAGS | 131072
}

func (r *X86registers) RemoveVM() {
	mask := ^131072
	r.EFLAGS &= uint32(mask)
}

// Alignment Check Flag (18bit)
func (r *X86registers) IsAC() bool {
	return (r.EFLAGS & 262144) != 0
}

func (r *X86registers) SetAC() {
	r.EFLAGS = r.EFLAGS | 262144
}

func (r *X86registers) RemoveAC() {
	mask := ^262144
	r.EFLAGS &= uint32(mask)
}

// Virtual Interrupt Flag (19bit)
func (r *X86registers) IsVIF() bool {
	return (r.EFLAGS & 524288) != 0
}

func (r *X86registers) SetVIF() {
	r.EFLAGS = r.EFLAGS | 524288
}

func (r *X86registers) RemoveVIF() {
	mask := ^524288
	r.EFLAGS &= uint32(mask)
}

// Virtual Interrupt Pending Flag (20bit)
func (r *X86registers) IsVIP() bool {
	return (r.EFLAGS & 1048576) != 0
}

func (r *X86registers) SetVIP() {
	r.EFLAGS = r.EFLAGS | 1048576
}

func (r *X86registers) RemoveVIP() {
	mask := ^1048576
	r.EFLAGS &= uint32(mask)
}

// Identification Flag (21bit)
func (r *X86registers) IsID() bool {
	return (r.EFLAGS & 2097152) != 0
}

func (r *X86registers) SetID() {
	r.EFLAGS = r.EFLAGS | 2097152
}

func (r *X86registers) RemoveID() {
	mask := ^2097152
	r.EFLAGS &= uint32(mask)
}
