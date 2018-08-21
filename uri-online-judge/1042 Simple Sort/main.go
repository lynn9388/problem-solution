/*********************************************************************
                             Simple Sort
    https://www.urionlinejudge.com.br/judge/en/problems/view/1042

Read three integers and sort them in ascending order. After, print
these values in ascending order, a blank line and then the values in
the sequence as they were readed.

Input
*****
The input contains three integer numbers.

Output
******
Present the output as requested above.

+--------------+---------------+
| INPUT SAMPLE | OUTPUT SAMPLE |
+--------------+---------------+
| 7 21 -14     | -14           |
|              | 7             |
|              | 21            |
|              |               |
|              | 7             |
|              | 21            |
|              | -14           |
+--------------+---------------+
| -14 21 7     | -14           |
|              | 7             |
|              | 21            |
|              |               |
|              | -14           |
|              | 21            |
|              | 7             |
+--------------+---------------+
*********************************************************************/

package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	temp := []int{a, b, c}
	sort(temp)
	fmt.Printf("%d\n%d\n%d\n\n%d\n%d\n%d\n", temp[0], temp[1], temp[2], a, b, c)
}

func sort(n []int) {
	for i := len(n) - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if n[j] > n[j+1] {
				n[j], n[j+1] = n[j+1], n[j]
			}
		}
	}
}
