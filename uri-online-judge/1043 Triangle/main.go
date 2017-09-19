/*********************************************************************
                               Triangle
    https://www.urionlinejudge.com.br/judge/en/problems/view/1043

Read three point floating values (A, B and C) and verify if is possible
to make a triangle with them. If it is possible, calculate the perimeter
of the triangle and print the message:

Perimetro = XX.X

If it is not possible, calculate the area of the trapezium which basis
A and B and C as height, and print the message:

Area = XX.X

Input
*****
The input file has tree floating point numbers.

Output
******
Print the result with one digit after the decimal point.

+--------------+------------------+
| SAMPLE INPUT |  SAMPLE OUTPUT   |
+--------------+------------------+
| 6.0 4.0 2.0  | Area = 10.0      |
+--------------+------------------+
| 6.0 4.0 2.1  | Perimetro = 12.1 |
+--------------+------------------+
*********************************************************************/

package main

import "fmt"

func main() {
	var A, B, C float64
	fmt.Scan(&A, &B, &C)
	if isTriangle(A, B, C) {
		fmt.Printf("Perimetro = %.1f\n", A+B+C)
	} else {
		fmt.Printf("Area = %.1f\n", (A+B)*C/2)
	}
}

func isTriangle(a, b, c float64) bool {
	if a+b > c && b+c > a && a+c > b {
		return true
	}
	return false
}
