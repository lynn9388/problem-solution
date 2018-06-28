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
		if 'A' <= v && v <= 'Z' {
			v += 13
			if v > 'Z' {
				v -= 26
			}
		} else if 'a' <= v && v <= 'z' {
			v += 13
			if v > 'z' {
				v -= 26
			}
		}
		b[i] = v
	}
	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
