/*********************************************************************
                             Age in Days
    https://www.urionlinejudge.com.br/judge/en/problems/view/1020

Read an integer value corresponding to a person's age (in days) and
print it in years, months and days, followed by its respective message
“ano(s)”, “mes(es)”, “dia(s)”.

Note: only to facilitate the calculation, consider the whole year with
365 days and 30 days every month. In the cases of test there will
never a situation that allows 12 months and some days, like 360, 363
or 364. This is just an exercise for the purpose of testing simple
mathematical reasoning.

Input
*****
The input file contains 1 integer value.

Output
******
Print the output, like the following example.

+--------------+----------------------------+
| SAMPLE INPUT |       SAMPLE OUTPUT        |
+--------------+----------------------------+
| 400          | 1 ano(s)                   |
|              | 1 mes(es)                  |
|              | 5 dia(s)                   |
+--------------+----------------------------+
| 800          | 2 ano(s)                   |
|              | 2 mes(es)                  |
|              | 10 dia(s)                  |
+--------------+----------------------------+
| 30           | 0 ano(s)                   |
|              | 1 mes(es)                  |
|              | 0 dia(s)                   |
+--------------+----------------------------+
*********************************************************************/

package main

import "fmt"

func main() {
	var days int
	fmt.Scanf("%d", &days)
	y := days / 365
	days %= 365
	m := days / 30
	d := days % 30
	fmt.Printf("%d ano(s)\n%d mes(es)\n%d dia(s)\n", y, m, d)
}
