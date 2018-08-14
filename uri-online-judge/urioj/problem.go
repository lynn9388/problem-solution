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
	"net/http"
	"net/url"
	"strconv"
	"time"

	"log"

	"github.com/PuerkitoBio/goquery"
	"github.com/tdewolff/minify"
	"github.com/tdewolff/minify/html"
)

const (
	nameSelector        = "div.header > h1"
	descriptionSelector = "div.description"
	inputSelector       = "div.input"
	outputSelector      = "div.output"
	tableSelector       = "table"
	sampleSelector      = "div.problem > " + tableSelector
)

type Problem struct {
	Id  int
	Url string
	Doc *goquery.Document
}

func NewProblem(id int) (*Problem, error) {
	p := Problem{Id: id, Url: GetURL(id)}

	proxyUrl, _ := url.Parse("socks5://localhost:1080")
	tr := &http.Transport{Proxy: http.ProxyURL(proxyUrl)}
	client := &http.Client{Transport: tr, Timeout: 5 * time.Second}
	res, err := client.Get(GetDescriptionUrl(p.Id))

	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	p.Doc, err = goquery.NewDocumentFromReader(res.Body)
	return &p, err
}

func GetURL(id int) string {
	return "https://www.urionlinejudge.com.br/judge/en/problems/view/" + strconv.Itoa(id)
}

func GetDescriptionUrl(id int) string {
	return "https://www.urionlinejudge.com.br/repository/UOJ_" + strconv.Itoa(id) + "_en.html"
}

func GetDocument(rawurl string) (*goquery.Document, error) {
	proxyURL, _ := url.Parse("socks5://localhost:1080")
	tr := &http.Transport{Proxy: http.ProxyURL(proxyURL)}
	client := &http.Client{Transport: tr, Timeout: 5 * time.Second}
	res, err := client.Get(rawurl)

	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	m := minify.New()
	m.Add("text/html", &html.Minifier{
		KeepConditionalComments: true,
		KeepDefaultAttrVals:     true,
		KeepDocumentTags:        true,
		KeepEndTags:             true,
		KeepWhitespace:          true,
	})
	mr := m.Reader("text/html", res.Body)

	return goquery.NewDocumentFromReader(mr)
}

func FindContent(d *goquery.Document, selector string) *goquery.Selection {
	return d.Find(selector)
}

func FindWholeTable(firstTableNode *goquery.Selection) *goquery.Selection {
	if !firstTableNode.Is(tableSelector) {
		h, _ := firstTableNode.Html()
		log.Fatalf("not find talbe from a table node:%v", h)
	}

	table := firstTableNode.First()
	c := firstTableNode.Parent().Children()

	for i := c.IndexOfSelection(firstTableNode) + 1; i < len(c.Nodes); i++ {
		n := c.Eq(i)
		if !n.Is(tableSelector) {
			break
		}
		table = table.AddSelection(n)
	}
	return table
}
