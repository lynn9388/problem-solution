/*********************************************************************
                              Average 3
    https://www.urionlinejudge.com.br/judge/en/problems/view/1040

Read four numbers (N₁, N₂, N₃, N₄), which one with 1 digit after the
decimal point, corresponding to 4 scores obtained by a student. Calculate
the average with weights 2, 3, 4 e 1 respectively, for these 4 scores
and print the message "Media: " (Average), followed by the calculated
result. If the average was 7.0 or more, print the message "Aluno aprovado."
(Approved Student). If the average was less than 5.0, print the message:
"Aluno reprovado." (Reproved Student). If the average was between 5.0
and 6.9, including these, the program must print the message "Aluno em
exame." (In exam student).

In case of exam, read one more score. Print the message "Nota do exame:
" (Exam score) followed by the typed score. Recalculate the average
(sum the exam score with the previous calculated average and divide by
2) and print the message “Aluno aprovado.” (Approved student) in case
of average 5.0 or more) or "Aluno reprovado." (Reproved student) in
case of average 4.9 or less. For these 2 cases (approved or reproved
after the exam) print the message "Media final: " (Final average)
followed by the final average for this student in the last line.

Input
*****
The input contains four floating point numbers that represent the
students' grades.

Output
******
Print all the answers with one digit after the decimal point.

+-----------------+--------------------+
|  INPUT SAMPLE   |   OUTPUT SAMPLE    |
+-----------------+--------------------+
| 2.0 4.0 7.5 8.0 | Media: 5.4         |
| 6.4             | Aluno em exame.    |
|                 | Nota do exame: 6.4 |
|                 | Aluno aprovado.    |
|                 | Media final: 5.9   |
+-----------------+--------------------+
| 2.0 6.5 4.0 9.0 | Media: 4.8         |
|                 | Aluno reprovado.   |
+-----------------+--------------------+
| 9.0 4.0 8.5 9.0 | Media: 7.3         |
|                 | Aluno aprovado.    |
+-----------------+--------------------+
*********************************************************************/

package main

import "fmt"

func main() {
	var N1, N2, N3, N4 float64
	fmt.Scan(&N1, &N2, &N3, &N4)
	average := (N1*2 + N2*3 + N3*4 + N4*1) / (2 + 3 + 4 + 1)
	fmt.Printf("Media: %.1f\n", average)
	if average >= 7 {
		fmt.Println("Aluno aprovado.")
	} else if average < 5 {
		fmt.Println("Aluno reprovado.")
	} else {
		fmt.Println("Aluno em exame.")
		var s float64
		fmt.Scan(&s)
		fmt.Printf("Nota do exame: %.1f\n", s)
		finalAverage := (average + s) / 2
		if finalAverage >= 5 {
			fmt.Println("Aluno aprovado.")
		} else if finalAverage <= 4.9 {
			fmt.Println("Aluno reprovado.")
		}
		fmt.Printf("Media final: %.1f\n", finalAverage)
	}
}
