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
	Samples     map[string]string
}

var tests = make([]problem, 2)

func init() {
	p, _ := NewProblem(1001)
	tests[0] = problem{
		p,
		"Extremely Basic",
		[]string{
			"Read 2 integer values and store them in variables, named A and B and make the sum of these two variables, assigning its result to the variable X. Print X as shown below. Don't present any message beyond what is being specified and don't forget to print the end of line after the result, otherwise you will receive “Presentation Error”.",
		},
		[]string{
			"The input file contain 2 integer values.",
		},
		[]string{
			"Print the variable X according to the following example, with a blank space before and after the equal signal. 'X' is uppercase and you have to print a blank space before and after the '=' signal.",
		},
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
		[]string{
			"Due to the continuous drought that happened recently in some regions of Brazil, the Federal Government created an agency to assess the consumption of these regions in order to verify the behavior of the population at the time of rationing. This agency will take some cities (for sampling) and it will verify the consumption of each of the townspeople and the average consumption per inhabitant of each town.",
		},
		[]string{
			"The input contains a number of test's cases. The first line of each case of test contains an integer N (1 ≤ N ≤ 1 * 10^6), indicating the amount of properties. The following N lines contains a pair of values X (1 ≤ X ≤ 10) and Y (1 ≤ Y ≤ 200) indicating the number of residents of each property and its total consumption (m^3). Surely, no residence consumes more than 200 m^3 per month. The input's end is represented by zero.",
		},
		[]string{
			"For each case of test you must present the message “Cidade# n:”, where n is the number of the city in the sequence (1, 2, 3, ...), and then you must list in ascending order of consumption, the people's amount followed by a hyphen and the consumption of these people, rounding the value down. In the third line of output you should present the consumption per person in that town, with two decimal places without rounding, considering the total real consumption. Print a blank line between two consecutives test's cases. There is no blank line at the end of output.",
		},
		map[string]string{
			"3\n3 22\n2 11\n3 39\n5\n1 25\n2 20\n3 31\n2 40\n6 70\n0": "Cidade# 1:\n2-5 3-7 3-13\nConsumo medio: 9.00 m3.\n\nCidade# 2:\n5-10 6-11 2-20 1-25\nConsumo medio: 13.28 m3.",
		},
	}
	p, _ = NewProblem(1239)
	tests[1] = problem{
		p,
		"Bloggo Shortcuts",
		[]string{
			"You are helping to develop a weblog-management system called bloggo. Although bloggo pushes all content to the front end of a website in HTML, not all content authors enjoy using HTML tags in their text. To make their lives easier, bloggo offers a simple syntax called shortcuts to achieve some HTML textual effects. Your job is to take a document written with shortcuts and translate it into proper HTML.",
			"One shortcut is used to make italicized text. HTML does this with the <i> and </i> tags, but in bloggo, an author can simply enclose a piece of text using two instances of the underscore character, '_'. Thus, where a content author writes",
			"  You _should_ see the baby elephant at the zoo!",
			"bloggo will publish the following instead.",
			"  You <i>should</i> see the baby elephant at the zoo!",
			"Another shortcut serves to render text in boldface, which HTML accomplishes with <b> and </b> tags. Bloggo lets content authors do the same with paired instances of the asterisk character, '*'. When a content author writes the text",
			"  Move it from *Receiving* to *Accounts Payable*.",
			"it will end up on the website as",
			"  Move it from <b>Receiving</b> to <b>Accounts Payable</b>.",
		},
		[]string{
			"The input contains several test cases. Each test case is composed by one line that contais a string text, containing zero or more usages of the italic and boldface shortcuts. Each text is between 1 and 50 characters long, inclusive. The only characters allowed in text are the alphabetic characters 'a' to 'z' and 'A' to 'Z', the underscore '_', the asterisk '*', the space character, and the punctuation symbols ',', ';', '.', '!', '?', '-', '(', and ')'. The underscore '_' occurs in text an even number of times. The asterisk '*' occurs in text an even number of times. No substring of text enclosed by a balanced pair of underscores or by a balanced pair of asterisks may contain any further underscores or asterisks.",
			"The end of input is determined by EOF.",
		},
		[]string{
			"Translate each input text into HTML as demonstrated by the examples above (and below). To render a span of text in italics in HTML, you must start with the <i> tag and end with the </i> tag. For boldface text, start with <b> and end with </b>. Print one translated text per line at standard output.",
		},
		map[string]string{
			"You _should_ see the new walrus at the zoo!\nMove it from *Accounts Payable* to *Receiving*.\nI saw _Chelydra serpentina_ in *Centennial Park*.\n_ _ __ _ yabba dabba _ * dooooo * ****\n_now_I_know_*my*_ABC_next_time_*sing*it_with_me":"You <i>should</i> see the new walrus at the zoo!\nMove it from <b>Accounts Payable</b> to <b>Receiving</b>.\nI saw <i>Chelydra serpentina</i> in <b>Centennial Park</b>.\n<i> </i> <i></i> <i> yabba dabba </i> <b> dooooo </b> <b></b><b></b>\n<i>now</i>I<i>know</i><b>my</b><i>ABC</i>next<i>time</i><b>sing</b>it<i>with</i>me",
		},
	}
}

func getField(p *problem, field string) reflect.Value {
	v := reflect.ValueOf(p)
	return reflect.Indirect(v).FieldByName(field)
}

func check(expect interface{}, get interface{}, t *testing.T) {
	expectType := reflect.TypeOf(expect)
	getType := reflect.TypeOf(get)
	if expectType != getType {
		t.Logf("The check value type doesn't match:\nExpect:%v\nGet   :%v", expectType, getType)
		t.Fail()
	}

	match := true
	switch expect.(type) {
	case []string:
		e := expect.([]string)
		for i, s := range e {
			g := get.([]string)
			if s != g[i] {
				match = false
				break
			}
		}
	default:
		match = reflect.DeepEqual(expect, get)
	}

	if !match {
		t.Logf("Check value doesn't match:\nExpect:%v\nGet   :%v", expect, get)
		t.Fail()
	}
}

func testString(f func(*Problem) string, field string, t *testing.T) {
	for _, p := range tests {
		expect := getField(&p, field).String()
		get := f(p.p)
		check(expect, get, t)
	}
}

func testStringSlice(f func(*Problem) []string, field string, t *testing.T) {
	for _, p := range tests {
		expect := getField(&p, field).Interface().([]string)
		get := f(p.p)
		check(expect, get, t)
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
		samples := p.p.Samples()
		for k, v := range samples {
			check(p.Samples[k], v, t)
		}
	}
}
