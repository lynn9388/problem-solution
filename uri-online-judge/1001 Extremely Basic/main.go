/*********************************************************************
                           Extremely Basic
    https://www.urionlinejudge.com.br/judge/en/problems/view/1001

Read 2 integer values and store them in variables, named A and B and
make the sum of these two variables, assigning its result to the variable
X. Print X as shown below. Don't present any message beyond what is
being specified and don't forget to print the end of line after the
result, otherwise you will receive “Presentation Error”.

Input
*****
The input file contain 2 integer values.

Output
******
Print the variable X according to the following example, with a blank
space before and after the equal signal. 'X' is uppercase and you have
to print a blank space before and after the '=' signal.

+--------------+---------------+
| SAMPLE INPUT | SAMPLE OUTPUT |
+--------------+---------------+
| 10           | X = 19        |
| 9            |               |
+--------------+---------------+
| -10          | X = -6        |
| 4            |               |
+--------------+---------------+
| 15           | X = 8         |
| -7           |               |
+--------------+---------------+
*********************************************************************/

package main

import "fmt"

func main() {
	var A, B int
	fmt.Scanf("%d\n%d", &A, &B)
	fmt.Printf("X = %d\n", A+B)
}
