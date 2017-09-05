/*********************************************************************
                                 Area
    https://www.urionlinejudge.com.br/judge/en/problems/view/1012

Make a program that reads three floating point values: A, B and C.
Then, calculate and show:

a) the area of the rectangled triangle that has base A and height C.

b) the area of the radius's circle C. (pi = 3.14159)

c) the area of the trapezium which has A and B by base, and C by height.

d) the area of ​​the square that has side B.

e) the area of the rectangle that has sides A and B.

Input
*****
The input file contains three double values with one digit after the
decimal point.

Output
******
The output file must contain 5 lines of data. Each line corresponds to
one of the areas described above, always with a corresponding message
(in Portuguese) and one space between the two points and the value.
The value calculated must be presented with 3 digits after the decimal
point.

+----------------+--------------------------------+
|  SAMPLE INPUT  |         SAMPLE OUTPUT          |
+----------------+--------------------------------+
| 3.0 4.0 5.2    | TRIANGULO: 7.800               |
|                | CIRCULO: 84.949                |
|                | TRAPEZIO: 18.200               |
|                | QUADRADO: 16.000               |
|                | RETANGULO: 12.000              |
+----------------+--------------------------------+
| 12.7 10.4 15.2 | TRIANGULO: 96.520              |
|                | CIRCULO: 725.833               |
|                | TRAPEZIO´: 175.560             |
|                | QUADRADO: 108.160              |
|                | RETANGULO: 132.080             |
+----------------+--------------------------------+
*********************************************************************/

package main

import "fmt"

const PI = 3.14159

func main() {
	var A, B, C float64
	fmt.Scanf("%f %f %f", &A, &B, &C)
	fmt.Printf("TRIANGULO: %.3f\nCIRCULO: %.3f\nTRAPEZIO: %.3f\nQUADRADO: %.3f\nRETANGULO: %.3f\n",
		A*C/2.0, PI*C*C, (A+B)*C/2.0, B*B, A*B)
}
