package controllers

import (
	"reflect"
	"testing"

	"github.com/dovudwkt/emu-exam-scrapper/entity"
)

func Test_scrapExams(t *testing.T) {
	tests := []struct {
		name string
		want entity.Exams
	}{
		{
			name: "test1",
			want: []entity.Exam{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := scrapExams(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("scrapExams() = %v, want %v", got, tt.want)
			}
		})
	}
}
