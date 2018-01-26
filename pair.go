package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func combine(glyphs []string) {
	for _, glyph := range glyphs {
		fmt.Println(glyph)

		for _, pair := range glyphs {
			fmt.Print(pair + glyph + pair)
			fmt.Print(pair + glyph + pair + " ")
		}

		fmt.Print("\n")
	}
}

func glyphs(str string) []string {
	return strings.Split(str, "")
}

func main() {
	args := strings.Join(os.Args[1:], "")

	if strings.HasSuffix(args, ".txt") {
		filepath := args
		fileBytes, err := ioutil.ReadFile(filepath)
		if err != nil {
			log.Fatal("Failed to read file: " + filepath)
		}

		str := strings.Trim(string(fileBytes), "\n")
		combine(glyphs(str))
	} else {
		combine(glyphs(args))
	}
}
