/*********************************************************************
                               Interval
    https://www.urionlinejudge.com.br/judge/en/problems/view/1037

You must make a program that read a float-point number and print a
message saying in which of following intervals the number belongs:
[0,25] (25,50], (50,75], (75,100]. If the read number is less than
zero or greather than 100, the program must print the message “Fora de
intervalo” that means "Out of Interval".

The symbol '(' represents greather than. For example:
[0,25] indicates numbers between 0 and 25.0000, including both.
(25,50] indicates numbers greather than 25 (25.00001) up to 50.0000000.

Input
*****
The input file contains a floating-point number.

Output
******
The output must be a message like following example.

+--------------+--------------------+
| INPUT SAMPLE |   OUTPUT SAMPLE    |
+--------------+--------------------+
| 25.01        | Intervalo (25,50]  |
+--------------+--------------------+
| 25.00        | Intervalo [0,25]   |
+--------------+--------------------+
| 100.00       | Intervalo (75,100] |
+--------------+--------------------+
| -25.02       | Fora de intervalo  |
+--------------+--------------------+
*********************************************************************/

package main

import (
	"fmt"
)

type bound struct {
	lower int
	upper int
}

var bounds = []bound{
	{0, 25},
	{25, 50},
	{50, 75},
	{75, 100},
}

func isInside(n float64, b bound) bool {
	if b.lower == 0 && n == 0 || float64(b.lower) < n && n <= float64(b.upper) {
		return true
	}
	return false
}

func main() {
	var n float64
	fmt.Scanf("%f", &n)
	if n < 0 || 100 < n {
		fmt.Println("Fora de intervalo")
	} else {
		for _, b := range bounds {
			if isInside(n, b) {
				fmt.Print("Intervalo ")
				if b.lower == 0 {
					fmt.Print("[")
				} else {
					fmt.Print("(")
				}
				fmt.Printf("%d,%d]\n", b.lower, b.upper)
				break
			}
		}
	}
}
