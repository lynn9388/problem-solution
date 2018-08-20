/*
 * Copyright © 2018 Lynn <lynn9388@gmail.com>
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
	"fmt"
	"testing"
)

var tests = []Problem{
	{ID: 1001, URL: "https://www.urionlinejudge.com.br/judge/en/problems/view/1001", Name: "Extremely Basic",
		Description: []Content{TextContent(`Read 2 variables, named A and B and make the sum of these two variables, assigning its result to the variable X. Print X as shown below. Print endline after the result otherwise you will get “Presentation Error”.`)},
		Input:       []Content{TextContent(`The input file will contain 2 integer numbers.`)},
		Output:      []Content{TextContent(`Print the letter X (uppercase) with a blank space before and after the equal signal followed by the value of X, according to the following example.`), TextContent(`Obs.: don't forget the endline after all.`)},
		Sample: []Content{*generateTable([]string{"Samples Input", "Samples Output"},
			TableData{TextContent("10\n9")}, TableData{TextContent("X = 19")},
			TableData{TextContent("-10\n4")}, TableData{TextContent("X = -6")},
			TableData{TextContent("15\n-7")}, TableData{TextContent("X = 8")})},
	},
	{ID: 1015, URL: "https://www.urionlinejudge.com.br/judge/en/problems/view/1015", Name: "Distance Between Two Points",
		Description: []Content{TextContent(`Read the four values corresponding to the x and y axes of two points in the plane, p1 (x1, y1) and p2 (x2, y2) and calculate the distance between them, showing four decimal places after the comma, according to the formula:`), FileContent(`https://urionlinejudge.r.worldssl.net/gallery/images/problems/UOJ_1015.png`), TextContent(`Distance = <img src="https://urionlinejudge.r.worldssl.net/gallery/images/problems/UOJ_1015.png">`)},
		Input:       []Content{TextContent(`The input file contains two lines of data. The first one contains two double values: x1 y1 and the second one also contains two double values with one digit after the decimal point: x2 y2.`)},
		Output:      []Content{TextContent(`Calculate and print the distance value using the provided formula, with 4 digits after the decimal point.`)},
		Sample: []Content{*generateTable([]string{"Input Sample", "Output Sample"},
			TableData{TextContent("1.0 7.0\n5.0 9.0")}, TableData{TextContent("4.4721")},
			TableData{TextContent("-2.5 0.4\n12.1 7.3")}, TableData{TextContent("16.1484")},
			TableData{TextContent("2.5 -0.4\n-12.2 7.0")}, TableData{TextContent("16.4575")})},
	},
	{ID: 1021, URL: "https://www.urionlinejudge.com.br/judge/en/problems/view/1021", Name: "Banknotes and Coins",
		Description: []Content{TextContent(`Read a value of floating point with two decimal places. This represents a monetary value. After this, calculate the smallest possible number of notes and coins on which the value can be decomposed. The considered notes are of 100, 50, 20, 10, 5, 2. The possible coins are of 1, 0.50, 0.25, 0.10, 0.05 and 0.01. Print the message “NOTAS:” followed by the list of notes and the message “MOEDAS:” followed by the list of coins.`)},
		Input:       []Content{TextContent(`The input file contains a value of floating point N (0 ≤ N ≤ 1000000.00).`)},
		Output:      []Content{TextContent(`Print the minimum quantity of banknotes and coins necessary to change the initial value, as the given example.`)},
		Sample: []Content{*generateTable([]string{"Input Sample", "Output Sample"},
			TableData{TextContent("576.73")}, TableData{TextContent("NOTAS:\n5 nota(s) de R$ 100.00\n1 nota(s) de R$ 50.00\n1 nota(s) de R$ 20.00\n0 nota(s) de R$ 10.00\n1 nota(s) de R$ 5.00\n0 nota(s) de R$ 2.00\nMOEDAS:\n1 moeda(s) de R$ 1.00\n1 moeda(s) de R$ 0.50\n0 moeda(s) de R$ 0.25\n2 moeda(s) de R$ 0.10\n0 moeda(s) de R$ 0.05\n3 moeda(s) de R$ 0.01")},
			TableData{TextContent("4.00")}, TableData{TextContent("NOTAS:\n0 nota(s) de R$ 100.00\n0 nota(s) de R$ 50.00\n0 nota(s) de R$ 20.00\n0 nota(s) de R$ 10.00\n0 nota(s) de R$ 5.00\n2 nota(s) de R$ 2.00\nMOEDAS:\n0 moeda(s) de R$ 1.00\n0 moeda(s) de R$ 0.50\n0 moeda(s) de R$ 0.25\n0 moeda(s) de R$ 0.10\n0 moeda(s) de R$ 0.05\n0 moeda(s) de R$ 0.01")},
			TableData{TextContent("91.01")}, TableData{TextContent("NOTAS:\n0 nota(s) de R$ 100.00\n1 nota(s) de R$ 50.00\n2 nota(s) de R$ 20.00\n0 nota(s) de R$ 10.00\n0 nota(s) de R$ 5.00\n0 nota(s) de R$ 2.00\nMOEDAS:\n1 moeda(s) de R$ 1.00\n0 moeda(s) de R$ 0.50\n0 moeda(s) de R$ 0.25\n0 moeda(s) de R$ 0.10\n0 moeda(s) de R$ 0.05\n1 moeda(s) de R$ 0.01")})},
	},
	{ID: 1023, URL: "https://www.urionlinejudge.com.br/judge/en/problems/view/1023", Name: "Drought",
		Description: []Content{TextContent(`Due to the continuous drought that happened recently in some regions of Brazil, the Federal Government created an agency to assess the consumption of these regions in order to verify the behavior of the population at the time of rationing. This agency will take some cities (for sampling) and it will verify the consumption of each of the townspeople and the average consumption per inhabitant of each town.`)},
		Input:       []Content{TextContent(`The input contains a number of test's cases. The first line of each case of test contains an integer N (1 ≤ N ≤ 1 * 10 ⁶), indicating the amount of properties. The following N lines contains a pair of values X (1 ≤ X ≤ 10) and Y ( 1 ≤ Y ≤ 200) indicating the number of residents of each property and its total consumption (m³). Surely, no residence consumes more than 200 m³ per month. The input's end is represented by zero.`)},
		Output:      []Content{TextContent(`For each case of test you must present the message “Cidade# n:”, where n is the number of the city in the sequence (1, 2, 3, ...), and then you must list in ascending order of consumption, the people's amount followed by a hyphen and the consumption of these people, rounding the value down. In the third line of output you should present the consumption per person in that town, with two decimal places without rounding, considering the total real consumption. Print a blank line between two consecutives test's cases. There is no blank line at the end of output.`)},
		Sample: []Content{*generateTable([]string{"Input Sample", "Output Sample"},
			TableData{TextContent("3\n3 22\n2 11\n3 39\n5\n1 25\n2 20\n3 31\n2 40\n6 70\n0")}, TableData{TextContent("Cidade# 1:\n2-5 3-7 3-13\nConsumo medio: 9.00 m3.\n\nCidade# 2:\n5-10 6-11 2-20 1-25\nConsumo medio: 13.28 m3.")})},
	},
	{ID: 1239, URL: "https://www.urionlinejudge.com.br/judge/en/problems/view/1239", Name: "Bloggo Shortcuts",
		Description: []Content{TextContent(`You are helping to develop a weblog-management system called bloggo. Although bloggo pushes all content to the front end of a website in HTML, not all content authors enjoy using HTML tags in their text. To make their lives easier, bloggo offers a simple syntax called shortcuts to achieve some HTML textual effects. Your job is to take a document written with shortcuts and translate it into proper HTML.`), TextContent(`One shortcut is used to make italicized text. HTML does this with the <i> and </i> tags, but in bloggo, an author can simply enclose a piece of text using two instances of the underscore character, '_'. Thus, where a content author writes`), TextContent(`  You _should_ see the baby elephant at the zoo!`), TextContent(`bloggo will publish the following instead.`), TextContent(`  You <i>should</i> see the baby elephant at the zoo!`), TextContent(`Another shortcut serves to render text in boldface, which HTML accomplishes with <b> and </b> tags. Bloggo lets content authors do the same with paired instances of the asterisk character, '*'. When a content author writes the text`), TextContent(`  Move it from *Receiving* to *Accounts Payable*.`), TextContent(`it will end up on the website as`), TextContent(`  Move it from <b>Receiving</b> to <b>Accounts Payable</b>.`)},
		Input:       []Content{TextContent(`The input contains several test cases. Each test case is composed by one line that contais a string text, containing zero or more usages of the italic and boldface shortcuts. Each text is between 1 and 50 characters long, inclusive. The only characters allowed in text are the alphabetic characters 'a' to 'z' and 'A' to 'Z', the underscore '_', the asterisk '*', the space character, and the punctuation symbols ',', ';', '.', '!', '?', '-', '(', and ')'. The underscore '_' occurs in text an even number of times. The asterisk '*' occurs in text an even number of times. No substring of text enclosed by a balanced pair of underscores or by a balanced pair of asterisks may contain any further underscores or asterisks.\n\nThe end of input is determined by EOF.`)},
		Output:      []Content{TextContent(`Translate each input text into HTML as demonstrated by the examples above (and below). To render a span of text in italics in HTML, you must start with the <i> tag and end with the </i> tag. For boldface text, start with <b> and end with </b>. Print one translated text per line at standard output.`)},
		Sample: []Content{*generateTable([]string{"Sample Input", "Sample Output"},
			TableData{TextContent(`You _should_ see the new walrus at the zoo!\nMove it from *Accounts Payable* to *Receiving*.\nI saw _Chelydra serpentina_ in *Centennial Park*.\n_ _ __ _ yabba dabba _ * dooooo * ****\n_now_I_know_*my*_ABC_next_time_*sing*it_with_me`)}, TableData{TextContent(`You <i>should</i> see the new walrus at the zoo!\nMove it from <b>Accounts Payable</b> to <b>Receiving</b>.\nI saw <i>Chelydra serpentina</i> in <b>Centennial Park</b>.\n<i> </i> <i></i> <i> yabba dabba </i> <b> dooooo </b> <b></b><b></b>\n<i>now</i>I<i>know</i><b>my</b><i>ABC</i>next<i>time</i><b>sing</b>it<i>with</i>me`)})},
	},
}

func generateTable(head []string, data ...TableData) *TableContent {
	t, _ := newTable(head, data...)
	return t
}

func TestGetDocument(t *testing.T) {
	d, err := getDocument(1001)
	if err != nil {
		t.Fatal(err)
	}
	if d == nil {
		t.FailNow()
	}
	html, err := d.Html()
	if err != nil {
		t.Fatal(err)
	}
	if len(html) == 0 {
		t.FailNow()
	}
}

func TestSelector(t *testing.T) {
	tests := map[string]string{
		nameSelector:        `<h1>Extremely Basic</h1>`,
		descriptionSelector: `<div class="description"><p>Read 2 variables, named <strong>A</strong> and <strong>B</strong> and make the sum of these two variables, assigning its result to the variable <strong>X</strong>. Print <strong>X</strong> as shown below. Print endline after the result otherwise you will get “<em>Presentation Error</em>”.</p></div>`,
		inputSelector:       `<div class="input"><p>The input file will contain 2 integer numbers.</p></div>`,
		outputSelector:      `<div class="output"><p>Print the letter <strong>X</strong> (uppercase) with a blank space before and after the equal signal followed by the value of X, according to the following example.</p><p>Obs.: don&#39;t forget the endline after all.</p></div>`,
	}
	d, _ := getDocument(1001)
	for selector, expect := range tests {
		get := getHTML(d.Find(selector))
		if get != expect {
			t.Errorf("content of %q doesn't match:\nExpect: %v\n   Get: %v\n", selector, expect, get)
		}
	}
}

func TestFindWholeTable(t *testing.T) {
	tests := map[int]struct {
		selector string
		numRow   int
	}{
		1001: {selector: sampleSelector, numRow: 3},
		1048: {selector: "div.description table", numRow: 1},
	}

	for id, v := range tests {
		d, err := getDocument(id)
		if err != nil {
			t.Fatal(err)
		}
		table, err := findWholeTable(d.Find(v.selector).First())
		if err != nil {
			t.Fatal(err)
		}
		if len(table.Nodes) != v.numRow {
			t.Fatalf("row number of %v %q doesn't match:\nExpect: %v\n   Get: %v\n", id, v.selector, v.numRow, len(table.Nodes))
		}
	}
}

func TestGetContent(t *testing.T) {
	for _, p := range tests {
		pp, err := NewProblem(p.ID)
		if err != nil {
			t.Fatalf(err.Error())
		}

		if p.URL != pp.URL {
			t.Errorf("url of %v doesn't match:\nExpect: %v\n   Get: %v\n", p.ID, p.URL, pp.URL)
		}

		if p.Name != pp.Name {
			t.Errorf("name of %v doesn't match:\nExpect: %v\n   Get: %v\n", p.ID, p.Name, pp.URL)
		}

		if err := checkContents(p.Description, pp.Description); err != nil {
			t.Errorf("description of %v doesn't match:\n%v", p.ID, err)
		}

		if err := checkContents(p.Input, pp.Input); err != nil {
			t.Errorf("input of %v doesn't match:\n%v", p.ID, err)
		}

		if err := checkContents(p.Output, pp.Output); err != nil {
			t.Errorf("output of %v doesn't match:\n%v", p.ID, err)
		}

		if err := checkContents(p.Sample, pp.Sample); err != nil {
			t.Errorf("sample of %v doesn't match:\n%v", p.ID, err)
		}
	}
}

func checkContents(expect []Content, get []Content) error {
	if len(expect) != len(get) {
		return fmt.Errorf("length of data doesn't match:\nExpect: %v\n   Get: %v", len(expect), len(get))
	}

	for i := range expect {
		if !expect[i].equal(get[i]) {
			return fmt.Errorf("data doesn't match:\nExpect: %v\n   Get: %v", expect[i], get[i])
		}
	}

	return nil
}
