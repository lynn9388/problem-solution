package urioj

import (
	"regexp"
	"strings"
	"unicode"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html"
)

var Prefix = map[string]string{
	"pre": "  ",
	"ul":  " • ",
}

type Sample struct {
	Input  []string
	Output []string
}

func (p *Problem) Name() string {
	return strings.TrimSpace(p.doc.Find("div.header > h1").Text())
}

func (p *Problem) Description() []string {
	return extractContent(p.doc.Find("div.description"))
}

func (p *Problem) Input() []string {
	return extractContent(p.doc.Find("div.input"))
}

func (p *Problem) Output() []string {
	return extractContent(p.doc.Find("div.output"))
}

func (p *Problem) Samples() []Sample {
	samples := make([]Sample, 0, 5)
	table := p.doc.Find("tbody")
	for i := range table.Nodes {
		sample := table.Eq(i).Find("td")
		input := formatSample(sample.First().Text())
		output := formatSample(sample.Last().Text())
		samples = append(samples, Sample{input, output})
	}
	return samples
}

func (p *Problem) Images() []string {
	var images []string

	var f func(*html.Node)
	f = func(n *html.Node) {
		switch n.Data {
		case "img":
			for _, attr := range n.Attr {
				if attr.Key == "src" {
					images = append(images, attr.Val)
					break
				}
			}
		default:
			if n.FirstChild != nil {
				for c := n.FirstChild; c != nil; c = c.NextSibling {
					f(c)
				}
			}
		}
	}

	for _, n := range p.doc.Find("div.problem").Nodes {
		f(n)
	}

	return images
}

func removeRedundantSpace(s string) string {
	reg := regexp.MustCompile(`[\s\p{Zs}]{2,}`)
	s = reg.ReplaceAllString(s, " ")
	replacer := strings.NewReplacer(" +", " ")
	return replacer.Replace(s)
}

func removeRedundantChar(s string) string {
	replacer := strings.NewReplacer("\r", "", "\t", "", "\n", " ", "\u200B", "", "\u00A0", " ")
	return replacer.Replace(s)
}

func text(n *html.Node) []string {
	var text []string

	var str string
	var f func(*html.Node)
	f = func(n *html.Node) {
		switch n.Data {
		case "br":
			if str != "\n" {
				if s := strings.TrimSpace(str); len(s) > 0 {
					text = append(text, removeRedundantSpace(s))
				}
				str = "\n"
			}
		case "sup":
			if len(str) > 0 && str[len(str)-1:] == " " {
				str = str[:len(str)-1]
			}
			str += "^" + strings.TrimSpace(n.FirstChild.Data)
		case "img":
			for _, attr := range n.Attr {
				if attr.Key == "src" {
					str += "<img src=\"" + attr.Val + "\">"
					break
				}
			}
		case "li":
			if s := strings.TrimSpace(str); len(s) > 0 {
				text = append(text, removeRedundantSpace(s))
			}
			str = ""
			fallthrough
		default:
			if n.Type == html.TextNode {
				data := removeRedundantChar(n.Data)
				if len(strings.TrimSpace(data)) > 0 {
					if n.Parent.Data == "strong" {
						data = strings.TrimLeftFunc(data, unicode.IsSpace)
						data = removeRedundantSpace(data)
					}
					str += data
				}
			} else if n.FirstChild != nil {
				for c := n.FirstChild; c != nil; c = c.NextSibling {
					f(c)
				}
			}
		}
	}

	f(n)
	if s := strings.TrimSpace(str); len(s) > 0 {
		text = append(text, removeRedundantSpace(s))
	}
	return text
}

func extractContent(s *goquery.Selection) []string {
	content := make([]string, 0, 10)

	var f func(*html.Node)
	f = func(n *html.Node) {
		switch n.Data {
		case "p":
			if len(content) != 0 {
				content = append(content, "")
			}
			content = append(content, text(n)...)
		case "pre":
			if len(content) != 0 {
				content = append(content, "")
			}
			for _, t := range text(n) {
				content = append(content, Prefix["pre"]+t)
			}
		case "ul":
			if len(content) != 0 {
				content = append(content, "")
			}
			for _, t := range text(n) {
				content = append(content, Prefix["ul"]+t)
			}
		default:
			if n.FirstChild != nil {
				for c := n.FirstChild; c != nil; c = c.NextSibling {
					f(c)
				}
			}
		}
	}
	for _, n := range s.Nodes {
		f(n)
	}

	return content
}

func formatSample(s string) []string {
	e := make([]string, 0, 5)
	lines := strings.Split(strings.TrimSpace(s), "\n")
	for _, line := range lines {
		e = append(e, removeRedundantSpace(strings.TrimSpace(line)))
	}
	return e
}
