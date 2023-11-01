package entity

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"sort"
	"time"
)

type Exams []Exam

type Exam struct {
	Course string
	Date   time.Time
}

type ResponseEl struct {
	Course string
	Date   string
	Time   string
}

type Response []ResponseEl

func (data Exams) ToResponse() Response {
	resp := make(Response, len(data))
	for i, exam := range data {

		date := exam.Date.Format("02 January 2006")
		period := exam.Date.Format("15:04")
		resp[i] = ResponseEl{
			Course: exam.Course,
			Date:   date,
			Time:   period,
		}
	}
	return resp
}

func (data Exams) Find(courseCodes []string) Exams {
	result := make(Exams, len(courseCodes))
	var resIDx int

	coursesMap := make(map[string]struct{})
	for _, v := range courseCodes {
		coursesMap[v] = struct{}{}
	}

	for _, exam := range data {
		if _, ok := coursesMap[exam.Course]; !ok {
			continue
		}

		result[resIDx] = exam
		resIDx++
	}

	return result[:resIDx]
}

func (data Exams) SaveJSON(fileName string) error {
	file, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		return errors.New("error on json marshal: " + err.Error())
	}
	err = ioutil.WriteFile(fileName, file, 0644)
	if err != nil {
		return errors.New("error on write file: " + err.Error())
	}

	return nil
}

func (data Exams) PrintExams() {
	for _, e := range data {
		fmt.Printf("%s - %s\n", e.Course, e.Date)
	}
}

func (data Exams) ToJSON() ([]byte, error) {
	return json.Marshal(data)
}

func (data Response) ToJSON() ([]byte, error) {
	return json.Marshal(data)
}

func (data Exams) SortByDate(asc bool) {
	if asc {
		sort.Sort(byDate(data))
	} else {
		sort.Sort(sort.Reverse(byDate(data)))
	}
}

type byDate Exams

func (f byDate) Len() int           { return len(f) }
func (f byDate) Swap(i, j int)      { f[i], f[j] = f[j], f[i] }
func (f byDate) Less(i, j int) bool { return f[i].Date.Before(f[j].Date) }
