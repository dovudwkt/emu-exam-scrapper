package entity

import (
	"reflect"
	"testing"
	"time"
)

func TestExams_SortByDate(t *testing.T) {
	now := time.Now()
	nextMonth := now.AddDate(0, 1, 0)
	inTwoMonths := now.AddDate(0, 2, 0)
	prevMonth := now.AddDate(0, -1, 0)

	input := []Exam{
		{
			Course: "now",
			Date:   now,
		},
		{
			Course: "nextMonth",
			Date:   nextMonth,
		},
		{
			Course: "prevMonth",
			Date:   prevMonth,
		},
		{
			Course: "inTwoMonths",
			Date:   inTwoMonths,
		},
	}

	type args struct {
		asc bool
	}
	tests := []struct {
		name string
		data Exams
		args args
		exp  Exams
	}{
		{
			name: "Sort in descending order",
			data: input,
			args: args{
				asc: false,
			},
			exp: []Exam{
				{
					Course: "inTwoMonths",
					Date:   inTwoMonths,
				},
				{
					Course: "nextMonth",
					Date:   nextMonth,
				},
				{
					Course: "now",
					Date:   now,
				},
				{
					Course: "prevMonth",
					Date:   prevMonth,
				},
			},
		},
		{
			name: "Sort in ascending order",
			data: input,
			args: args{
				asc: true,
			},
			exp: []Exam{
				{
					Course: "prevMonth",
					Date:   prevMonth,
				},
				{
					Course: "now",
					Date:   now,
				},
				{
					Course: "nextMonth",
					Date:   nextMonth,
				},
				{
					Course: "inTwoMonths",
					Date:   inTwoMonths,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.data.SortByDate(tt.args.asc)
			if !reflect.DeepEqual(tt.data, tt.exp) {
				t.Errorf("expected %v but got %v", tt.exp, tt.data)
			}
		})
	}
}
