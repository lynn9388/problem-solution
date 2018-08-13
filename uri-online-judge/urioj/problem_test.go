package urioj

import (
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
	if getURL(1001) != url {
		t.FailNow()
	}
}

func TestGetDescriptionUrl(t *testing.T) {
	url := "https://www.urionlinejudge.com.br/repository/UOJ_1001_en.html"
	if getDescriptionUrl(1001) != url {
		t.FailNow()
	}
}
