/*********************************************************************
                              Banknotes
    https://www.urionlinejudge.com.br/judge/en/problems/view/1018

In this problem you have to read an integer value and calculate the
smallest possible number of banknotes in which the value may be decomposed.
The possible banknotes are 100, 50, 20, 10, 5, 2 e 1. Print the read
value and the list of banknotes.

Input
*****
The input file contains an integer value N (0 < N < 1000000).

Output
******
Print the read number and the minimum quantity of each necessary banknotes
in Portuguese language, as the given example. Do not forget to print
the end of line after each line, otherwise you will receive “Presentation
Error”.

+--------------+--------------------------+
| INPUT SAMPLE |      OUTPUT SAMPLE       |
+--------------+--------------------------+
| 576          | 576                      |
|              | 5 nota(s) de R$ 100,00   |
|              | 1 nota(s) de R$ 50,00    |
|              | 1 nota(s) de R$ 20,00    |
|              | 0 nota(s) de R$ 10,00    |
|              | 1 nota(s) de R$ 5,00     |
|              | 0 nota(s) de R$ 2,00     |
|              | 1 nota(s) de R$ 1,00     |
+--------------+--------------------------+
| 11257        | 11257                    |
|              | 112 nota(s) de R$ 100,00 |
|              | 1 nota(s) de R$ 50,00    |
|              | 0 nota(s) de R$ 20,00    |
|              | 0 nota(s) de R$ 10,00    |
|              | 1 nota(s) de R$ 5,00     |
|              | 1 nota(s) de R$ 2,00     |
|              | 0 nota(s) de R$ 1,00     |
+--------------+--------------------------+
| 503          | 503                      |
|              | 5 nota(s) de R$ 100,00   |
|              | 0 nota(s) de R$ 50,00    |
|              | 0 nota(s) de R$ 20,00    |
|              | 0 nota(s) de R$ 10,00    |
|              | 0 nota(s) de R$ 5,00     |
|              | 1 nota(s) de R$ 2,00     |
|              | 1 nota(s) de R$ 1,00     |
+--------------+--------------------------+
*********************************************************************/

package main

import "fmt"

func main() {
	var m int
	banknote := []int{100, 50, 20, 10, 5, 2, 1}
	fmt.Scanf("%d", &m)
	fmt.Println(m)
	for _, b := range banknote {
		fmt.Printf("%d nota(s) de R$ %d,00\n", m/b, b)
		m %= b
	}
}
