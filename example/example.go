package example

import (
	"time"
	"unsafe"
)

//go:generate go run ../main.go -source example.go

//fixtureGen
type Elf struct {
	Lifetime time.Time
	Age      uint64
	Name     string
	Friend   []unsafe.Pointer
	Familiar
}

//fixtureGen
type Familiar struct {
	Lifetime time.Time
	Name     string
}

type NotGen struct {
	M int
}
