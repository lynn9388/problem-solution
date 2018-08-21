/*********************************************************************
                                Sphere
    https://www.urionlinejudge.com.br/judge/en/problems/view/1011

Make a program that calculates and shows the volume of a sphere being
provided the value of its radius (R) . The formula to calculate the
volume is: (4/3) * pi * R³. Consider (assign) for pi the value 3.14159.

Tip: Use (4/3.0) or (4.0/3) in your formula, because some languages
(including C++) assume that the division's result between two integers
is another integer. :)

Input
*****
The input contains a value of floating point (double precision).

Output
******
The output must be a message "VOLUME" like the following example with
a space before and after the equal signal. The value must be presented
with 3 digits after the decimal point.

+---------------+--------------------------+
| INPUT SAMPLES |      OUTPUT SAMPLES      |
+---------------+--------------------------+
| 3             | VOLUME = 113.097         |
+---------------+--------------------------+
| 15            | VOLUME = 14137.155       |
+---------------+--------------------------+
| 1523          | VOLUME = 14797486501.627 |
+---------------+--------------------------+
*********************************************************************/

package main

import (
	"fmt"
)

var PI = 3.14159

func main() {
	var R float64
	fmt.Scanf("%f", &R)
	fmt.Printf("VOLUME = %.3f\n", 4.0/3*PI*R*R*R)
}
