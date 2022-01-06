package main

import (
	"crypto/sha1"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/schollz/progressbar/v3"
)

var colorGreen = "\033[32m"
var colorRed = "\033[31m"
var colorYellow = "\033[33m"
var colorReset = "\033[0m"

type Results struct {
	Bad     []string
	Missing []string
	Good    []string
}

type Item struct {
	Filename string
}

func timestampToTime(timestamp string) time.Time {
	i, err := strconv.ParseInt(timestamp, 10, 64)
	if err != nil {
		panic(err)
	}

	return time.Unix(i, 0)
}

func fixTimestamp(f *os.File, timestamp time.Time) {
	if err := os.Chtimes(f.Name(), timestamp, timestamp); err != nil {
		log.Fatal(err)
	}
	//fmt.Println("Fixed timestamp on" + f.Name())
}

func printReport(results Results) {
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
}

func printSummary(results Results) {
	fmt.Printf("Missing: %d | Bad: %d | Good: %d\n", len(results.Missing), len(results.Bad), len(results.Good))
}

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

	bar := progressbar.Default(int64(len(xml.Files)))

	for _, file := range xml.Files {
		bar.Add(1)

		// Ignore the _files.xml
		if file.Name == matches[0] {
			continue
		}

		f, err := os.Open("./" + file.Name)

		// Check for presence of file
		if err != nil {
			results.Missing = append(results.Missing, file.Name)
			f.Close()
			continue
		}

		stat, err := f.Stat()
		if err != nil {
			log.Fatal(err)
		}

		// Check the file is the correct size
		if stat.Size() != file.Size {
			results.Bad = append(results.Bad, file.Name)
			f.Close()
			continue
		}

		// Modify timestamp if it doesn't match the xml file
		expectedTimestamp := timestampToTime(file.MTime)
		if stat.ModTime().Unix() != expectedTimestamp.Unix() {
			fixTimestamp(f, expectedTimestamp)
		}

		///////////
		// Really hacky hash verify
		// TODO: allow choice of hashing crc32 for faster etc.
		h := sha1.New()
		if _, err := io.Copy(h, f); err != nil {
			log.Fatal(err)
		}

		hash := fmt.Sprintf("%x", h.Sum(nil))

		// TODO: better matching. Case sensitivity will fail.
		if hash != file.SHA1 {
			results.Bad = append(results.Bad, file.Name)
			f.Close()
			continue
		}

		////////////

		results.Good = append(results.Good, file.Name)
		f.Close()
	}

	printReport(results)
	printSummary(results)
}
