package main

import (
	"fmt"
	"sort"
	"strings"
)

func groupAnagrams(str []string) [][]string {
	wordMap := make(map[string][]string)
	result := [][]string{}

	for _, word := range str {
		sortedWord := sortString(word)
		wordMap[sortedWord] = append(wordMap[sortedWord], word)
	}

	for _, group := range wordMap {
		result = append(result, group)
	}

	return result
}

func sortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

func main() {
	str := []string{"cook", "save", "taste", "aves", "vase", "state", "map"}
	result := groupAnagrams(str)

	fmt.Println("[")
	for i, group := range result {
		fmt.Print("  [")
		for j, word := range group {
			fmt.Printf("'%s'", word)
			if j < len(group)-1 {
				fmt.Print(",")
			}
		}
		if i < len(result)-1 {
			fmt.Println("],")
		} else {
			fmt.Println("]")
		}
	}
	fmt.Println("]")
}
