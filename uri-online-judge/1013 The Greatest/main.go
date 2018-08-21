/*********************************************************************
                             The Greatest
    https://www.urionlinejudge.com.br/judge/en/problems/view/1013

Make a program that reads 3 integer values and present the greatest
one followed by the message "eh o maior". Use the following formula:

<img src="https://urionlinejudge.r.worldssl.net/gallery/images/problems/UOJ_1013.png">

Input
*****
The input file contains 3 integer values.

Output
******
Print the greatest of these three values followed by a space and the
message “eh o maior”.

+---------------+----------------+
| INPUT SAMPLES | OUTPUT SAMPLES |
+---------------+----------------+
| 7 14 106      | 106 eh o maior |
+---------------+----------------+
| 217 14 6      | 217 eh o maior |
+---------------+----------------+
*********************************************************************/

package main

import (
	"fmt"
)

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

func main() {
	var x, y, z int
	fmt.Scanf("%d %d %d", &x, &y, &z)
	fmt.Printf("%d eh o maior\n", max(max(x, y), z))
}
