package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
)

type Entry struct {
	ID         string     `xml:"id"`
	Published  string     `xml:"published"`
	Updated    string     `xml:"updated"`
	Categories []Category `xml:"category"`
	Title      string     `xml:"title"`
	Content    Content    `xml:"content"`
}

type Category struct {
	Scheme string `xml:"scheme,attr"`
	Term   string `xml:"term,attr"`
}

type Content struct {
	Type  string `xml:"type,attr"`
	Value string `xml:",chardata"`
}

type Feed struct {
	Entries []Entry `xml:"entry"`
}

func main() {
	xmlContent, err := ioutil.ReadFile("Data\\blog-12-12-2023.xml")
	if err != nil {
		fmt.Println("Error reading XML file:", err)
		return
	}

	var feed Feed
	err = xml.Unmarshal(xmlContent, &feed)
	if err != nil {
		fmt.Println("Error unmarshaling XML:", err)
		return
	}

	// Create a new Feed to store extracted information
	var extractedFeed Feed

	// Iterate through entries and extract information
	for _, entry := range feed.Entries {
		// You can customize this part to filter out entries or modify content as needed
		extractedFeed.Entries = append(extractedFeed.Entries, entry)
	}

	// Marshal the extracted information back to XML
	extractedXML, err := xml.MarshalIndent(extractedFeed, "", "  ")
	if err != nil {
		fmt.Println("Error marshaling XML:", err)
		return
	}

	// Save the extracted information to a new XML file
	err = ioutil.WriteFile("extracted_blog_info.xml", extractedXML, 0644)
	if err != nil {
		fmt.Println("Error writing XML file:", err)
		return
	}

	fmt.Println("Extracted information saved to extracted_blog_info.xml")
}
