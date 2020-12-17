// https://tour.golang.org/methods/23
package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (r rot13Reader) Read(b []byte) (int, error) {
	n, err := r.r.Read(b)
	for i, v := range b {
		if 'Z'-13 < v && v < 'Z' || 'z'-13 < v && v < 'z' {
			v -= 26
		}
		v += 13
		b[i] = v
	}
	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
