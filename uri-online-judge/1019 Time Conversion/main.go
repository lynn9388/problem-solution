/*********************************************************************
                           Time Conversion
    https://www.urionlinejudge.com.br/judge/en/problems/view/1019

Read an integer value, which is the duration in seconds of a certain
event in a factory, and inform it expressed in hours:minutes:seconds.

Input
*****
The input file contains an integer N.

Output
******
Print the read time in the input file (seconds) converted in
hours:minutes:seconds like the following example.

+--------------+---------------+
| SAMPLE INPUT | SAMPLE OUTPUT |
+--------------+---------------+
| 556          | 0:9:16        |
+--------------+---------------+
| 1            | 0:0:1         |
+--------------+---------------+
| 140153       | 38:55:53      |
+--------------+---------------+
*********************************************************************/

package main

import "fmt"

func main() {
	var t int
	fmt.Scanf("%d", &t)
	h := t / 3600
	t %= 3600
	m := t / 60
	s := t % 60
	fmt.Printf("%d:%d:%d\n", h, m, s)
}
