package main

import (
	"fmt"
	"strings"
)

// replacer replaces spaces with commas and tabs with commas.
// It's a package-level variable so we can easily reuse it, but
// this program doesn't take advantage of that fact.
var replacer = strings.NewReplacer(" ", ",", "\t", ",")

func main() {
	str := "a space- and\ttab-separated string"
	str = replacer.Replace(str)
	fmt.Println(str)
}
