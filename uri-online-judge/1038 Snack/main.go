/*********************************************************************
                                Snack
    https://www.urionlinejudge.com.br/judge/en/problems/view/1038

Using the following table, write a program that reads a code and the
amount of an item. After, print the value to pay. This is a very simple
program with the only intention of practice of selection commands.

<img src="https://urionlinejudge.r.worldssl.net/gallery/images/problems/UOJ_1038_en.png">

Input
*****
The input file contains two integer numbers X and Y. X is the product
code and Y is the quantity of this item according to the above table.

Output
******
The output must be a message "Total: R$ " followed by the total value
to be paid, with 2 digits after the decimal point.

+--------------+-----------------+
| INPUT SAMPLE |  OUTPUT SAMPLE  |
+--------------+-----------------+
| 3 2          | Total: R$ 10.00 |
+--------------+-----------------+
| 4 3          | Total: R$ 6.00  |
+--------------+-----------------+
| 2 3          | Total: R$ 13.50 |
+--------------+-----------------+
*********************************************************************/

package main

import "fmt"

type goods struct {
	name  string
	price float64
}

var items = map[int]goods{
	1: {"Cachorro Quente", 4},
	2: {"X-Salada", 4.5},
	3: {"X-Bacon", 5},
	4: {"Torrada siples", 2},
	5: {"Refrigerante", 1.5},
}

func main() {
	var X, Y int
	fmt.Scanf("%d %d", &X, &Y)
	fmt.Printf("Total: R$ %.2f\n", items[X].price*float64(Y))
}
