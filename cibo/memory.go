package cibo

import "log"

type Memory struct {
	RAM []byte
}

func (m *Memory) Read(address uint64) byte {
  log.Printf("read")
	return 0
}

func (m *Memory) Write(address uint64, value byte) {
	log.Printf("write")
}
