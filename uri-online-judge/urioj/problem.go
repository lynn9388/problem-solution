package urioj

import (
	"github.com/PuerkitoBio/goquery"
	"strconv"
)

type Problem struct {
	Id  string
	Url string
	doc *goquery.Document
}

func NewProblem(id int) (*Problem, error) {
	p := new(Problem)
	var err error
	p.Id = strconv.Itoa(id)
	p.Url = getUrl(p.Id)
	p.doc, err = goquery.NewDocument(getDescriptionUrl(p.Id))
	return p, err
}

func getUrl(id string) string {
	return "https://www.urionlinejudge.com.br/judge/en/problems/view/" + id
}

func getDescriptionUrl(id string) string {
	return "https://www.urionlinejudge.com.br/repository/UOJ_" + id + "_en.html"
}
