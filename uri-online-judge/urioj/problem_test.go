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
	"errors"
	"fmt"
	"testing"
)

var tests []Problem = []Problem{
	{Id: 1001, URL: "https://www.urionlinejudge.com.br/judge/en/problems/view/1001", Name: "Extremely Basic",
		Description: []Content{TextContent(`Read 2 variables, named A and B and make the sum of these two variables, assigning its result to the variable X. Print X as shown below. Print endline after the result otherwise you will get “Presentation Error”.`)},
		Input:       []Content{TextContent(`The input file will contain 2 integer numbers.`)},
		Output:      []Content{TextContent(`Print the letter X (uppercase) with a blank space before and after the equal signal followed by the value of X, according to the following example.`), TextContent(`Obs.: don't forget the endline after all.`)},
		Sample: []Content{*generateTable([]string{"Samples Input", "Samples Output"},
			TableData{"10", "9"}, TableData{"X = 19"},
			TableData{"-10", "4"}, TableData{"X = -6"},
			TableData{"15", "-7"}, TableData{"X = 8"},
		)},
	},
	{Id: 1015, URL: "https://www.urionlinejudge.com.br/judge/en/problems/view/1015", Name: "Distance Between Two Points",
		Description: []Content{TextContent(`Read the four values corresponding to the x and y axes of two points in the plane, p1 (x1, y1) and p2 (x2, y2) and calculate the distance between them, showing four decimal places after the comma, according to the formula:`), FileContent(`https://urionlinejudge.r.worldssl.net/gallery/images/problems/UOJ_1015.png`), TextContent(`Distance = <img src="https://urionlinejudge.r.worldssl.net/gallery/images/problems/UOJ_1015.png">`)},
		Input:       []Content{TextContent(`The input file contains two lines of data. The first one contains two double values: x1 y1 and the second one also contains two double values with one digit after the decimal point: x2 y2.`)},
		Output:      []Content{TextContent(`Calculate and print the distance value using the provided formula, with 4 digits after the decimal point.`)},
		Sample: []Content{*generateTable([]string{"Input Sample", "Output Sample"},
			TableData{"1.0 7.0", "5.0 9.0"}, TableData{"4.4721"},
			TableData{"-2.5 0.4", "12.1 7.3"}, TableData{"16.1484"},
			TableData{"2.5 -0.4", "-12.2 7.0"}, TableData{"16.4575"},
		)},
	},
	{Id: 1021, URL: "https://www.urionlinejudge.com.br/judge/en/problems/view/1021", Name: "Banknotes and Coins",
		Description: []Content{TextContent(`Read a value of floating point with two decimal places. This represents a monetary value. After this, calculate the smallest possible number of notes and coins on which the value can be decomposed. The considered notes are of 100, 50, 20, 10, 5, 2. The possible coins are of 1, 0.50, 0.25, 0.10, 0.05 and 0.01. Print the message “NOTAS:” followed by the list of notes and the message “MOEDAS:” followed by the list of coins.`)},
		Input:       []Content{TextContent(`The input file contains a value of floating point N (0 ≤ N ≤ 1000000.00).`)},
		Output:      []Content{TextContent(`Print the minimum quantity of banknotes and coins necessary to change the initial value, as the given example.`)},
		Sample: []Content{*generateTable([]string{"Input Sample", "Output Sample"},
			TableData{"576.73"}, TableData{"NOTAS:", "5 nota(s) de R$ 100.00", "1 nota(s) de R$ 50.00", "1 nota(s) de R$ 20.00", "0 nota(s) de R$ 10.00", "1 nota(s) de R$ 5.00", "0 nota(s) de R$ 2.00", "MOEDAS:", "1 moeda(s) de R$ 1.00", "1 moeda(s) de R$ 0.50", "0 moeda(s) de R$ 0.25", "2 moeda(s) de R$ 0.10", "0 moeda(s) de R$ 0.05", "3 moeda(s) de R$ 0.01"},
			TableData{"4.00"}, TableData{"NOTAS:", "0 nota(s) de R$ 100.00", "0 nota(s) de R$ 50.00", "0 nota(s) de R$ 20.00", "0 nota(s) de R$ 10.00", "0 nota(s) de R$ 5.00", "2 nota(s) de R$ 2.00", "MOEDAS:", "0 moeda(s) de R$ 1.00", "0 moeda(s) de R$ 0.50", "0 moeda(s) de R$ 0.25", "0 moeda(s) de R$ 0.10", "0 moeda(s) de R$ 0.05", "0 moeda(s) de R$ 0.01"},
			TableData{"91.01"}, TableData{"NOTAS:", "0 nota(s) de R$ 100.00", "1 nota(s) de R$ 50.00", "2 nota(s) de R$ 20.00", "0 nota(s) de R$ 10.00", "0 nota(s) de R$ 5.00", "0 nota(s) de R$ 2.00", "MOEDAS:", "1 moeda(s) de R$ 1.00", "0 moeda(s) de R$ 0.50", "0 moeda(s) de R$ 0.25", "0 moeda(s) de R$ 0.10", "0 moeda(s) de R$ 0.05", "1 moeda(s) de R$ 0.01"},
		)},
	},
	{Id: 1023, URL: "https://www.urionlinejudge.com.br/judge/en/problems/view/1023", Name: "Drought",
		Description: []Content{TextContent(`Due to the continuous drought that happened recently in some regions of Brazil, the Federal Government created an agency to assess the consumption of these regions in order to verify the behavior of the population at the time of rationing. This agency will take some cities (for sampling) and it will verify the consumption of each of the townspeople and the average consumption per inhabitant of each town.`)},
		Input:       []Content{TextContent(`The input contains a number of test's cases. The first line of each case of test contains an integer N (1 ≤ N ≤ 1 * 10 6), indicating the amount of properties. The following N lines contains a pair of values X (1 ≤ X ≤ 10) and Y ( 1 ≤ Y ≤ 200) indicating the number of residents of each property and its total consumption (m3). Surely, no residence consumes more than 200 m3 per month. The input's end is represented by zero.`)},
		Output:      []Content{TextContent(`For each case of test you must present the message “Cidade# n:”, where n is the number of the city in the sequence (1, 2, 3, ...), and then you must list in ascending order of consumption, the people's amount followed by a hyphen and the consumption of these people, rounding the value down. In the third line of output you should present the consumption per person in that town, with two decimal places without rounding, considering the total real consumption. Print a blank line between two consecutives test's cases. There is no blank line at the end of output.`)},
		Sample: []Content{*generateTable([]string{"Input Sample", "Output Sample"},
			TableData{"3", "3 22", "2 11", "3 39", "5", "1 25", "2 20", "3 31", "2 40", "6 70", "0"}, TableData{"Cidade# 1:", "2-5 3-7 3-13", "Consumo medio: 9.00 m3.", "", "Cidade# 2:", "5-10 6-11 2-20 1-25", "Consumo medio: 13.28 m3."},
		)},
	},
	{Id: 1239, URL: "https://www.urionlinejudge.com.br/judge/en/problems/view/1239", Name: "Bloggo Shortcuts",
		Description: []Content{TextContent(`You are helping to develop a weblog-management system called bloggo. Although bloggo pushes all content to the front end of a website in HTML, not all content authors enjoy using HTML tags in their text. To make their lives easier, bloggo offers a simple syntax called shortcuts to achieve some HTML textual effects. Your job is to take a document written with shortcuts and translate it into proper HTML.`), TextContent(`One shortcut is used to make italicized text. HTML does this with the <i> and </i> tags, but in bloggo, an author can simply enclose a piece of text using two instances of the underscore character, '_'. Thus, where a content author writes`), TextContent(`  You _should_ see the baby elephant at the zoo!`), TextContent(`bloggo will publish the following instead.`), TextContent(`  You <i>should</i> see the baby elephant at the zoo!`), TextContent(`Another shortcut serves to render text in boldface, which HTML accomplishes with <b> and </b> tags. Bloggo lets content authors do the same with paired instances of the asterisk character, '*'. When a content author writes the text`), TextContent(`  Move it from *Receiving* to *Accounts Payable*.`), TextContent(`it will end up on the website as`), TextContent(`  Move it from <b>Receiving</b> to <b>Accounts Payable</b>.`)},
		Input:       []Content{TextContent(`The input contains several test cases. Each test case is composed by one line that contais a string text, containing zero or more usages of the italic and boldface shortcuts. Each text is between 1 and 50 characters long, inclusive. The only characters allowed in text are the alphabetic characters 'a' to 'z' and 'A' to 'Z', the underscore '_', the asterisk '*', the space character, and the punctuation symbols ',', ';', '.', '!', '?', '-', '(', and ')'. The underscore '_' occurs in text an even number of times. The asterisk '*' occurs in text an even number of times. No substring of text enclosed by a balanced pair of underscores or by a balanced pair of asterisks may contain any further underscores or asterisks.`), TextContent(`The end of input is determined by EOF.`)},
		Output:      []Content{TextContent(`Translate each input text into HTML as demonstrated by the examples above (and below). To render a span of text in italics in HTML, you must start with the <i> tag and end with the </i> tag. For boldface text, start with <b> and end with </b>. Print one translated text per line at standard output.`)},
		Sample: []Content{*generateTable([]string{"Sample Input", "Sample Output"},
			TableData{`You _should_ see the new walrus at the zoo!`, `Move it from *Accounts Payable* to *Receiving*.`, `I saw _Chelydra serpentina_ in *Centennial Park*.`, `_ _ __ _ yabba dabba _ * dooooo * ****`, `_now_I_know_*my*_ABC_next_time_*sing*it_with_me`}, TableData{`You <i>should</i> see the new walrus at the zoo!`, `Move it from <b>Accounts Payable</b> to <b>Receiving</b>.`, `I saw <i>Chelydra serpentina</i> in <b>Centennial Park</b>.`, `<i> </i> <i></i> <i> yabba dabba </i> <b> dooooo </b> <b></b><b></b>`, `<i>now</i>I<i>know</i><b>my</b><i>ABC</i>next<i>time</i><b>sing</b>it<i>with</i>me`},
		)},
	},
}

