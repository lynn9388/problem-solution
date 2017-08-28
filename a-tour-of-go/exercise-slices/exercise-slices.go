// https://tour.golang.org/moretypes/18
package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	m := make([][]uint8, dy)
	for y := 0; y < dy; y++ {
		m[y] = make([]uint8, dx)
		for x := 0; x < dx; x++ {
			m[y][x] = uint8(x ^ y)
		}
	}
	return m
}

func main() {
	pic.Show(Pic)
}
