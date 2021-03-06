/*********************************************************************
                           Simple Calculate
    https://www.urionlinejudge.com.br/judge/en/problems/view/1010

In this problem, the task is to read a code of a product 1, the number
of units of product 1, the price for one unit of product 1, the code
of a product 2, the number of units of product 2 and the price for one
unit of product 2. After this, calculate and show the amount to be
paid.

Input
*****
The input file contains two lines of data. In each line there will be
3 values: two integers and a floating value with 2 digits after the
decimal point.

Output
******
The output file must be a message like the following example where
"Valor a pagar" means Value to Pay. Remember the space after ":" and
after "R$" symbol. The value must be presented with 2 digits after the
point.

+---------------+-------------------------+
| INPUT SAMPLES |     OUTPUT SAMPLES      |
+---------------+-------------------------+
| 12 1 5.30     | VALOR A PAGAR: R$ 15.50 |
| 16 2 5.10     |                         |
+---------------+-------------------------+
| 13 2 15.30    | VALOR A PAGAR: R$ 51.40 |
| 161 4 5.20    |                         |
+---------------+-------------------------+
| 1 1 15.10     | VALOR A PAGAR: R$ 30.20 |
| 2 1 15.10     |                         |
+---------------+-------------------------+
*********************************************************************/

package main

import "fmt"

func main() {
	var code, num int
	var price, amount float64
	for i := 0; i < 2; i++ {
		fmt.Scanf("%d %d %f", &code, &num, &price)
		amount += float64(num) * price
	}
	fmt.Printf("VALOR A PAGAR: R$ %.2f\n", amount)
}
