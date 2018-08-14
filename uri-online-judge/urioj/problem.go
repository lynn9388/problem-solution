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

type Content interface{}
type TextContent string
type FileContent string
type TableContent struct {
	head []string
	data [][]string
}

type Problem struct {
	Id          int
	Url         string
	Name        string
	Description []Content
	Input       []Content
	Output      []Content
	Sample      []Content
}

func NewProblem(id int) (*Problem, error) {
	p := Problem{Id: id, Url: getURL(id)}

	d, err := getDocument(getDescriptionUrl(id))
	if err != nil {
		return nil, err
	}

	p.Name = getName(d)

	return &p, nil
}

func getDescriptionUrl(id int) string {
	return "https://www.urionlinejudge.com.br/repository/UOJ_" + strconv.Itoa(id) + "_en.html"
}

func getDocument(rawurl string) (*goquery.Document, error) {
	proxyURL, _ := url.Parse("socks5://localhost:1080")
	tr := &http.Transport{Proxy: http.ProxyURL(proxyURL)}
	client := &http.Client{Transport: tr, Timeout: 5 * time.Second}
	res, err := client.Get(rawurl)

	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	m := minify.New()
	m.AddFunc("text/html", html.Minify)
	mr := m.Reader("text/html", res.Body)

	return goquery.NewDocumentFromReader(mr)
}

func findContent(d *goquery.Document, selector string) *goquery.Selection {
	return d.Find(selector)
}

func findWholeTable(firstTableNode *goquery.Selection) *goquery.Selection {
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

func getURL(id int) string {
	return "https://www.urionlinejudge.com.br/judge/en/problems/view/" + strconv.Itoa(id)
}

func getName(d *goquery.Document) string {
	return findContent(d, nameSelector).Text()
}
