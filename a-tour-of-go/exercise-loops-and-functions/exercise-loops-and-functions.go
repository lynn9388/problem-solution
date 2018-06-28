// https://tour.golang.org/flowcontrol/8
package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	diff := math.Abs(z*z - x)
	for diff > 1e-10 {
		z -= (z*z - x) / (2 * z)
		diff = math.Abs(z*z - x)
		fmt.Println(z, diff)
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
}
