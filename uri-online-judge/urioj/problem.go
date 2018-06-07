package urioj

import (
	"strconv"

	"github.com/PuerkitoBio/goquery"
	"net/http"
	"errors"
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
	res, _ := http.Get(getDescriptionUrl(p.Id))
	if res.StatusCode == 404 {
		err = errors.New("The problem does not exists ")
	} else {
		p.doc, err = goquery.NewDocumentFromReader(res.Body)
	}
	return p, err
}

func getUrl(id string) string {
	return "https://www.urionlinejudge.com.br/judge/en/problems/view/" + id
}

func getDescriptionUrl(id string) string {
	return "https://www.urionlinejudge.com.br/repository/UOJ_" + id + "_en.html"
}
