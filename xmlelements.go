package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
)

func AnalyseXmlContent(content string) error {
	xmlfile, err := os.Open(content)
	if err != nil {
		return fmt.Errorf("error loading file: %v", err)
	}
	defer xmlfile.Close()
	xmlData, err := io.ReadAll(xmlfile)
	if err != nil {
		return fmt.Errorf("error reading XML data: %v", err)
	}

	type Element struct {
		XMLName xml.Name
		Attrs   []xml.Attr `xml:",any,attr"`
		Content string     `xml:",chardata"`
		Children []Element `xml:",any"`
	}
	var root Element
	if err := xml.Unmarshal(xmlData, &root); err != nil {
		return fmt.Errorf("error unmarshaling XML data: %v", err)
	}
	var analyzeFields func(Element, string)
	analyzeFields = func(elem Element, prefix string) {
		if prefix != "" {
			prefix += "+"
		}
		for _, child := range elem.Children {
			key := prefix + child.XMLName.Local
			fmt.Println("Field:", key)
			analyzeFields(child, key)
		}
	}

	// Start analyzing fields
	analyzeFields(root, "")

	return nil
}
func main() {
	xmlContent := "Data\\blog-12-12-2023.xml"
	err := AnalyseXmlContent(xmlContent)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
