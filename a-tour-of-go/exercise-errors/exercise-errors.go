// https://tour.golang.org/methods/20
package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprint("cannot Sqrt negative number: ", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return float64(0), ErrNegativeSqrt(x)
	} else {
		z := 1.0
		diff := math.Abs(z*z - x)
		for diff > 1e-10 {
			z -= (z*z - x) / (2 * z)
			diff = math.Abs(z*z - x)
		}
		return z, nil
	}
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
