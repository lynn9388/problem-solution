//                              Average 2

// Read three values (variables A, B and C), which are the three student's
// grades. Then, calculate the average, considering that grade A has
// weight 2, grade B has weight 3 and the grade C has weight 5. Consider
// that each grade can go from 0 to 10.0, always with one decimal
// place.

// Input
//
// The input file contains 3 values of floating points with one digit
// after the decimal point.

// Output
//
// Print MEDIA(average in Portuguese) according to the following example,
// with a blank space before and after the equal signal.

// +--------------+---------------+
// | SAMPLE INPUT | SAMPLE OUTPUT |
// +--------------+---------------+
// | 5.0          | MEDIA = 6.3   |
// | 6.0          |               |
// | 7.0          |               |
// +--------------+---------------+
// | 5.0          | MEDIA = 9.0   |
// | 10.0         |               |
// | 10.0         |               |
// +--------------+---------------+
// | 10.0         | MEDIA = 7.5   |
// | 10.0         |               |
// | 5.0          |               |
// +--------------+---------------+

// https://www.urionlinejudge.com.br/judge/en/problems/view/1006

package main

import "fmt"

func main() {
	var A, B, C float64
	weightA, weightB, weightC := 2.0, 3.0, 5.0
	fmt.Scanf("%f\n%f\n%f", &A, &B, &C)
	fmt.Printf("MEDIA = %.1f\n", (A*weightA+B*weightB+C*weightC)/(weightA+weightB+weightC))
}
