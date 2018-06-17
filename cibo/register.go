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

func (r *Registers) setCF() {
	// Carry Flag (0 bit)
	r.RFLAGS = r.RFLAGS | 1
}

func (r *Registers) removeCF() {
	// Carry Flag (0 bit)
	r.RFLAGS = r.RFLAGS ^ 1
}

func (r *Registers) setPF() {
	// Parity Flag (2bit)
	r.RFLAGS = r.RFLAGS | 4
}

func (r *Registers) removePF() {
	// Parity Flag (2bit)
	r.RFLAGS = r.RFLAGS ^ 4
}

func (r *Registers) setAF() {
	// Adjust Flag (4bit)
	r.RFLAGS = r.RFLAGS | 16
}

func (r *Registers) removeAF() {
	// Adjust Flag (4bit)
	r.RFLAGS = r.RFLAGS ^ 16
}

func (r *Registers) setZF() {
	// Zero Flag (6bit)
	r.RFLAGS = r.RFLAGS | 64
}

func (r *Registers) removeZF() {
	// Zero Flag (6bit)
	r.RFLAGS = r.RFLAGS ^ 64
}

func (r *Registers) setSF() {
	// Sign Flag (7bit)
	r.RFLAGS = r.RFLAGS | 128
}

func (r *Registers) removeSF() {
	// Sign Flag (7bit)
	r.RFLAGS = r.RFLAGS ^ 128
}

func (r *Registers) setTF() {
	// Trap Flag (8bit)
	r.RFLAGS = r.RFLAGS | 256
}

func (r *Registers) removeTF() {
	// Trap Flag (8bit)
	r.RFLAGS = r.RFLAGS ^ 256
}

func (r *Registers) setIF() {
	// Interrupt Enable Flag (9bit)
	r.RFLAGS = r.RFLAGS | 512
}

func (r *Registers) removeIF() {
	// Interrupt Enable Flag (9bit)
	r.RFLAGS = r.RFLAGS ^ 512
}

func (r *Registers) setDF() {
	// Direction Flag (10bit)
	r.RFLAGS = r.RFLAGS | 1024
}

func (r *Registers) removeDF() {
	// Direction Flag (10bit)
	r.RFLAGS = r.RFLAGS ^ 1024
}

func (r *Registers) setOF() {
	// Overflow Flag (11bit)
	r.RFLAGS = r.RFLAGS | 2048
}

func (r *Registers) removeOF() {
	// Overflow Flag (11bit)
	r.RFLAGS = r.RFLAGS ^ 2048
}

func (r *Registers) setIOPL() {
	// I/O Privilege Level Field (12-13bit)
	r.RFLAGS = r.RFLAGS | 4096 // TODO: fix later
}

func (r *Registers) removeIOPL() {
	// I/O Privilege Level Field (12-13bit)
	r.RFLAGS = r.RFLAGS ^ 4096 // TODO: fix later
}

func (r *Registers) setNT() {
	// Nested Task Flag (14bit)
	r.RFLAGS = r.RFLAGS | 16384
}

func (r *Registers) removeNT() {
	// Nested Task Flag (14bit)
	r.RFLAGS = r.RFLAGS ^ 16384
}

func (r *Registers) setRF() {
	// Resume Flag (16bit)
	r.RFLAGS = r.RFLAGS | 65536
}

func (r *Registers) removeRF() {
	// Resume Flag (16bit)
	r.RFLAGS = r.RFLAGS ^ 65536
}

func (r *Registers) setVM() {
	// Virtual x86 Mode Flag (17bit)
	r.RFLAGS = r.RFLAGS | 131072
}

func (r *Registers) removeVM() {
	// Virtual x86 Mode Flag (17bit)
	r.RFLAGS = r.RFLAGS ^ 131072
}

func (r *Registers) setAC() {
	// Alignment Check Flag (18bit)
	r.RFLAGS = r.RFLAGS | 262144
}

func (r *Registers) removeAC() {
	// Alignment Check Flag (18bit)
	r.RFLAGS = r.RFLAGS ^ 262144
}

func (r *Registers) setVIF() {
	// Virtual Interrupt Flag (19bit)
	r.RFLAGS = r.RFLAGS | 524288
}

func (r *Registers) removeVIF() {
	// Virtual Interrupt Flag (19bit)
	r.RFLAGS = r.RFLAGS ^ 524288
}

func (r *Registers) setVIP() {
	// Virtual Interrupt Pending Flag (20bit)
	r.RFLAGS = r.RFLAGS | 1048576
}

func (r *Registers) removeVIP() {
	// Virtual Interrupt Pending Flag (20bit)
	r.RFLAGS = r.RFLAGS ^ 1048576
}

func (r *Registers) setID() {
	// Identification Flag (21bit)
	r.RFLAGS = r.RFLAGS | 2097152
}

func (r *Registers) removeID() {
	// Identification Flag (21bit)
	r.RFLAGS = r.RFLAGS ^ 2097152
}
