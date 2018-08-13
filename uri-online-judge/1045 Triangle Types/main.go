/*********************************************************************
                            Triangle Types
    https://www.urionlinejudge.com.br/judge/en/problems/view/1045

Read 3 double numbers (A, B and C) representing the sides of a triangle
and arrange them in decreasing order, so that the side A is the biggest
of the three sides. Next, determine the type of triangle that they can
make, based on the following cases always writing an appropriate message:

 • if A ≥ B + C, write the message: NAO FORMA TRIANGULO
 • if A^2 = B^2 + C^2, write the message: TRIANGULO RETANGULO
 • if A^2 > B^2 + C^2, write the message: TRIANGULO OBTUSANGULO
 • if A^2 < B^2 + C^2, write the message: TRIANGULO ACUTANGULO
 • if the three sides are the same size, write the message: TRIANGULO
   EQUILATERO
 • if only two sides are the same and the third one is different, write
   the message: TRIANGULO ISOSCELES

Input
*****
The input contains three double numbers, A (0 < A) , B (0 < B) and C
(0 < C).

Output
******
Print all the classifications of the triangle presented in the input.

+--------------+-----------------------+
| SAMPLE INPUT |     SAMPLE OUTPUT     |
+--------------+-----------------------+
| 7.0 5.0 7.0  | TRIANGULO ACUTANGULO  |
|              | TRIANGULO ISOSCELES   |
+--------------+-----------------------+
| 6.0 6.0 10.0 | TRIANGULO OBTUSANGULO |
|              | TRIANGULO ISOSCELES   |
+--------------+-----------------------+
| 6.0 6.0 6.0  | TRIANGULO ACUTANGULO  |
|              | TRIANGULO EQUILATERO  |
+--------------+-----------------------+
| 5.0 7.0 2.0  | NAO FORMA TRIANGULO   |
+--------------+-----------------------+
| 6.0 8.0 10.0 | TRIANGULO RETANGULO   |
+--------------+-----------------------+
*********************************************************************/

package main

import "fmt"

func main() {
	var A, B, C float64
	fmt.Scan(&A, &B, &C)
	if A < B {
		A, B = B, A
	}
	if A < C {
		A, C = C, A
	}

	if A >= B+C {
		fmt.Println("NAO FORMA TRIANGULO")
		return
	}

	switch {
	case A*A == B*B+C*C:
		fmt.Println("TRIANGULO RETANGULO")
	case A*A > B*B+C*C:
		fmt.Println("TRIANGULO OBTUSANGULO")
	case A*A < B*B+C*C:
		fmt.Println("TRIANGULO ACUTANGULO")
	}

	if A == B && B == C {
		fmt.Println("TRIANGULO EQUILATERO")
	} else if A == B || B == C || A == C {
		fmt.Println("TRIANGULO ISOSCELES")
	}
}
