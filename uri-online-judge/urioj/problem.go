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

// Package urioj parse html for problem page from RUI online judge
package urioj

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"time"

	"strings"

	"unicode"

	"github.com/PuerkitoBio/goquery"
	"github.com/dedis/student_18/dgcosi/code/onet/log"
	"github.com/lynn9388/supsub"
	"github.com/tdewolff/minify"
	html2 "github.com/tdewolff/minify/html"
	"golang.org/x/net/html"
)

// Problem selectors.
const (
	nameSelector        = "div.header > h1"     // selector for problem name
	descriptionSelector = "div.description"     // selector for problem main description block
	inputSelector       = "div.input"           // selector for problem input description block
	outputSelector      = "div.output"          // selector for problem output description block
	tableSelector       = "table"               // selector for table element
	sampleSelector      = "div.problem > table" // selector for test sample block
)

// Content is the interface of presented content in page
type Content interface {
	equal(interface{}) bool
	empty() bool
}

// TextContent is the plain text content
type TextContent string

//FileContent is the URL and plain text presentation of a file
type FileContent struct {
	URL  string
	Text string
}

// TableData is the content in a table cell
type TableData []Content

// TableRow is the content in a row of table
type TableRow []TableData

// TableContent is the content of a table
type TableContent struct {
	Head []string
	Data []TableRow
}

// Problem is the description of a problem
type Problem struct {
	ID          int       //Problem id
	URL         string    //page URL
	Name        string    //problem name
	Description []Content //main description
	Input       []Content //input description
	Output      []Content //output description
	Sample      []Content //test sample
}

func (t TextContent) equal(c interface{}) bool {
	return t == c.(TextContent)
}

func (t TextContent) empty() bool {
	return len(t) == 0
}

func (f FileContent) equal(c interface{}) bool {
	return reflect.DeepEqual(f, c.(FileContent))
}

func (f FileContent) empty() bool {
	return len(f.URL) == 0
}

func (t TableContent) equal(c interface{}) bool {
	return reflect.DeepEqual(t, c.(TableContent))
}

func (t TableContent) empty() bool {
	return len(t.Head) == 0
}

// NewProblem create a initialized problem based on the problem id. It returns
// a error if the problem page is inaccessible.
func NewProblem(id int) (*Problem, error) {
	d, err := getDocument(id)
	if err != nil {
		return nil, err
	}

	p := Problem{
		ID:          id,
		URL:         getURL(id),
		Name:        getName(d),
		Description: getDescription(d),
		Input:       getInput(d),
		Output:      getOutput(d),
		Sample:      getSample(d),
	}
	return &p, nil
}

// getDocument downloads the minimized problem page
func getDocument(id int) (*goquery.Document, error) {
	proxyURL, _ := url.Parse("socks5://localhost:1080")
	tr := &http.Transport{Proxy: http.ProxyURL(proxyURL)}
	client := &http.Client{Transport: tr, Timeout: 5 * time.Second}
	res, err := client.Get("https://www.urionlinejudge.com.br/repository/UOJ_" + strconv.Itoa(id) + "_en.html")

	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	m := minify.New()
	m.AddFunc("text/html", html2.Minify)
	mr := m.Reader("text/html", res.Body)

	return goquery.NewDocumentFromReader(mr)
}

// findWholeTable finds the <table> elements right next to each other
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

func getDescription(d *goquery.Document) []Content {
	return removeEmptyContent(getContent(d.Find(descriptionSelector)))
}

func getInput(d *goquery.Document) []Content {
	return removeEmptyContent(getContent(d.Find(inputSelector)))
}

func getOutput(d *goquery.Document) []Content {
	return removeEmptyContent(getContent(d.Find(outputSelector)))
}

func getSample(d *goquery.Document) []Content {
	return getContent(d.Find(sampleSelector))
}

func getContent(s *goquery.Selection) []Content {
	var content []Content

	var f func(*goquery.Selection)
	f = func(s *goquery.Selection) {
		for i := 0; i < s.Length(); i++ {
			n := s.Nodes[i]
			switch n.Data {
			case "p":
				content = append(content, renderParagraph(n)...)
			case "pre":
				content = append(content, TextContent(strings.TrimRightFunc(n.FirstChild.Data, unicode.IsSpace)))
			case "table":
				table, err := findWholeTable(s.Eq(i))
				if err != nil {
					log.Fatal(err)
				}
				tableContent, err := renderTable(table)
				if err != nil {
					log.Fatal(err)
				}
				content = append(content, *tableContent)
				i += len(tableContent.Data) - 1
			default:
				c := s.Eq(i).Children()
				for i := range c.Nodes {
					f(c.Eq(i))
				}
			}
		}
	}
	f(s)
	return content
}

func renderParagraph(n *html.Node) []Content {
	var content []Content
	var textBuf bytes.Buffer

	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.TextNode {
			textBuf.WriteString(processText(n.Data))
			return
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			switch c.Data {
			case "sup":
				textBuf.WriteString(supsub.ToSup(c.FirstChild.Data))
			case "sub":
				textBuf.WriteString(supsub.ToSub(c.FirstChild.Data))
			case "br":
				if textBuf.Len() != 0 {
					content = append(content, TextContent(textBuf.String()))
					textBuf.Reset()
				} else {
					content = append(content, TextContent(""))
				}
			case "img":
				image := renderFile(c)
				content = append(content, image)
				textBuf.WriteString(image.Text)
			default:
				f(c)
			}
		}
	}

	f(n)
	if textBuf.Len() != 0 {
		content = append(content, TextContent(textBuf.String()))
	}
	return content
}

func processText(s string) string {
	return strings.Map(func(r rune) rune {
		switch {
		case unicode.IsSpace(r):
			return ' '
		default:
			return r
		}
	}, s)
}

func removeEmptyContent(c []Content) []Content {
	for i := 0; i < len(c); i++ {
		if c[i].empty() {
			c = append(c[:i], c[i+1:]...)
		}
	}
	return c
}

func renderFile(n *html.Node) FileContent {
	var file FileContent
	for _, attr := range n.Attr {
		if attr.Key != "src" {
			continue
		}
		file.URL = attr.Val
		file.Text = fmt.Sprintf("<%v src=%q>", n.Data, attr.Val)
	}
	return file
}

func renderTable(s *goquery.Selection) (*TableContent, error) {
	var head []string
	var data []TableData

	th := s.Find("thead").Find("td")
	for i := range th.Nodes {
		head = append(head, th.Eq(i).Text())
	}

	tdp := s.Find("tbody").Find("td").Find("p")
	for _, n := range tdp.Nodes {
		var td TableData
		for _, c := range renderParagraph(n) {
			td = append(td, c)
		}
		data = append(data, td)
	}
	return newTable(head, data...)
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

func newTable(head []string, data ...TableData) (*TableContent, error) {
	numColumn := len(head)
	if len(data)%numColumn != 0 {
		return nil, errors.New("number of table data is not enough: " + strconv.Itoa(len(data)))
	}

	var table TableContent
	table.Head = head
	numRow := len(data) / numColumn
	for r := 0; r < numRow; r++ {
		table.Data = append(table.Data, data[r*numColumn:r*numColumn+numColumn])
	}
	return &table, nil
}
