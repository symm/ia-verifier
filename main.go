package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type Results struct {
	Bad     []string
	Missing []string
	Good    []string
}

type Item struct {
	Filename string
}

var colorGreen = "\033[32m"
var colorRed = "\033[31m"
var colorYellow = "\033[33m"
var colorReset = "\033[0m"

func main() {

	matches, err := filepath.Glob("*_files.xml")

	if err != nil {
		log.Fatal(err)
	}

	if len(matches) > 1 || len(matches) == 0 {
		log.Fatal("Could not determine which XML file to use. Found: " + strings.Join(matches, ","))
	}

	xml := ReadXMLFile(matches[0])

	results := Results{}

	for _, file := range xml.Files {
		if file.Name == matches[0] {
			continue
		}

		f, err := os.Open("./" + file.Name)
		defer f.Close()

		if err != nil {
			results.Missing = append(results.Missing, file.Name)
			continue
		}

		stat, err := f.Stat()
		if err != nil {
			log.Fatal(err)
		}
		if stat.Size() != file.Size {
			results.Bad = append(results.Bad, file.Name)
		}

		results.Good = append(results.Good, file.Name)

	}

	fmt.Println("Bad:")
	for _, filename := range results.Bad {
		fmt.Println(string(colorRed), filename, string(colorReset))
	}

	fmt.Println("Good:")
	for _, filename := range results.Good {
		fmt.Println(string(colorGreen), filename, string(colorReset))
	}

	fmt.Println("Missing:")
	for _, filename := range results.Missing {
		fmt.Println(string(colorYellow), filename, string(colorReset))
	}

	fmt.Printf("Missing: %d | Bad: %d | Good: %d\n", len(results.Missing), len(results.Bad), len(results.Good))
}
