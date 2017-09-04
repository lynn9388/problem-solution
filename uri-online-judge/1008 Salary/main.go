//                               Salary

// Write a program that reads an employee's number, his/her worked
// hours number in a month and the amount he received per hour. Print
// the employee's number and salary that he/she will receive at end of
// the month, with two decimal places.

// Input
//
// The input file contains 2 integer numbers and 1 value of floating
// point, representing the number, worked hours amount and the amount
// the employee receives per worked hour.

// Output
//
// Print the number and the employee's salary, according to the given
// example, with a blank space before and after the equal signal.

// +--------------+-------------------------------+
// | SAMPLE INPUT |         SAMPLE OUTPUT         |
// +--------------+-------------------------------+
// | 25           | NUMBER = 25                   |
// | 100          | SALARY = U$ 550.00            |
// | 5.50         |                               |
// +--------------+-------------------------------+
// | 1            | NUMBER = 1                    |
// | 200          | SALARY = U$ 4100.00           |
// | 20.50        |                               |
// +--------------+-------------------------------+
// | 6            | NUMBER = 6                    |
// | 145          | SALARY = U$ 2254.75           |
// | 15.55        |                               |
// +--------------+-------------------------------+

// https://www.urionlinejudge.com.br/judge/en/problems/view/1008

package main

import "fmt"

func main() {
	var number, hours int
	var rate float64
	fmt.Scanf("%d\n%d\n%f", &number, &hours, &rate)
	fmt.Printf("NUMBER = %d\nSALARY = U$ %.2f\n", number, float64(hours)*rate)
}
