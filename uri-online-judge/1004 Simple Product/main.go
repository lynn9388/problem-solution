//                           Simple Product

// Read two integer values. After this, calculate the product between
// them and store the result in a variable named PROD. Print the result
// like the example below. Do not forget to print the end of line
// after the result, otherwise you will receive “Presentation Error”.

// Input
//
// The input file contains 2 integer numbers.

// Output
//
// Print PROD according to the following example, with a blank space
// before and after the equal signal.

// +--------------+---------------+
// | SAMPLE INPUT | SAMPLE OUTPUT |
// +--------------+---------------+
// | 3            | PROD = 27     |
// | 9            |               |
// +--------------+---------------+
// | -30          | PROD = -300   |
// | 10           |               |
// +--------------+---------------+
// | 0            | PROD = 0      |
// | 9            |               |
// +--------------+---------------+

// https://www.urionlinejudge.com.br/judge/en/problems/view/1004

package main

import "fmt"

func main() {
	var x, y int
	fmt.Scanf("%d\n%d", &x, &y)
	fmt.Printf("PROD = %d\n", x*y)
}
