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
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/dedis/student_18/dgcosi/code/onet/log"
	"github.com/olekukonko/tablewriter"
)

const lineWidth = 70

var description = "/" + strings.Repeat("*", lineWidth-1) + `
PROBLEM-NAME
PROBLEM-URL

PROBLEM-DESCRIPTION

Input
*****
PROBLEM-INPUT

Output
******
PROBLEM-OUTPUT

PROBLEM-SAMPLE
` + strings.Repeat("*", lineWidth-1) + "/"

var dir = "."

// NewDescription render a problem to plain text comment.
func NewDescription(id int) (string, error) {
	p, err := NewProblem(id)
	if err != nil {
		return "", err
	}

	description := strings.Replace(description, "PROBLEM-NAME", alignCenter(p.Name), 1)
	description = strings.Replace(description, "PROBLEM-URL", alignCenter(p.URL), 1)
	description = strings.Replace(description, "PROBLEM-DESCRIPTION", processContent(p.Description), 1)
	description = strings.Replace(description, "PROBLEM-INPUT", processContent(p.Input), 1)
	description = strings.Replace(description, "PROBLEM-OUTPUT", processContent(p.Output), 1)
	description = strings.Replace(description, "PROBLEM-SAMPLE", processContent(p.Sample), 1)

	return description, nil
}

// NewDescriptionFile creates a description file of a problem in a folder with
// relative resources (like images). Default directory to save file will be
// updated at the same time
func NewDescriptionFile(id int, path string) error {
	dir = filepath.Dir(path)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err := os.MkdirAll(dir, 0777); err != nil {
			return err
		}
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		description, err := NewDescription(id)
		if err != nil {
			return err
		}

		if err := ioutil.WriteFile(path, []byte(description), 0664); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("file already exists: %v", path)
	}

	return nil
}

// processContent processes content with correspond behavior (likes download image)
// and returns a string represents these contents.
func processContent(cs []Content) string {
	var buf bytes.Buffer

	var f func(c Content)
	f = func(c Content) {
		switch c.(type) {
		case TextContent:
			buf.WriteString(limitWidth(c) + "\n")
		case ListText:
			buf.WriteString(strings.Replace(limitWidth(c), "   ", " • ", 1))
		case ListContent:
			for _, item := range c.(ListContent) {
				for _, i := range item {
					f(i)
				}
			}
			buf.WriteString("\n")
		case FileContent:
			url := c.(FileContent).URL
			downloadFile(dir+"/"+path.Base(url), url)
		case TableContent:
			buf.WriteString(tableToText(c.(TableContent)) + "\n")
		}
	}

	for _, c := range cs {
		f(c)
	}
	return strings.TrimRightFunc(buf.String(), unicode.IsSpace)
}

// limitWidth formats content to a limit width.
func limitWidth(c Content) string {
	var lines []string
	var prefix string

	switch c.(type) {
	case TextContent:
		lines = strings.Split(string(c.(TextContent)), "\n")
		prefix = ""
	case ListText:
		lines = strings.Split(string(c.(ListText)), "\n")
		prefix = "   "
	}

	prefixWidth := utf8.RuneCountInString(prefix)
	var buf bytes.Buffer
	for _, line := range lines {
		words := strings.Split(line, " ")
		width := 0
		for i, word := range words {
			wl := utf8.RuneCountInString(word)
			if width+wl < lineWidth || width < lineWidth/2 ||
				width < lineWidth-5 && width+wl < lineWidth+5 {
				if width == 0 {
					buf.WriteString(prefix)
					width += prefixWidth
				}
				if wl == 0 || i > 0 && utf8.RuneCountInString(words[i-1]) != 0 {
					buf.WriteByte(' ')
					width++
				}
				buf.WriteString(word)
				width += wl
			} else {
				buf.WriteString("\n" + prefix + word)
				width = prefixWidth + wl
			}
		}
		buf.WriteByte('\n')
	}
	return buf.String()
}

// alignCenter formats a string on the center of a fixed width.
func alignCenter(s string) string {
	spaceNum := (lineWidth - utf8.RuneCountInString(s)) / 2
	return strings.Repeat(" ", spaceNum) + s
}

// tableToText converts a table to plain text.
func tableToText(t TableContent) string {
	var buf bytes.Buffer

	table := tablewriter.NewWriter(&buf)
	table.SetHeader(t.Head)
	table.SetRowLine(true)
	table.SetRowSeparator("-")
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetAutoWrapText(false)

	for _, td := range t.Data {
		var row []string
		for _, d := range td {
			row = append(row, processContent(d))
		}
		table.Append(row)
	}

	table.Render()
	return buf.String()
}

// downloadFile downloads a file
func downloadFile(path string, url string) {
	out, err := os.Create(path)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(out, resp.Body)
}
