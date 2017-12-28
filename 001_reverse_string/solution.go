package main

import (
	"fmt"
	"regexp"
	"strings"
)

func reverse(items []string) []string {
	for i := 0; i < len(items)/2; i++ {
		j := len(items) - 1 - i
		items[i], items[j] = items[j], items[i]
	}

	return items
}

func main() {
	s := "Hello  world hello stars    hello   universe"
	a := regexp.MustCompile(`\s+`).Split(s, -1)

	fmt.Println(strings.Join(reverse(a), " "))
}
