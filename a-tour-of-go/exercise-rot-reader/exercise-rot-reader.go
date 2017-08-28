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
		b[i] = v
		if 'A' <= v && v <= 'Z' {
			b[i] += 13
			if b[i] > 'Z' {
				b[i] -= 26
			}
		} else if 'a' <= v && v <= 'z' {
			b[i] += 13
			if b[i] > 'z' {
				b[i] -= 26
			}
		}
	}
	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
