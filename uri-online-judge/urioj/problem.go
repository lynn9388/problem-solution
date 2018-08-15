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
	"bytes"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/pkg/errors"
	"github.com/tdewolff/minify"
	html2 "github.com/tdewolff/minify/html"
	"golang.org/x/net/html"
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
	m.AddFunc("text/html", html2.Minify)
	mr := m.Reader("text/html", res.Body)

	return goquery.NewDocumentFromReader(mr)
}

func findWholeTable(firstRow *goquery.Selection) (*goquery.Selection, error) {
	if firstRow.Length() != 1 {
		return nil, errors.New("firstRow is not one row: " + strconv.Itoa(firstRow.Length()))
	}

	if !firstRow.Is(tableSelector) {
		return nil, errors.New("firstRow is not a table: " + getHTML(firstRow))
	}

	table := firstRow
	c := firstRow.Parent().Children()
	for i := c.IndexOfSelection(firstRow) + 1; i < c.Length(); i++ {
		n := c.Eq(i)
		if !n.Is(tableSelector) {
			break
		}
		table = table.AddSelection(n)
	}
	return table, nil
}

func getURL(id int) string {
	return "https://www.urionlinejudge.com.br/judge/en/problems/view/" + strconv.Itoa(id)
}

func getName(d *goquery.Document) string {
	return d.Find(nameSelector).Text()
}

func getHTML(s *goquery.Selection) string {
	var buf bytes.Buffer
	for _, n := range s.Nodes {
		if html.Render(&buf, n) != nil {
			return ""
		}
	}
	return buf.String()
}
