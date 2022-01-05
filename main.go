package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type Results struct {
	//	Bad     []string
	Missing []string
	Present []string
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

	if len(matches) > 1 {
		log.Fatal("Could not determine which XML file to use. Found more than one: " + strings.Join(matches, ","))
	}

	xml := ReadXMLFile(matches[0])

	results := Results{}

	for _, file := range xml.Files {
		f, err := os.Open("./" + file.Name)
		if file.Name == matches[0] {
			continue
		}

		if err != nil {
			results.Missing = append(results.Missing, file.Name)
			continue
		}

		results.Present = append(results.Present, file.Name)
		defer f.Close()
	}

	fmt.Println("Missing:")
	for _, filename := range results.Missing {
		fmt.Println(string(colorYellow), filename, string(colorReset))
	}

	// fmt.Println("Bad:")
	// for _, filename := range results.Bad {
	// 	fmt.Println(string(colorRed), filename, string(colorReset))
	// }

	fmt.Println("Present:")
	for _, filename := range results.Present {
		fmt.Println(string(colorGreen), filename, string(colorReset))
	}

	//fmt.Printf("Missing: %d | Bad: %d | Present: %d\n", len(results.Missing), len(results.Bad), len(results.Present))
	fmt.Printf("Missing: %d | Present: %d\n", len(results.Missing), len(results.Present))

}
