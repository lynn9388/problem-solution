/*********************************************************************
                              Simple Sum
    https://www.urionlinejudge.com.br/judge/en/problems/view/1003

Read two integer values, in this case, the variables A and B. After
this, calculate the sum between them and assign it to the variable
SOMA. Write the value of this variable.

Input
*****
The input file contains 2 integer numbers.

Output
******
Print the variable SOMA with all the capital letters, with a blank
space before and after the equal signal followed by the corresponding
value to the sum of A and B. Like all the problems, don't forget to
print the end of line, otherwise you will receive "Presentation Error"

+---------------+----------------+
| INPUT SAMPLES | OUTPUT SAMPLES |
+---------------+----------------+
| 30            | SOMA = 40      |
| 10            |                |
+---------------+----------------+
| -30           | SOMA = -20     |
| 10            |                |
+---------------+----------------+
| 0             | SOMA = 0       |
| 0             |                |
+---------------+----------------+
*********************************************************************/

package main

import "fmt"

func main() {
	var A, B int
	fmt.Scanf("%d\n%d", &A, &B)
	fmt.Printf("SOMA = %d\n", A+B)
}
