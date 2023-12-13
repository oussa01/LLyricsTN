package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
	"time"

	"golang.org/x/net/html"
)

type Item struct {

	Title   string   `xml:"title"`
	Content string   `xml:"content"`
	Term    string   `xml:"term"`
}

func ExtractTextFromHTML(htmlContent string) string {
	startTime := time.Now()
	doc, err := html.Parse(strings.NewReader(htmlContent))
	if err != nil {
		fmt.Println("Error parsing HTML:", err)
		return ""
	}

	var textContent []string
	var extractText func(*html.Node)
	extractText = func(n *html.Node) {
		if n.Type == html.TextNode {
			textContent = append(textContent, strings.TrimSpace(n.Data))
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			extractText(c)
		}
	}

	extractText(doc)

	cleanText := strings.Join(textContent, "\n")

	// Perform text cleaning
	re := regexp.MustCompile(`\n+`)
	cleanText = re.ReplaceAllString(cleanText, "\n")

	re = regexp.MustCompile(`\s{2,}`)
	cleanText = re.ReplaceAllString(cleanText, " ")

	re = regexp.MustCompile(`&nbsp;`)
	cleanText = re.ReplaceAllString(cleanText, " ")

	re = regexp.MustCompile(`<.*?>`)
	cleanText = re.ReplaceAllString(cleanText, "")
	elapsedTime := time.Since(startTime)
	fmt.Println("Execution time:", elapsedTime)
	return cleanText
}

func main() {
	xmlContent := `Data\\blog-12-12-2023.xml`
	xmlfile, err := os.Open(xmlContent)
	if err != nil {
		fmt.Println("error loading file :", err)
	}
	xmlData, _ := io.ReadAll(xmlfile)
	defer xmlfile.Close()

	var feed []Item
	err = xml.Unmarshal(xmlData, &feed)
	if err != nil {
		fmt.Println("Error unmarshaling XML:", err)
		return
	}
	var data []Item
	for _, item := range feed {
		if strings.HasPrefix(item.Term, "http") || strings.HasPrefix(item.Term, "https") {
			continue
		}
		item.Content = ExtractTextFromHTML(xmlContent)
		data = append(data, item)
	}
	for _, item := range data {
		fmt.Println("Title:", item.Title)
		fmt.Println("Content:", item.Content)
		fmt.Println("Term:", item.Term)
		fmt.Println("----------------------")
	}

	phrases := ExtractTextFromHTML(xmlContent)
	// Print the result
	fmt.Println("extracted :", phrases)
}
