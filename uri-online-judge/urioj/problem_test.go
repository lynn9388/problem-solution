package urioj

import (
	"strings"
	"testing"
)

func TestNewProblem(t *testing.T) {
	p, err := NewProblem(1001)
	if err != nil {
		t.Fatal(err)
	}
	if p.Doc == nil {
		t.FailNow()
	}
	html, err := p.Doc.Html()
	if err != nil {
		t.Fatal(err)
	}
	if len(html) == 0 {
		t.FailNow()
	}
}

func TestGetURL(t *testing.T) {
	url := "https://www.urionlinejudge.com.br/judge/en/problems/view/1001"
	if GetURL(1001) != url {
		t.FailNow()
	}
}

func TestGetDescriptionUrl(t *testing.T) {
	url := "https://www.urionlinejudge.com.br/repository/UOJ_1001_en.html"
	if GetDescriptionUrl(1001) != url {
		t.FailNow()
	}
}

func TestGetDocument(t *testing.T) {
	d, err := GetDocument(GetDescriptionUrl(1001))
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

func TestFindContent(t *testing.T) {
	tests := map[string]string{
		nameSelector: "Extremely Basic",
		descriptionSelector: `<p>
Read 2 variables, named <strong>A</strong> and <strong>B</strong> and make the sum of these two variables, assigning its result to the variable <strong>X</strong>. Print <strong>X</strong> as shown below. Print endline after the result otherwise you will get “<em>Presentation Error</em>”.
</p>`,
		inputSelector: `<p>
The input file will contain 2 integer numbers.
</p>`,
		outputSelector: `<p>
Print the letter <strong>X</strong> (uppercase) with a blank space before and after the equal signal followed by the value of X, according to the following example. </p> <p> Obs.: don&#39;t forget the endline after all.
</p>`,
	}
	d, _ := GetDocument(GetDescriptionUrl(1001))
	for selector, expect := range tests {
		content := FindContent(d, selector)

		var get string
		for i := range content.Nodes {
			h, _ := content.Eq(i).Html()
			get += h
		}
		get = strings.TrimSpace(get)
		if get != expect {
			t.Errorf("content of %q doesn't match:\nExpect:\n%v\nGet:\n%v\n", selector, expect, get)
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
		d, _ := GetDocument(GetDescriptionUrl(id))
		table := FindWholeTable(FindContent(d, v.selector))
		if len(table.Nodes) != v.numRow {
			t.Errorf("row number of %v %q doesn't match:\nExpect:%v\nGet:%v\n", id, v.selector, v.numRow, len(table.Nodes))
		}
	}
}
