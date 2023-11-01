package controllers

import (
	"testing"
)

func Test_scrapExams(t *testing.T) {
	t.Run("test", func(t *testing.T) {
		got := scrapExams()
		if len(got) < 1 {
			t.Errorf("scrapExams(): expecting more than 0 exams")
		}
	})
}
