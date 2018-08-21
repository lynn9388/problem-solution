/*********************************************************************
                        Game Time with Minutes
    https://www.urionlinejudge.com.br/judge/en/problems/view/1047

Read the start time and end time of a game, in hours and minutes (initial
hour, initial minute, final hour, final minute). Then print the duration
of the game, knowing that the game can begin in a day and finish in
another day,

Obs.: With a maximum game time of 24 hours and the minimum game time
of 1 minute.

Input
*****
Four integer numbers representing the start and end time of the game.

Output
******
Print the duration of the game in hours and minutes, in this format:
“O JOGO DUROU XXX HORA(S) E YYY MINUTO(S)” . Which means: the game
lasted XXX hour(s) and YYY minutes.

+--------------+---------------------------------------+
| INPUT SAMPLE |             OUTPUT SAMPLE             |
+--------------+---------------------------------------+
| 7 8 9 10     | O JOGO DUROU 2 HORA(S) E 2 MINUTO(S)  |
+--------------+---------------------------------------+
| 7 7 7 7      | O JOGO DUROU 24 HORA(S) E 0 MINUTO(S) |
+--------------+---------------------------------------+
| 7 10 8 9     | O JOGO DUROU 0 HORA(S) E 59 MINUTO(S) |
+--------------+---------------------------------------+
*********************************************************************/

package main

import "fmt"

func main() {
	var sh, sm, eh, em, dur int
	fmt.Scan(&sh, &sm, &eh, &em)
	s := sh*60 + sm
	e := eh*60 + em
	if s < e {
		dur = e - s
	} else {
		dur = 24*60 + e - s
	}
	fmt.Printf("O JOGO DUROU %v HORA(S) E %v MINUTO(S)\n", dur/60, dur%60)
}
