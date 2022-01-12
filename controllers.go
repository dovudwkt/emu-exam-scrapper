package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const examsFileName = "allExams.json"

var (
	// available exam periods in order as in the website
	periods = []string{"08:30", "12:30", "16:00"}

	// URL to be used for scrapping
	targetURL = "https://stdportal.emu.edu.tr/examlist.asp"
)

func importExamsHandler(w http.ResponseWriter, r *http.Request) {
	allExams := scrapExams()

	err := allExams.SaveJSON(examsFileName)
	if err != nil {
		err = errors.New("error saving exams: " + err.Error())
		reply(w, nil, http.StatusInternalServerError, err)
		return
	}

	log.Println("exams imported")
	reply(w, nil, http.StatusOK, nil)
}

func searchExamsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	coursesURI := r.URL.Query().Get("courses")
	coursesURI = strings.TrimSpace(strings.ToUpper(coursesURI))
	courses := strings.Split(coursesURI, ",")

	allExams, err := ParseJSON(examsFileName)
	if err != nil {
		reply(w, nil, http.StatusInternalServerError, err)
		return
	}

	exams := allExams.Find(courses)
	exams.printExams()

	respBody, err := exams.toJSON()
	if err != nil {
		reply(w, nil, http.StatusInternalServerError, err)
		return
	}

	reply(w, respBody, http.StatusOK, nil)
}

func scrapExams() Exams {
	data := Exams{}
	dates := []string{}

	doc, err := parseURL(targetURL)
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

	return data
}

// ========= HELPERS ==========

func ParseJSON(fileName string) (Exams, error) {
	result := make(Exams, 0)

	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, errors.New("error reading file: " + err.Error())
	}

	err = json.Unmarshal(file, &result)
	if err != nil {
		return nil, errors.New("error on json unmarshal: " + err.Error())
	}

	return result, nil
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

func reply(w http.ResponseWriter, body []byte, statusCode int, err error) {
	w.WriteHeader(statusCode)
	if err != nil {
		log.Println(err)
	}
	if len(body) > 0 {
		w.Write(body)
	}
	if err != nil {
		w.Write([]byte(err.Error()))
	}
}
