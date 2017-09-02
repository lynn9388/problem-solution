package urioj

import (
	"strings"
	"regexp"
)

func (p *Problem) GetName() string {
	return p.doc.Find("div.header > h1").Text()
}

func (p *Problem) GetDescription() string {
	return removeRedundantSpace(strings.TrimSpace(p.doc.Find("div.description").Text()))
}

func (p *Problem) GetInput() string {
	return removeRedundantSpace(strings.TrimSpace(p.doc.Find("div.input").Text()))
}

func (p *Problem) GetOutput() string {
	return removeRedundantSpace(strings.TrimSpace(p.doc.Find("div.output").Text()))
}

func (p *Problem) GetSamples() map[string]string {
	samples := make(map[string]string)
	table := p.doc.Find("tbody")
	for i := range table.Nodes {
		sample := table.Eq(i).Find("td")
		input := format(sample.First().Text())
		output := format(sample.Last().Text())
		samples[input] = output
	}
	return samples
}

func removeRedundantSpace(s string) string {
	reg := regexp.MustCompile(`[\s\p{Zs}]{2,}`)
	return reg.ReplaceAllString(s, " ")
}

func format(s string) string {
	lines := strings.Split(strings.TrimSpace(s), "\n")
	for i, line := range lines {
		lines[i] = removeRedundantSpace(strings.TrimSpace(line))
	}
	return strings.Join(lines, "\n")
}
