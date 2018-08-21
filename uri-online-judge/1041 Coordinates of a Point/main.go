/*********************************************************************
                        Coordinates of a Point
    https://www.urionlinejudge.com.br/judge/en/problems/view/1041

Write an algorithm that reads two floating values (x and y), which
should represent the coordinates of a point in a plane. Next, determine
which quadrant the point belongs, or if you are over one of the Cartesian
axes or the origin (x = y = 0).

<img src="https://urionlinejudge.r.worldssl.net/gallery/images/problems/UOJ_1041.png">

If the point is at the origin, write the message "Origem".

If the point is over X axis write "Eixo X", else if the point is over
Y axis write "Eixo Y".

Input
*****
The input contains the coordinates of a point.

Output
******
The output should display the quadrant in which the point is.

+--------------+---------------+
| INPUT SAMPLE | OUTPUT SAMPLE |
+--------------+---------------+
| 4.5 -2.2     | Q4            |
+--------------+---------------+
| 0.1 0.1      | Q1            |
+--------------+---------------+
| 0.0 0.0      | Origem        |
+--------------+---------------+
*********************************************************************/

package main

import "fmt"

func main() {
	var x, y float64
	fmt.Scan(&x, &y)
	var loc string
	if x == 0 {
		if y == 0 {
			loc = "Origem"
		} else {
			loc = "Eixo Y"
		}
	} else if y == 0 {
		loc = "Eixo X"
	} else if x > 0 {
		if y > 0 {
			loc = "Q1"
		} else {
			loc = "Q4"
		}
	} else if x < 0 {
		if y > 0 {
			loc = "Q2"
		} else {
			loc = "Q3"
		}
	}
	fmt.Println(loc)
}
