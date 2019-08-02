package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

var split = flag.Bool("split", false, "Split results in multiple files")

func printToFiles(glyphs []string) {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	for _, glyph := range glyphs {
		filename := dir + glyph + "_kerning_pairs.txt"

		fmt.Println("Creating ", filename)

		file, err := os.Create(filename)
		if err != nil {
			panic(err)
		}
		defer file.Close()

		file.WriteString(glyph + "\n")

		for _, pair := range glyphs {
			file.WriteString(pair + glyph + pair)
			file.WriteString(pair + glyph + pair + " ")
		}

		file.WriteString("\n")
	}
}

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
	flag.Parse()

	args := strings.Join(flag.Args()[0:1], "")

	if strings.HasSuffix(args, ".txt") {
		filepath := args
		fileBytes, err := ioutil.ReadFile(filepath)
		if err != nil {
			log.Fatal("Failed to read file: " + filepath)
		}

		str := strings.Trim(string(fileBytes), "\n")

		if *split {
			printToFiles(glyphs(str))
		}

		combine(glyphs(str))
	} else {
		combine(glyphs(args))
	}
}
