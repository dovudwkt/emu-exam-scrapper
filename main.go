package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Exams []exam

type exam struct {
	Period string
	Course string
	Date   string
}

var (
	periods = []string{"08:30", "12:30", "16:00"}
	url     = "https://stdportal.emu.edu.tr/examlist.asp"
)

func main() {
	scrapExams()
}

func scrapExams() {
	data := Exams{}
	dates := []string{}

	doc, err := parseURL(url)
	if err != nil {
		fmt.Println(err)
	}

	doc.Find("table").Each(func(tableIdx int, tableNode *goquery.Selection) {
		if tableIdx == 0 {
			tableNode.Find("tbody tr td font").Each(func(j int, dateStr *goquery.Selection) {
				dates = append(dates, strings.TrimSpace(dateStr.Text()))
			})
		} else {
			tableNode.Find("tbody tr").EachWithBreak(func(trIdx int, tr *goquery.Selection) bool {
				if trIdx > 2 {
					return false
				}

				tr.Find("td font").EachWithBreak(func(tdIdx int, courseCode *goquery.Selection) bool {
					cCode := strings.TrimSpace(courseCode.Text())
					if cCode == "" {
						return true
					}
					entry := exam{
						Period: periods[tableIdx-1],
						Course: cCode,
						Date:   dates[tdIdx],
					}

					data = append(data, entry)
					return true
				})

				return true
			})
		}
	})

	data.printExams()
}

func (data Exams) printExams() {
	for _, e := range data {
		fmt.Printf("%s - %s - %s\n", e.Course, e.Date, e.Period)
	}
}

func parseURL(url string) (*goquery.Document, error) {
	// Request the HTML page.
	res, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("Error: %w", err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("status code error: %d %s %w", res.StatusCode, res.Status, err)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, err
	}

	return doc, nil
}
