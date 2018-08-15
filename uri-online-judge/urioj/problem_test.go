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

	"github.com/pkg/errors"
)

var tests []Problem = []Problem{
	{Id: 1001, Url: "https://www.urionlinejudge.com.br/judge/en/problems/view/1001", Name: "Extremely Basic",
		Description: []Content{TextContent(`Read 2 variables, named A and B and make the sum of these two variables, assigning its result to the variable X. Print X as shown below. Print endline after the result otherwise you will get “Presentation Error”.`)},
		Input:       []Content{TextContent(`The input file will contain 2 integer numbers.`)},
		Output:      []Content{TextContent(`Print the letter X (uppercase) with a blank space before and after the equal signal followed by the value of X, according to the following example.`), TextContent(`Obs.: don't forget the endline after all.`)},
		Sample: []Content{*generateTable([]string{"Samples Input", "Samples Output"},
			tableData{"10", "9"}, tableData{"X = 19"},
			tableData{"-10", "4"}, tableData{"X = -6"},
			tableData{"15", "-7"}, tableData{"X = 8"},
		)},
	},
}

func generateTable(head []string, data ...tableData) *TableContent {
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
			t.Errorf("content of %q doesn't match:\nExpect: %v\nGet: %v\n", selector, expect, get)
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
			t.Fatalf("row number of %v %q doesn't match:\nExpect: %v\nGet: %v\n", id, v.selector, v.numRow, len(table.Nodes))
		}
	}
}

func TestGetURL(t *testing.T) {
	for _, p := range tests {
		url := getURL(p.Id)
		if url != p.Url {
			t.Fatalf("url of %v doesn't match:\nExpect: %v\nGet: %v\n", p.Id, p.Url, url)
		}
	}
}

func TestGetName(t *testing.T) {
	for _, p := range tests {
		d, _ := getDocument(getDescriptionUrl(p.Id))
		name := getName(d)
		if name != p.Name {
			t.Fatalf("name of %v doesn't match:\nExpect: %v\nGet: %v\n", p.Id, p.Name, name)
		}
	}
}

func TestGetDescription(t *testing.T) {
	for _, p := range tests {
		d, _ := getDocument(getDescriptionUrl(p.Id))
		des := getDescription(d)
		if err := checkContents(p.Description, des); err != nil {
			t.Fatalf("description of %v doesn't match:\n%v", p.Id, err)
		}
	}
}

func checkContents(expect []Content, get []Content) error {
	if len(expect) != len(get) {
		return errors.New(fmt.Sprintf("length of data doesn't match:\nExpect: %v\nGet: %v\n", len(expect), len(get)))
	}

	for i := range expect {
		if !expect[i].equal(get[i]) {
			return errors.New(fmt.Sprintf("data doesn't match:\nExpect: %v\nGet: %v\n", expect[i], get[i]))
		}
	}

	return nil
}