func generateTable(head []string, data ...TableData) *TableContent {
	t, _ := newTable(head, data...)
	return t
}

func TestGetDescriptionUrl(t *testing.T) {
	url := "https://www.urionlinejudge.com.br/repository/UOJ_1001_en.html"
	if getDescriptionUrl(1001) != url {
		t.FailNow()
	}
}

func TestGetDocument(t *testing.T) {
	d, err := getDocument(getDescriptionUrl(1001))
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
	d, _ := getDocument(getDescriptionUrl(1001))
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
		d, err := getDocument(getDescriptionUrl(id))
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
		pp, err := NewProblem(p.Id)
		if err != nil {
			t.Fatalf(err.Error())
		}

		if p.URL != pp.URL {
			t.Errorf("url of %v doesn't match:\nExpect: %v\n   Get: %v\n", p.Id, p.URL, pp.URL)
		}

		if p.Name != pp.Name {
			t.Errorf("name of %v doesn't match:\nExpect: %v\n   Get: %v\n", p.Id, p.Name, pp.URL)
		}

		if err := checkContents(p.Description, pp.Description); err != nil {
			t.Errorf("description of %v doesn't match:\n%v", p.Id, err)
		}

		if err := checkContents(p.Input, pp.Input); err != nil {
			t.Errorf("input of %v doesn't match:\n%v", p.Id, err)
		}

		if err := checkContents(p.Output, pp.Output); err != nil {
			t.Errorf("output of %v doesn't match:\n%v", p.Id, err)
		}

		if err := checkContents(p.Sample, pp.Sample); err != nil {
			t.Errorf("sample of %v doesn't match:\n%v", p.Id, err)
		}
	}
}

func checkContents(expect []Content, get []Content) error {
	if len(expect) != len(get) {
		return errors.New(fmt.Sprintf("length of data doesn't match:\nExpect: %v\n   Get: %v\n", len(expect), len(get)))
	}

	for i := range expect {
		if !expect[i].Equal(get[i]) {
			return errors.New(fmt.Sprintf("data doesn't match:\nExpect: %v\n   Get: %v\n", expect[i], get[i]))
		}
	}

	return nil
}
