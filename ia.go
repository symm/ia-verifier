package main

import (
	"encoding/xml"
	"io/ioutil"
	"log"
	"os"
)

type Files struct {
	XMLName xml.Name `xml:"files"`
	Files   []File   `xml:"file"`
}

type File struct {
	XMLName xml.Name `xml:"file"`
	Name    string   `xml:"name,attr"`
	Source  string   `xml:"source,attr"`
	MTime   string   `xml:"mtime"`
	Size    int64    `xml:"size"`
	MD5     string   `xml:"md5"`
	CRC32   string   `xml:"crc32"`
	SHA1    string   `xml:"sha1"`
	Format  string   `xml:"format"`
}

func ReadXMLFile(filename string) Files {
	xmlFile, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	defer xmlFile.Close()

	byteValue, err := ioutil.ReadAll(xmlFile)

	if err != nil {
		log.Fatal(err)
	}

	var files Files

	err = xml.Unmarshal(byteValue, &files)

	if err != nil {
		log.Fatal(err)
	}

	return files
}
