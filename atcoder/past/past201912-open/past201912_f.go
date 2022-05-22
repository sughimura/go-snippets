// 20:56 - 21:32(例1,2->OK) - 21:34(AC)

package main

import (
	"fmt"
	"sort"
	"strings"
	"unicode"
)

func main() {
	var s string
	fmt.Scan(&s)
	var words []string
	var status string
	status = "end"
	var word string
	for _, si := range s {
		word = word + string(si)
		if unicode.IsUpper(si) {
			if status == "end" {
				status = "wording"
			} else if status == "wording" {
				status = "end"
				words = append(words, word)
				word = ""
			}
		}
	}

	n := len(words)
	wordsLower := make([]string, n)
	wordsMap := make(map[string]string)
	for i, word := range words {
		wordLower := strings.ToLower(word)
		wordsLower[i] = wordLower
		wordsMap[wordLower] = word
	}

	var result string
	sort.Strings(wordsLower)
	for _, word := range wordsLower {
		result = result + wordsMap[word]
	}

	fmt.Println(result)
}
