/*********************************************************************
                          Bhaskara's Formula
    https://www.urionlinejudge.com.br/judge/en/problems/view/1036

Read 3 floating-point numbers. After, print the roots of bhaskara’s
formula. If it's impossible to calculate the roots because a division
by zero or a square root of a negative number, presents the message
“Impossivel calcular”.

Input
*****
Read 3 floating-point numbers A, B and C.

Output
******
Print the result with 5 digits after the decimal point or the message
if it is impossible to calculate.

+----------------+-----------------------------+
|  SAMPLE INPUT  |        SAMPLE OUTPUT        |
+----------------+-----------------------------+
| 10.0 20.1 5.1  | R1 = -0.29788               |
|                | R2 = -1.71212               |
+----------------+-----------------------------+
| 0.0 20.0 5.0   | Impossivel calcular         |
+----------------+-----------------------------+
| 10.3 203.0 5.0 | R1 = -0.02466               |
|                | R2 = -19.68408              |
+----------------+-----------------------------+
| 10.0 3.0 5.0   | Impossivel calcular         |
+----------------+-----------------------------+
*********************************************************************/

package main

import (
	"fmt"
	"math"
)

func main() {
	var A, B, C float64
	fmt.Scanf("%f %f %f", &A, &B, &C)
	delta := B*B - 4*A*C
	if A == 0 || delta < 0 {
		fmt.Println("Impossivel calcular")
	} else {
		R1 := (-B + math.Sqrt(delta)) / (2 * A)
		R2 := (-B - math.Sqrt(delta)) / (2 * A)
		fmt.Printf("R1 = %.5f\nR2 = %.5f\n", R1, R2)
	}
}
