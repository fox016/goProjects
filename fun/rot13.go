package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rot rot13Reader) Read (b []byte) (int, error) {
	bytesRead, err := rot.r.Read(b)
	if err != nil {
		return 0, err
	}
	for i := range b {
		if b[i] >= 'A' && b[i] <= 'M' {
			b[i] += 13
		} else if b[i] >= 'N' && b[i] <= 'Z' {
			b[i] -= 13
		} else if b[i] >= 'a' && b[i] <= 'm' {
			b[i] += 13
		} else if b[i] >= 'n' && b[i] <= 'z' {
			b[i] -= 13
		}
	}
	return bytesRead, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!\n")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
