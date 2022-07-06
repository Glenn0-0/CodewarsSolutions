package main

import (
	"fmt"
	"strings"
	"regexp"
)

func arrayToString(words []string) string {
	result := ""
	for index, word := range words {
		if index != 0 {
		  result += strings.Title(word)
		} else {
		  result += word
		}
	  }
	return result
}

func ToCamelCase(s string) string {
	words := regexp.MustCompile("-|_").Split(s, -1)
	camelString := arrayToString(words)
	return camelString
}

func main() {
	fmt.Println(ToCamelCase("the-very_beginning"))
}

