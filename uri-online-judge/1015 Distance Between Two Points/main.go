/*********************************************************************
                     Distance Between Two Points
    https://www.urionlinejudge.com.br/judge/en/problems/view/1015

Read the four values corresponding to the x and y axes of two points
in the plane, p1 (x1, y1) and p2 (x2, y2) and calculate the distance
between them, showing four decimal places after the comma, according
to the formula:

Distance = <img src="https://urionlinejudge.r.worldssl.net/gallery/images/problems/UOJ_1015.png">

Input
*****
The input file contains two lines of data. The first one contains two
double values: x1 y1 and the second one also contains two double values
with one digit after the decimal point: x2 y2.

Output
******
Calculate and print the distance value using the provided formula,
with 4 digits after the decimal point.

+--------------+---------------+
| INPUT SAMPLE | OUTPUT SAMPLE |
+--------------+---------------+
| 1.0 7.0      | 4.4721        |
| 5.0 9.0      |               |
+--------------+---------------+
| -2.5 0.4     | 16.1484       |
| 12.1 7.3     |               |
+--------------+---------------+
| 2.5 -0.4     | 16.4575       |
| -12.2 7.0    |               |
+--------------+---------------+
*********************************************************************/

package main

import (
	"fmt"
	"math"
)

func main() {
	var x1, x2, y1, y2 float64
	fmt.Scanf("%f %f\n%f %f", &x1, &y1, &x2, &y2)
	fmt.Printf("%.4f\n", math.Sqrt(math.Pow(x2-x1, 2.0)+math.Pow(y2-y1, 2.0)))
}
