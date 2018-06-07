/*
 * Copyright 2018 Lynn
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package urioj

import (
	"reflect"
	"testing"
)

type problem struct {
	p           *Problem
	Name        string
	Description []string
	Input       []string
	Output      []string
	Samples     []Sample
	Images      []string
}

var tests = make([]problem, 0)

func init() {
	p, _ := NewProblem(1001)
	tests = append(tests, problem{
		p,
		"Extremely Basic",
		[]string{
			"Read 2 variables, named A and B and make the sum of these two variables, assigning its result to the variable X. Print X as shown below. Print endline after the result otherwise you will get “Presentation Error”.",
		},
		[]string{
			"The input file will contain 2 integer numbers.",
		},
		[]string{
			"Print the letter X (uppercase) with a blank space before and after the equal signal followed by the value of X, according to the following example.",
			"",
			"Obs.: don't forget the endline after all.",
		},
		[]Sample{
			{[]string{"10", "9"}, []string{"X = 19"}},
			{[]string{"-10", "4"}, []string{"X = -6"}},
			{[]string{"15", "-7"}, []string{"X = 8"}},
		},
		[]string{},
	})

	p, _ = NewProblem(1015)
	tests = append(tests, problem{
		p,
		"Distance Between Two Points",
		[]string{
			"Read the four values corresponding to the x and y axes of two points in the plane, p1 (x1, y1) and p2 (x2, y2) and calculate the distance between them, showing four decimal places after the comma, according to the formula:",
			"",
			"Distance = <img src=\"https://urionlinejudge.r.worldssl.net/gallery/images/problems/UOJ_1015.png\">",
		},
		[]string{
			"The input file contains two lines of data. The first one contains two double values: x1 y1 and the second one also contains two double values with one digit after the decimal point: x2 y2.",
		},
		[]string{
			"Calculate and print the distance value using the provided formula, with 4 digits after the decimal point.",
		},
		[]Sample{
			{[]string{"1.0 7.0", "5.0 9.0"}, []string{"4.4721"}},
			{[]string{"-2.5 0.4", "12.1 7.3"}, []string{"16.1484"}},
			{[]string{"2.5 -0.4", "-12.2 7.0"}, []string{"16.4575"}},
		},
		[]string{
			"https://urionlinejudge.r.worldssl.net/gallery/images/problems/UOJ_1015.png",
		},
	})

	p, _ = NewProblem(1021)
	tests = append(tests, problem{
		p,
		"Banknotes and Coins",
		[]string{
			"Read a value of floating point with two decimal places. This represents a monetary value. After this, calculate the smallest possible number of notes and coins on which the value can be decomposed. The considered notes are of 100, 50, 20, 10, 5, 2. The possible coins are of 1, 0.50, 0.25, 0.10, 0.05 and 0.01. Print the message “NOTAS:” followed by the list of notes and the message “MOEDAS:” followed by the list of coins.",
		},
		[]string{
			"The input file contains a value of floating point N (0 ≤ N ≤ 1000000.00).",
		},
		[]string{
			"Print the minimum quantity of banknotes and coins necessary to change the initial value, as the given example.",
		},
		[]Sample{
			{[]string{"576.73"}, []string{"NOTAS:", "5 nota(s) de R$ 100.00", "1 nota(s) de R$ 50.00", "1 nota(s) de R$ 20.00", "0 nota(s) de R$ 10.00", "1 nota(s) de R$ 5.00", "0 nota(s) de R$ 2.00", "MOEDAS:", "1 moeda(s) de R$ 1.00", "1 moeda(s) de R$ 0.50", "0 moeda(s) de R$ 0.25", "2 moeda(s) de R$ 0.10", "0 moeda(s) de R$ 0.05", "3 moeda(s) de R$ 0.01"}},
			{[]string{"4.00"}, []string{"NOTAS:", "0 nota(s) de R$ 100.00", "0 nota(s) de R$ 50.00", "0 nota(s) de R$ 20.00", "0 nota(s) de R$ 10.00", "0 nota(s) de R$ 5.00", "2 nota(s) de R$ 2.00", "MOEDAS:", "0 moeda(s) de R$ 1.00", "0 moeda(s) de R$ 0.50", "0 moeda(s) de R$ 0.25", "0 moeda(s) de R$ 0.10", "0 moeda(s) de R$ 0.05", "0 moeda(s) de R$ 0.01"}},
			{[]string{"91.01"}, []string{"NOTAS:", "0 nota(s) de R$ 100.00", "1 nota(s) de R$ 50.00", "2 nota(s) de R$ 20.00", "0 nota(s) de R$ 10.00", "0 nota(s) de R$ 5.00", "0 nota(s) de R$ 2.00", "MOEDAS:", "1 moeda(s) de R$ 1.00", "0 moeda(s) de R$ 0.50", "0 moeda(s) de R$ 0.25", "0 moeda(s) de R$ 0.10", "0 moeda(s) de R$ 0.05", "1 moeda(s) de R$ 0.01"}},
		},
		[]string{},
	})

	p, _ = NewProblem(1023)
	tests = append(tests, problem{
		p,
		"Drought",
		[]string{
			"Due to the continuous drought that happened recently in some regions of Brazil, the Federal Government created an agency to assess the consumption of these regions in order to verify the behavior of the population at the time of rationing. This agency will take some cities (for sampling) and it will verify the consumption of each of the townspeople and the average consumption per inhabitant of each town.",
		},
		[]string{
			"The input contains a number of test's cases. The first line of each case of test contains an integer N (1 ≤ N ≤ 1 * 10^6), indicating the amount of properties. The following N lines contains a pair of values X (1 ≤ X ≤ 10) and Y (1 ≤ Y ≤ 200) indicating the number of residents of each property and its total consumption (m^3). Surely, no residence consumes more than 200 m^3 per month. The input's end is represented by zero.",
		},
		[]string{
			"For each case of test you must present the message “Cidade# n:”, where n is the number of the city in the sequence (1, 2, 3, ...), and then you must list in ascending order of consumption, the people's amount followed by a hyphen and the consumption of these people, rounding the value down. In the third line of output you should present the consumption per person in that town, with two decimal places without rounding, considering the total real consumption. Print a blank line between two consecutives test's cases. There is no blank line at the end of output.",
		},
		[]Sample{
			{[]string{"3", "3 22", "2 11", "3 39", "5", "1 25", "2 20", "3 31", "2 40", "6 70", "0"}, []string{"Cidade# 1:", "2-5 3-7 3-13", "Consumo medio: 9.00 m3.", "", "Cidade# 2:", "5-10 6-11 2-20 1-25", "Consumo medio: 13.28 m3."}},
		},
		[]string{},
	})

	p, _ = NewProblem(1239)
	tests = append(tests, problem{
		p,
		"Bloggo Shortcuts",
		[]string{
			"You are helping to develop a weblog-management system called bloggo. Although bloggo pushes all content to the front end of a website in HTML, not all content authors enjoy using HTML tags in their text. To make their lives easier, bloggo offers a simple syntax called shortcuts to achieve some HTML textual effects. Your job is to take a document written with shortcuts and translate it into proper HTML.",
			"",
			"One shortcut is used to make italicized text. HTML does this with the <i> and </i> tags, but in bloggo, an author can simply enclose a piece of text using two instances of the underscore character, '_'. Thus, where a content author writes",
			"",
			"  You _should_ see the baby elephant at the zoo!",
			"",
			"bloggo will publish the following instead.",
			"",
			"  You <i>should</i> see the baby elephant at the zoo!",
			"",
			"Another shortcut serves to render text in boldface, which HTML accomplishes with <b> and </b> tags. Bloggo lets content authors do the same with paired instances of the asterisk character, '*'. When a content author writes the text",
			"",
			"  Move it from *Receiving* to *Accounts Payable*.",
			"",
			"it will end up on the website as",
			"",
			"  Move it from <b>Receiving</b> to <b>Accounts Payable</b>.",
		},
		[]string{
			"The input contains several test cases. Each test case is composed by one line that contais a string text, containing zero or more usages of the italic and boldface shortcuts. Each text is between 1 and 50 characters long, inclusive. The only characters allowed in text are the alphabetic characters 'a' to 'z' and 'A' to 'Z', the underscore '_', the asterisk '*', the space character, and the punctuation symbols ',', ';', '.', '!', '?', '-', '(', and ')'. The underscore '_' occurs in text an even number of times. The asterisk '*' occurs in text an even number of times. No substring of text enclosed by a balanced pair of underscores or by a balanced pair of asterisks may contain any further underscores or asterisks.",
			"The end of input is determined by EOF.",
		},
		[]string{
			"Translate each input text into HTML as demonstrated by the examples above (and below). To render a span of text in italics in HTML, you must start with the <i> tag and end with the </i> tag. For boldface text, start with <b> and end with </b>. Print one translated text per line at standard output.",
		},
		[]Sample{
			{[]string{"You _should_ see the new walrus at the zoo!", "Move it from *Accounts Payable* to *Receiving*.", "I saw _Chelydra serpentina_ in *Centennial Park*.", "_ _ __ _ yabba dabba _ * dooooo * ****", "_now_I_know_*my*_ABC_next_time_*sing*it_with_me"}, []string{"You <i>should</i> see the new walrus at the zoo!", "Move it from <b>Accounts Payable</b> to <b>Receiving</b>.", "I saw <i>Chelydra serpentina</i> in <b>Centennial Park</b>.", "<i> </i> <i></i> <i> yabba dabba </i> <b> dooooo </b> <b></b><b></b>", "<i>now</i>I<i>know</i><b>my</b><i>ABC</i>next<i>time</i><b>sing</b>it<i>with</i>me"}},
		},
		[]string{},
	})
}

func getField(p *problem, field string) reflect.Value {
	v := reflect.ValueOf(p)
	return reflect.Indirect(v).FieldByName(field)
}

func checkStringSlice(expect []string, get []string) bool {
	if len(expect) != len(get) {
		return false
	}
	for i, s := range expect {
		if s != get[i] {
			return false
		}
	}
	return true
}

func check(id string, expect interface{}, get interface{}, t *testing.T) {
	expectType := reflect.TypeOf(expect)
	getType := reflect.TypeOf(get)
	if expectType != getType {
		t.Logf("The check value type doesn't match (problem %s):\nExpect:%v\nGet   :%v", id, expectType, getType)
		t.Fail()
	}

	match := true
	switch expect.(type) {
	case []string:
		match = checkStringSlice(expect.([]string), get.([]string))
	case []Sample:
		es := expect.([]Sample)
		gs := get.([]Sample)
		if len(es) != len(gs) {
			match = false
			break
		}
		for i, s := range es {
			match = checkStringSlice(s.Input, gs[i].Input) && checkStringSlice(s.Output, gs[i].Output)
			if !match {
				break
			}
		}
	default:
		match = reflect.DeepEqual(expect, get)
	}

	if !match {
		t.Logf("Check value doesn't match (problem %s):\nExpect:%v\nGet   :%v", id, expect, get)
		t.Fail()
	}
}

func testString(f func(*Problem) string, field string, t *testing.T) {
	for _, p := range tests {
		expect := getField(&p, field).String()
		get := f(p.p)
		check(p.p.Id, expect, get, t)
	}
}

func testStringSlice(f func(*Problem) []string, field string, t *testing.T) {
	for _, p := range tests {
		expect := getField(&p, field).Interface().([]string)
		get := f(p.p)
		check(p.p.Id, expect, get, t)
	}
}

func TestProblem_GetName(t *testing.T) {
	testString((*Problem).Name, "Name", t)
}

func TestProblem_GetDescription(t *testing.T) {
	testStringSlice((*Problem).Description, "Description", t)
}

func TestProblem_GetInput(t *testing.T) {
	testStringSlice((*Problem).Input, "Input", t)
}

func TestProblem_GetOutput(t *testing.T) {
	testStringSlice((*Problem).Output, "Output", t)
}

func TestProblem_GetSamples(t *testing.T) {
	for _, p := range tests {
		expect := p.Samples
		get := p.p.Samples()
		check(p.p.Id, expect, get, t)
	}
}

func TestProblem_Images(t *testing.T) {
	testStringSlice((*Problem).Images, "Images", t)
}
