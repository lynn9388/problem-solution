# problem-solution
Solutions for programming puzzles.

## urioj

[![GoDoc](https://godoc.org/github.com/lynn9388/problem-solution/urioj?status.svg)](https://godoc.org/github.com/lynn9388/problem-solution/urioj)

Package urioj parses html for problem page from [URI Online Judge](https://www.urionlinejudge.com.br/). It can format html page to plain text and download relative files.

For example ([Problem 1048](https://www.urionlinejudge.com.br/judge/en/problems/view/1045)):

```go
d, _ := NewDescription(1045)
fmt.Println(d)
```

Output:

```text
/*********************************************************************
                            Triangle Types
    https://www.urionlinejudge.com.br/judge/en/problems/view/1045

Read 3 double numbers (A, B and C) representing the sides of a triangle
and arrange them in decreasing order, so that the side A is the biggest
of the three sides. Next, determine the type of triangle that they can
make, based on the following cases always writing an appropriate message:

 • if A ≥ B + C, write the message: NAO FORMA TRIANGULO
 • if A² = B² + C², write the message: TRIANGULO RETANGULO
 • if A² > B² + C², write the message: TRIANGULO OBTUSANGULO
 • if A² < B² + C², write the message: TRIANGULO ACUTANGULO
 • if the three sides are the same size, write the message: TRIANGULO
   EQUILATERO
 • if only two sides are the same and the third one is different,
   write the message: TRIANGULO ISOSCELES

Input
*****
The input contains three double numbers, A (0 < A) , B (0 < B) and C
(0 < C).

Output
******
Print all the classifications of the triangle presented in the input.

+---------------+-----------------------+
| INPUT SAMPLES |    OUTPUT SAMPLES     |
+---------------+-----------------------+
| 7.0 5.0 7.0   | TRIANGULO ACUTANGULO  |
|               | TRIANGULO ISOSCELES   |
+---------------+-----------------------+
| 6.0 6.0 10.0  | TRIANGULO OBTUSANGULO |
|               | TRIANGULO ISOSCELES   |
+---------------+-----------------------+
| 6.0 6.0 6.0   | TRIANGULO ACUTANGULO  |
|               | TRIANGULO EQUILATERO  |
+---------------+-----------------------+
| 5.0 7.0 2.0   | NAO FORMA TRIANGULO   |
+---------------+-----------------------+
| 6.0 8.0 10.0  | TRIANGULO RETANGULO   |
+---------------+-----------------------+
*********************************************************************/
```


