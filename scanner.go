package duration

import (
	"bufio"
	"io"
	"unsafe"
)

type durationScanner struct {
	*bufio.Scanner
}

func (s *durationScanner) Text() string {
	return byteSliceToString(s.Bytes()) //  avoid allocations
}

func newScanner(r io.Reader) *durationScanner {
	return &durationScanner{
		Scanner: bufio.NewScanner(r),
	}
}

func byteSliceToString(bs []byte) string {
	return *(*string)(unsafe.Pointer(&bs))
}
