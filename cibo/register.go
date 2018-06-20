package cibo

import (
	"fmt"
	"reflect"
)

type Registers struct {
	// GPR
	RAX uint64
	RBX uint64
	RCX uint64
	RDX uint64
	RSP uint64
	RBP uint64
	RSI uint64
	RDI uint64
	R8  uint64
	R9  uint64
	R10 uint64
	R11 uint64
	R12 uint64
	R13 uint64
	R14 uint64
	R15 uint64
	// Instruction Register
	RIP uint64
	// Segment Registers
	CS uint16
	DS uint16
	SS uint16
	ES uint16
	FS uint16
	GS uint16
	// FLAGS Register
	RFLAGS uint64
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
	CR0  uint64
	CR1  uint64
	CR2  uint64
	CR3  uint64
	CR4  uint64
	CR5  uint64
	CR6  uint64
	CR7  uint64
	CR8  uint64
	CR9  uint64
	CR10 uint64
	CR11 uint64
	CR12 uint64
	CR13 uint64
	CR14 uint64
	CR15 uint64
	// Extended Feature Enable Register
  IA32_EFER uint64
}

func (r *Registers) Init() {
	v := reflect.ValueOf(r).Elem()
	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		f := v.Field(i)
		switch f.Kind() {
		case reflect.Uint16:
			f.Set(reflect.ValueOf(uint16(0)))

		case reflect.Uint32:
			f.Set(reflect.ValueOf(uint32(0)))

		case reflect.Uint64:
			if t.Field(i).Name == "RFLAGS" {
				// Reserved 1st bit, it's always 1 in RFLAGS.
				f.Set(reflect.ValueOf(uint64(2))) // 0b10
			} else {
				f.Set(reflect.ValueOf(uint64(0)))
			}
		}
	}
}

func (r Registers) Dump() {
	v := reflect.ValueOf(&r).Elem()
	t := v.Type()

	for i := 0; i < 24; i++ {
		registerName := t.Field(i).Name
		registerValue := v.Field(i).Interface()

		switch registerName {
		case "RFLAGS":
			fmt.Printf("%02d: %s = %d (%064b)\n",
				i+1, registerName, registerValue, registerValue)
		default:
			fmt.Printf("%02d: %s = %d\n",
				i+1, registerName, registerValue)
		}
	}
}

// FLAGS Register
// Carry Flag (0 bit)
func (r *Registers) IsCF() bool {
	return (r.RFLAGS & 1) != 0
}

func (r *Registers) SetCF() {
	r.RFLAGS = r.RFLAGS | 1
}

func (r *Registers) RemoveCF() {
	r.RFLAGS = r.RFLAGS ^ 1
}

// Parity Flag (2bit)
func (r *Registers) IsPF() bool {
	return (r.RFLAGS & 4) != 0
}

func (r *Registers) SetPF() {
	r.RFLAGS = r.RFLAGS | 4
}

func (r *Registers) RemovePF() {
	r.RFLAGS = r.RFLAGS ^ 4
}


// Adjust Flag (4bit)
func (r *Registers) IsAF() bool {
	return (r.RFLAGS & 16) != 0
}

func (r *Registers) SetAF() {
	r.RFLAGS = r.RFLAGS | 16
}

func (r *Registers) RemoveAF() {
	r.RFLAGS = r.RFLAGS ^ 16
}


// Zero Flag (6bit)
func (r *Registers) IsZF() bool {
	return (r.RFLAGS & 64) != 0
}

func (r *Registers) SetZF() {
	r.RFLAGS = r.RFLAGS | 64
}

func (r *Registers) RemoveZF() {
	r.RFLAGS = r.RFLAGS ^ 64
}

// Sign Flag (7bit)
func (r *Registers) IsSF() bool {
	return (r.RFLAGS & 128) != 0
}

func (r *Registers) SetSF() {
	r.RFLAGS = r.RFLAGS | 128
}

func (r *Registers) RemoveSF() {
	r.RFLAGS = r.RFLAGS ^ 128
}

// Trap Flag (8bit)
func (r *Registers) IsTF() bool {
	return (r.RFLAGS & 256) != 0
}

