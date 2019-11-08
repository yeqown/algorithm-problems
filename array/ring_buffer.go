package array

import (
	"errors"
)

var (
	errEmpty = errors.New("empty buffer")
	errFull  = errors.New("full ring buffer")
)

// Byts ...
type Byts []byte

// RingBuffer . not concurrent safe
type RingBuffer struct {
	data   []Byts
	pRead  uint
	pWrite uint
	size   uint
}

// NewRingBuffer . at least 2
func NewRingBuffer(size uint) *RingBuffer {
	if size <= 1 {
		panic("too small size")
	}
	return &RingBuffer{
		data:   make([]Byts, size),
		pRead:  0,
		pWrite: 0,
		size:   size,
	}
}

func (buf *RingBuffer) Write(d Byts) error {
	if (buf.pWrite+1)%buf.size == buf.pRead {
		return errFull
	}
	buf.data[buf.pWrite] = d
	buf.pWrite = (buf.pWrite + 1) % buf.size
	return nil
}

func (buf *RingBuffer) Read() (*Byts, error) {
	if buf.pRead == buf.pWrite {
		return nil, errEmpty
	}
	d := buf.data[buf.pRead]
	buf.pRead = (buf.pRead + 1) % buf.size
	return &d, nil
}

// Reset . drop all data
func (buf *RingBuffer) Reset() {
	buf.pRead = 0
	buf.pWrite = 0
}
