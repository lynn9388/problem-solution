// https://tour.golang.org/flowcontrol/8
package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64, times int) float64 {
	z := 1.0
	for i := 0; i < times; i++ {
		z = z - (z*z-x)/(2*z)
	}
	return z
}

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(Sqrt(2, i))
	}
	fmt.Println(math.Sqrt(2))
}
