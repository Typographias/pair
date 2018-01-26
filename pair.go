package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	args := strings.Join(os.Args[1:], "")
	glyphs := strings.Split(args, "")

	for _, glyph := range glyphs {
		fmt.Println(glyph)

		for _, pair := range glyphs {
			fmt.Print(pair + glyph + pair)
			fmt.Print(pair + glyph + pair + " ")
		}

		fmt.Print("\n")
	}
}
