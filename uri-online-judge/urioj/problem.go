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

	"github.com/PuerkitoBio/goquery"
)

type Problem struct {
	Id  int
	Url string
	Doc *goquery.Document
}

func NewProblem(id int) (*Problem, error) {
	p := Problem{Id: id, Url: getURL(id)}

	proxyUrl, _ := url.Parse("socks5://localhost:1080")
	tr := &http.Transport{Proxy: http.ProxyURL(proxyUrl)}
	client := &http.Client{Transport: tr, Timeout: 5 * time.Second}
	res, err := client.Get(getDescriptionUrl(p.Id))

	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	p.Doc, err = goquery.NewDocumentFromReader(res.Body)
	return &p, err
}

func getURL(id int) string {
	return "https://www.urionlinejudge.com.br/judge/en/problems/view/" + strconv.Itoa(id)
}

func getDescriptionUrl(id int) string {
	return "https://www.urionlinejudge.com.br/repository/UOJ_" + strconv.Itoa(id) + "_en.html"
}
