package array_test

import (
	"testing"

	"github.com/yeqown/algorithm-problems/array"
)

func Test_RingBuffer(t *testing.T) {
	// var size uint = 10
	size := 10
	buf := array.NewRingBuffer(uint(size))

	// start write
	for i := 0; i < size; i++ {
		c := byte(i) + byte(65)
		d := []byte{c}

		err := buf.Write(array.Byts(d))
		// t.Logf("idx=%d,d=%v,err=%v", i, d, err)
		if i == size-1 {
			if err == nil {
				t.Errorf("could not write, but got no error")
				t.FailNow()
			}
		} else {
			if err != nil {
				t.Errorf("invalid write: %v", err)
				t.FailNow()
			}
		}
	}
	// start read
	for i := 0; i < size; i++ {
		d, err := buf.Read()
		t.Logf("idx=%d,d=%v,err=%v", i, d, err)
		if i == size-1 {
			if err == nil {
				t.Errorf("could not read, but got no error")
				t.FailNow()
			}
		} else {
			if err != nil {
				t.Errorf("invalid write: %v", err)
				t.FailNow()
			}
		}
	}
}
