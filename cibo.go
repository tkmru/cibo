package main

import (
	"fmt"
	"reflect"
)

type Registers struct {
	// GPR
	EAX uint32
	EBX uint32
	ECX uint32
	EDX uint32
	ESP uint32
	EBP uint32
	ESI uint32
	EDI uint32
	// Instruction Register
	EIP uint32
	// Segment Register
	GS uint16
	FS uint16
	ES uint16
	DS uint16
	CS uint16
	SS uint16
	// FLAGS Register
	EFLAGS uint32
}

func (r *Registers) init() {
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
				f.Set(reflect.ValueOf(uint32(2)))
			} else {
				f.Set(reflect.ValueOf(uint32(0)))
			}
		}
	}
}

func (r Registers) dump() {
	v := reflect.ValueOf(&r).Elem()
	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		registerName := t.Field(i).Name
		registerValue := v.Field(i).Interface()

		switch registerName {
		case "EFLAGS":
			fmt.Printf("%d: %s = %d (%032b)\n",
				i+1, registerName, registerValue, registerValue)
		default:
			fmt.Printf("%d: %s = %d\n",
				i+1, registerName, registerValue)
		}
	}
}

func (r *Registers) setCF() {
	// Carry Flag (0 bit)
	r.EFLAGS = r.EFLAGS | 1
}

func (r *Registers) removeCF() {
	// Carry Flag (0 bit)
	r.EFLAGS = r.EFLAGS ^ 1
}

func (r *Registers) setPF() {
	// Parity Flag (2bit)
	r.EFLAGS = r.EFLAGS | 4
}

func (r *Registers) removePF() {
	// Parity Flag (2bit)
	r.EFLAGS = r.EFLAGS ^ 4
}

func (r *Registers) setAF() {
	// Adjust Flag (4bit)
	r.EFLAGS = r.EFLAGS | 16
}

func (r *Registers) removeAF() {
	// Adjust Flag (4bit)
	r.EFLAGS = r.EFLAGS ^ 16
}

func (r *Registers) setZF() {
	// Zero Flag (6bit)
	r.EFLAGS = r.EFLAGS | 64
}

func (r *Registers) removeZF() {
	// Zero Flag (6bit)
	r.EFLAGS = r.EFLAGS ^ 64
}

func (r *Registers) setSF() {
	// Sign Flag (7bit)
	r.EFLAGS = r.EFLAGS | 128
}

func (r *Registers) removeSF() {
	// Sign Flag (7bit)
	r.EFLAGS = r.EFLAGS ^ 128
}

func (r *Registers) setTF() {
	// Trap Flag (8bit)
	r.EFLAGS = r.EFLAGS | 256
}

func (r *Registers) removeTF() {
	// Trap Flag (8bit)
	r.EFLAGS = r.EFLAGS ^ 256
}

func (r *Registers) setIF() {
	// Interrupt Enable Flag (9bit)
	r.EFLAGS = r.EFLAGS | 512
}

func (r *Registers) removeIF() {
	// Interrupt Enable Flag (9bit)
	r.EFLAGS = r.EFLAGS ^ 512
}

func (r *Registers) setDF() {
	// Direction Flag (10bit)
	r.EFLAGS = r.EFLAGS | 1024
}

func (r *Registers) removeDF() {
	// Direction Flag (10bit)
	r.EFLAGS = r.EFLAGS ^ 1024
}

func (r *Registers) setOF() {
	// Overflow Flag (11bit)
	r.EFLAGS = r.EFLAGS | 2048
}

func (r *Registers) removeOF() {
	// Overflow Flag (11bit)
	r.EFLAGS = r.EFLAGS ^ 2048
}

func (r *Registers) setIOPL() {
	// I/O Privilege Level Field (12-13bit)
	r.EFLAGS = r.EFLAGS | 4096 // TODO: fix later
}

func (r *Registers) removeIOPL() {
	// I/O Privilege Level Field (12-13bit)
	r.EFLAGS = r.EFLAGS ^ 4096 // TODO: fix later
}

func (r *Registers) setNT() {
	// Nested Task Flag (14bit)
	r.EFLAGS = r.EFLAGS | 16384
}

func (r *Registers) removeNT() {
	// Nested Task Flag (14bit)
	r.EFLAGS = r.EFLAGS ^ 16384
}

func (r *Registers) setRF() {
	// Resume Flag (16bit)
	r.EFLAGS = r.EFLAGS | 65536
}

func (r *Registers) removeRF() {
	// Resume Flag (16bit)
	r.EFLAGS = r.EFLAGS ^ 65536
}

func (r *Registers) setVM() {
	// Virtual x86 Mode Flag (17bit)
	r.EFLAGS = r.EFLAGS | 131072
}

func (r *Registers) removeVM() {
	// Virtual x86 Mode Flag (17bit)
	r.EFLAGS = r.EFLAGS ^ 131072
}

func (r *Registers) setAC() {
	// Alignment Check Flag (18bit)
	r.EFLAGS = r.EFLAGS | 262144
}

func (r *Registers) removeAC() {
	// Alignment Check Flag (18bit)
	r.EFLAGS = r.EFLAGS ^ 262144
}

func (r *Registers) setVIF() {
	// Virtual Interrupt Flag (19bit)
	r.EFLAGS = r.EFLAGS | 524288
}

func (r *Registers) removeVIF() {
	// Virtual Interrupt Flag (19bit)
	r.EFLAGS = r.EFLAGS ^ 524288
}

func (r *Registers) setVIP() {
	// Virtual Interrupt Pending Flag (20bit)
	r.EFLAGS = r.EFLAGS | 1048576
}

func (r *Registers) removeVIP() {
	// Virtual Interrupt Pending Flag (20bit)
	r.EFLAGS = r.EFLAGS ^ 1048576
}

func (r *Registers) setID() {
	// Identification Flag (21bit)
	r.EFLAGS = r.EFLAGS | 2097152
}

func (r *Registers) removeID() {
	// Identification Flag (21bit)
	r.EFLAGS = r.EFLAGS ^ 2097152
}

func main() {
	r := Registers{}
	r.init()
	r.dump()
}
