/*********************************************************************
                              Multiples
    https://www.urionlinejudge.com.br/judge/en/problems/view/1044

Read two nteger values (A and B). After, the program should print the
message "Sao Multiplos" (are multiples) or "Nao sao Multiplos" (aren’t
multiples), corresponding to the read values.

Input
*****
The input has two integer numbers.

Output
******
Print the relative message to the input as stated above.

+--------------+-------------------+
| SAMPLE INPUT |   SAMPLE OUTPUT   |
+--------------+-------------------+
| 6 24         | Sao Multiplos     |
+--------------+-------------------+
| 6 25         | Nao sao Multiplos |
+--------------+-------------------+
*********************************************************************/

package main

import "fmt"

func main() {
	var A, B int
	fmt.Scan(&A, &B)
	if A%B == 0 || B%A == 0 {
		fmt.Println("Sao Multiplos")
	} else {
		fmt.Println("Nao sao Multiplos")
	}
}
