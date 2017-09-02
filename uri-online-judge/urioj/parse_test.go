package urioj

import (
	"reflect"
	"testing"
)

type problem struct {
	p           *Problem
	name        string
	description string
	input       string
	output      string
	samples     map[string]string
}

var tests = make([]problem, 2)

func init() {
	p, _ := NewProblem(1001)
	tests[0] = problem{
		p,
		"Extremely Basic",
		"Read 2 integer values and store them in variables, named A and B and make the sum of these two variables, assigning its result to the variable X. Print X as shown below. Don't present any message beyond what is being specified and don't forget to print the end of line after the result, otherwise you will receive “Presentation Error”.",
		"The input file contain 2 integer values.",
		"Print the variable X according to the following example, with a blank space before and after the equal signal. 'X' is uppercase and you have to print a blank space before and after the '=' signal.",
		map[string]string{
			"10\n9":  "X = 19",
			"-10\n4": "X = -6",
			"15\n-7": "X = 8",
		},
	}
	p, _ = NewProblem(1023)
	tests[1] = problem{
		p,
		"Drought",
		"Due to the continuous drought that happened recently in some regions of Brazil, the Federal Government created an agency to assess the consumption of these regions in order to verify the behavior of the population at the time of rationing. This agency will take some cities (for sampling) and it will verify the consumption of each of the townspeople and the average consumption per inhabitant of each town.",
		"The input contains a number of test's cases. The first line of each case of test contains an integer N (1 ≤ N ≤ 1 * 10 6), indicating the amount of properties. The following N lines contains a pair of values X (1 ≤ X ≤ 10) and Y (	1 ≤ Y ≤ 200) indicating the number of residents of each property and its total consumption (m3). Surely, no residence consumes more than 200 m3 per month. The input's end is represented by zero.",
		"For each case of test you must present the message “Cidade# n:”, where n is the number of the city in the sequence (1, 2, 3, ...), and then you must list in ascending order of consumption, the people's amount followed by a hyphen and the consumption of these people, rounding the value down. In the third line of output you should present the consumption per person in that town, with two decimal places without rounding, considering the total real consumption. Print a blank line between two consecutives test's cases. There is no blank line at the end of output.",
		map[string]string{
			"3\n3 22\n2 11\n3 39\n5\n1 25\n2 20\n3 31\n2 40\n6 70\n0": "Cidade# 1:\n2-5 3-7 3-13\nConsumo medio: 9.00 m3.\n\nCidade# 2:\n5-10 6-11 2-20 1-25\nConsumo medio: 13.28 m3.",
		},
	}

}

func getField(p *problem, field string) string {
	v := reflect.ValueOf(p)
	f := reflect.Indirect(v).FieldByName(field)
	return f.String()
}

func check(expect string, get string, t *testing.T) {
	if  expect != get {
		t.Logf("\nExpect:%s\nGet:%s\n", expect, get)
		t.Fail()
	}
}

func test(f func(*Problem) string, field string, t *testing.T) {
	for _, p := range tests {
		expect := getField(&p, field)
		get := f(p.p)
		check(expect, get, t)
	}
}

func TestProblem_GetName(t *testing.T) {
	test((*Problem).GetName, "name", t)
}

func TestProblem_GetDescription(t *testing.T) {
	test((*Problem).GetDescription, "description", t)
}

func TestProblem_GetInput(t *testing.T) {
	test((*Problem).GetInput, "input", t)
}

func TestProblem_GetOutput(t *testing.T) {
	test((*Problem).GetOutput, "output", t)
}

func TestProblem_GetSamples(t *testing.T) {
	for _, p := range tests {
		samples := p.p.GetSamples()
		for k, v := range samples {
			check(p.samples[k], v, t)
		}
	}
}
