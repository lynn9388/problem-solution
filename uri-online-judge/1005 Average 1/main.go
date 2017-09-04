//                              Average 1

// Read two floating points' values of double precision A and B,
// corresponding to two student's grades. After this, calculate the
// student's average, considering that grade A has weight 3.5 and B
// has weight 7.5. Each grade can be from zero to ten, always with one
// digit after the decimal point. Don’t forget to print the end of
// line after the result, otherwise you will receive “Presentation
// Error”. Don’t forget the space before and after the equal sign.

// Input
//
// The input file contains 2 floating points' values with one digit
// after the decimal point.

// Output
//
// Print MEDIA(average in Portuguese) according to the following example,
// with 5 digits after the decimal point and with a blank space before
// and after the equal signal.

// +--------------+------------------+
// | SAMPLE INPUT |  SAMPLE OUTPUT   |
// +--------------+------------------+
// | 5.0          | MEDIA = 6.43182  |
// | 7.1          |                  |
// +--------------+------------------+
// | 0.0          | MEDIA = 4.84091  |
// | 7.1          |                  |
// +--------------+------------------+
// | 10.0         | MEDIA = 10.00000 |
// | 10.0         |                  |
// +--------------+------------------+

// https://www.urionlinejudge.com.br/judge/en/problems/view/1005

package main

import "fmt"

func main() {
	var A, B float64
	weightA, weightB := 3.5, 7.5
	fmt.Scanf("%f\n%f", &A, &B)
	fmt.Printf("MEDIA = %.5f\n", (A*weightA+B*weightB)/(weightA+weightB))
}