func (r *Registers) SetTF() {
	r.RFLAGS = r.RFLAGS | 256
}

func (r *Registers) RemoveTF() {
	r.RFLAGS = r.RFLAGS ^ 256
}

// Interrupt Enable Flag (9bit)
func (r *Registers) IsEF() bool {
	return (r.RFLAGS & 512) != 0
}

func (r *Registers) SetIF() {
	r.RFLAGS = r.RFLAGS | 512
}

func (r *Registers) RemoveIF() {
	r.RFLAGS = r.RFLAGS ^ 512
}

// Direction Flag (10bit)
func (r *Registers) IsDF() bool {
	return (r.RFLAGS & 1024) != 0
}

func (r *Registers) SetDF() {
	r.RFLAGS = r.RFLAGS | 1024
}

func (r *Registers) RemoveDF() {
	r.RFLAGS = r.RFLAGS ^ 1024
}

// Overflow Flag (11bit)
func (r *Registers) IsOF() bool {
	return (r.RFLAGS & 2048) != 0
}

func (r *Registers) SetOF() {
	r.RFLAGS = r.RFLAGS | 2048
}

func (r *Registers) RemoveOF() {
	r.RFLAGS = r.RFLAGS ^ 2048
}

// I/O Privilege Level Field (12-13bit)
func (r *Registers) IsIOPL() bool {
	return (r.RFLAGS & 4096) != 0
}

func (r *Registers) SetIOPL() {
	r.RFLAGS = r.RFLAGS | 4096 // TODO: fix later
}

func (r *Registers) RemoveIOPL() {
	r.RFLAGS = r.RFLAGS ^ 4096 // TODO: fix later
}

// Nested Task Flag (14bit)
func (r *Registers) IsNT() bool {
	return (r.RFLAGS & 16384) != 0
}

func (r *Registers) SetNT() {
	r.RFLAGS = r.RFLAGS | 16384
}

func (r *Registers) RemoveNT() {
	r.RFLAGS = r.RFLAGS ^ 16384
}

// Resume Flag (16bit)
func (r *Registers) IsRF() bool {
	return (r.RFLAGS & 65536) != 0
}

func (r *Registers) SetRF() {
	r.RFLAGS = r.RFLAGS | 65536
}

func (r *Registers) RemoveRF() {
	r.RFLAGS = r.RFLAGS ^ 65536
}

// Virtual x86 Mode Flag (17bit)
func (r *Registers) IsVM() bool {
	return (r.RFLAGS & 131072) != 0
}

func (r *Registers) SetVM() {
	r.RFLAGS = r.RFLAGS | 131072
}

func (r *Registers) RemoveVM() {
	r.RFLAGS = r.RFLAGS ^ 131072
}

// Alignment Check Flag (18bit)
func (r *Registers) IsAC() bool {
	return (r.RFLAGS & 262144) != 0
}

func (r *Registers) SetAC() {
	r.RFLAGS = r.RFLAGS | 262144
}

func (r *Registers) RemoveAC() {
	r.RFLAGS = r.RFLAGS ^ 262144
}

// Virtual Interrupt Flag (19bit)
func (r *Registers) IsVIF() bool {
	return (r.RFLAGS & 524288) != 0
}

func (r *Registers) SetVIF() {
	r.RFLAGS = r.RFLAGS | 524288
}

func (r *Registers) RemoveVIF() {
	r.RFLAGS = r.RFLAGS ^ 524288
}

// Virtual Interrupt Pending Flag (20bit)
func (r *Registers) IsVIP() bool {
	return (r.RFLAGS & 1048576) != 0
}

func (r *Registers) SetVIP() {
	r.RFLAGS = r.RFLAGS | 1048576
}

func (r *Registers) RemoveVIP() {
	r.RFLAGS = r.RFLAGS ^ 1048576
}

// Identification Flag (21bit)
func (r *Registers) IsID() bool {
	return (r.RFLAGS & 2097152) != 0
}

func (r *Registers) SetID() {
	r.RFLAGS = r.RFLAGS | 2097152
}

func (r *Registers) RemoveID() {
	r.RFLAGS = r.RFLAGS ^ 2097152
}
