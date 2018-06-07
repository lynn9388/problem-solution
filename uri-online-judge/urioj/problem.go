/*
 * Copyright 2018 Lynn
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
