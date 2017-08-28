// https://tour.golang.org/moretypes/23
package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	a := make(map[string]int)
	w := strings.Fields(s)
	for i := 0; i < len(w); i++ {
		a[w[i]]++
	}
	return a
}

func main() {
	wc.Test(WordCount)
}
