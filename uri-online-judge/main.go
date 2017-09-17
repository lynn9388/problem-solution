package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"unicode/utf8"

	"github.com/lynn9388/problem-solution/uri-online-judge/urioj"
	"github.com/olekukonko/tablewriter"
)

const lineWidth = 70

func main() {
	var id int
	fmt.Print("Please input the problem id:")
	for {
		fmt.Scanf("%d", &id)
		p, _ := urioj.NewProblem(id)
		createFile(p)
	}
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func comment(p *urioj.Problem) string {
	name := p.Name()
	description := p.Description()
	input := p.Input()
	output := p.Output()
	samples := p.Samples()

	command := new(bytes.Buffer)
	command.WriteByte('/')
	for i := 0; i < lineWidth-1; i++ {
		command.WriteByte('*')
	}
	command.WriteString("\n" +
		centerString(name) + "\n" +
		centerString(p.Url) + "\n\n" +
		formatStringSlice(description).String() + "\nInput\n*****\n" +
		formatStringSlice(input).String() + "\nOutput\n******\n" +
		formatStringSlice(output).String() + "\n" +
		formatSample(samples).String())
	for i := 0; i < lineWidth-1; i++ {
		command.WriteByte('*')
	}
	command.WriteString("/\n\n")

	return command.String()
}

func createFile(p *urioj.Problem) {
	inputFile := "template/main.go"
	code, err := ioutil.ReadFile(inputFile)
	check(err)

	outputDir := p.Id + " " + p.Name()
	if _, err := os.Stat(outputDir); os.IsNotExist(err) {
		err := os.Mkdir(outputDir, 0777)
		check(err)
	}

	outputFile := outputDir + "/main.go"
	if _, err := os.Stat(outputFile); os.IsNotExist(err) {
		err = ioutil.WriteFile(outputDir+"/main.go",
			append([]byte(comment(p)), code...), 0664)
		check(err)
		fmt.Println("File has been created:", outputFile)
	} else {
		fmt.Println("File already exists:", outputFile)
	}
}

func centerString(s string) string {
	spaceNum := (lineWidth - utf8.RuneCountInString(s)) / 2
	return strings.Repeat(" ", spaceNum) + s
}

func formatString(s string) *bytes.Buffer {
	words := strings.Split(s, " ")
	buf := new(bytes.Buffer)
	width := 0
	var word string
	var length int

	if strings.HasPrefix(s, urioj.Prefix) {
		buf.WriteString(urioj.Prefix)
	}

	for i := 0; i < len(words); i++ {
		word = words[i]
		length = utf8.RuneCountInString(word)
		if width+length < lineWidth || (width < lineWidth/2) ||
			(width < lineWidth-5 && width+length < lineWidth+5) {
			if width != 0 {
				buf.WriteByte(' ')
				width++
			}
			buf.WriteString(word)
			width += length
		} else {
			buf.WriteString("\n" + word)
			width = length
		}
	}
	buf.WriteString("\n")
	return buf
}

func formatStringSlice(s []string) *bytes.Buffer {
	buf := new(bytes.Buffer)
	for i, line := range s {
		buf.Write(formatString(line).Bytes())
		if i != len(s)-1 {
			buf.WriteString("\n")
		}
	}
	return buf
}

func formatSample(sample []urioj.Sample) *bytes.Buffer {
	buf := new(bytes.Buffer)

	table := tablewriter.NewWriter(buf)
	table.SetHeader([]string{"Sample Input", "Sample Output"})
	table.SetRowLine(true)
	table.SetRowSeparator("-")
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetAutoWrapText(false)

	for _, s := range sample {
		inputStr := ""
		for i, input := range s.Input {
			inputStr += input
			if i != len(s.Input)-1 {
				inputStr += "\n"
			}
		}
		outputStr := ""
		for i, output := range s.Output {
			outputStr += output
			if i != len(s.Output)-1 {
				outputStr += "\n"
			}
		}
		table.Append([]string{inputStr, outputStr})
	}

	table.Render()
	return buf
}
