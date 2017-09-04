//                             Difference

// Read four integer values named A, B, C and D. Calculate and print
// the difference of product A and B by the product of C and D (A * B
// - C * D).

// Input
//
// The input file contains 4 integer values.

// Output
//
// Print DIFERENCA (DIFFERENCE in Portuguese) with all the capital
// letters, according to the following example, with a blank space
// before and after the equal signal.

// +--------------+-----------------+
// | SAMPLE INPUT |  SAMPLE OUTPUT  |
// +--------------+-----------------+
// | 5            | DIFERENCA = -26 |
// | 6            |                 |
// | 7            |                 |
// | 8            |                 |
// +--------------+-----------------+
// | 0            | DIFERENCA = -56 |
// | 0            |                 |
// | 7            |                 |
// | 8            |                 |
// +--------------+-----------------+
// | 5            | DIFERENCA = 86  |
// | 6            |                 |
// | -7           |                 |
// | 8            |                 |
// +--------------+-----------------+

// https://www.urionlinejudge.com.br/judge/en/problems/view/1007

package main

import "fmt"

func main() {
	var A, B, C, D int
	fmt.Scanf("%d\n%d\n%d\n%d", &A, &B, &C, &D)
	fmt.Printf("DIFERENCA = %d\n", A*B-C*D)
}
