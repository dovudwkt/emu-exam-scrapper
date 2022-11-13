package entity

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
)

type Exams []Exam

type Exam struct {
	Course string
	Date   string
	Period string
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

	return result
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
		fmt.Printf("%s - %s - %s\n", e.Course, e.Date, e.Period)
	}
}

func (data Exams) ToJSON() ([]byte, error) {
	return json.Marshal(data)
}
