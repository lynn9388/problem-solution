// https://tour.golang.org/flowcontrol/8
package main

import (
	"fmt"
)

func Sqrt1(x float64) float64 {
	z := 1.0
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
		fmt.Println(z, x-z*z)
	}
	return z
}

func Sqrt2(x float64) float64 {
	z, diff := 1.0, 1.0
	for diff < -1e-10 || 1e-10 < diff {
		z -= (z*z - x) / (2 * z)
		diff = x - z*z
		fmt.Println(z, diff)
	}
	return z
}

func Sqrt3(x float64) float64 {
	z, diff := x, 1.0
	for diff < -1e-10 || 1e-10 < diff {
		z -= (z*z - x) / (2 * z)
		diff = x - z*z
		fmt.Println(z, diff)
	}
	return z
}

func Sqrt4(x float64) float64 {
	z, diff := x/2, 1.0
	for diff < -1e-10 || 1e-10 < diff {
		z -= (z*z - x) / (2 * z)
		diff = x - z*z
		fmt.Println(z, diff)
	}
	return z
}

func main() {
	fmt.Println(Sqrt1(2))
	fmt.Println(Sqrt2(2))
	fmt.Println(Sqrt2(100))
	fmt.Println(Sqrt3(100))
	fmt.Println(Sqrt4(100))
}
