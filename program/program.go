// program/program.go
//
// Copyright (C) 2018-present Karim Alibhai. All rights reserved

package program

import (
	"os"

	"github.com/karimsa/basic/constants"
	"github.com/karimsa/basic/data"
)

var (
	wordSize = constants.WordSize / 8
)

type Program struct {
	fd *os.File

	instructions [][2]uint16
	data         []uint16
}

func ReadProgram(fpath string) (*Program, error) {
	fd, err := os.Open(fpath)
	if err != nil {
		return nil, err
	}

	return &Program{
		fd: fd,
	}, nil
}

func NewProgram(fpath string) (*Program, error) {
	fd, err := os.OpenFile(fpath, os.O_CREATE, 0755)
	if err != nil {
		return nil, err
	}

	return &Program{
		fd: fd,
	}, nil
}

func (p *Program) Load() {
	addr := 0

	for {
		word, ok := p.readWord()
		if !ok {
			return
		}

		data.UnsafeMemWrite(addr, word)
		addr++
	}
}

func (p *Program) readWord() (uint16, bool) {
	word := make([]byte, wordSize)
	bytesRead, err := p.fd.Read(word)
	if bytesRead != wordSize || err != nil {
		return 0, false
	}

	return (uint16(word[0]) << 8) | uint16(word[1]), true
}

func (p *Program) writeWord(word uint16) bool {
	buff := make([]byte, wordSize)
	a := word >> 8
	buff[0] = uint8(word >> 8)
	buff[1] = uint8(word ^ (a << 8))
	bytesWritten, err := p.fd.Write(buff)
	return bytesWritten == wordSize && err == nil
}

func (p *Program) close() {
	p.fd.Close()
}
