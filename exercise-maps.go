// https://tour.golang.org/moretypes/23
package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	count := make(map[string]int)
	words := strings.Fields(s)
	for _, word := range words {
		count[word]++
	}	
	return count
}

func main() {
	wc.Test(WordCount)
}
