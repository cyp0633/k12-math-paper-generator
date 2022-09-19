package models

import (
	"reflect"
	"testing"
)

func TestAnswers(t *testing.T) {
	tests := []struct {
		name     string
		username string
		want     []float64
	}{
		{"", "123456", []float64{288, 12}},
	}
	p := &Problem{
		User:       "123456",
		Expression: "24*12",
		Answer:     288,
	}
	p2 := &Problem{
		User:       "123456",
		Expression: "24+12-24",
		Answer:     12,
	}
	DB.Create(p)
	DB.Create(p2)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Answers(tt.username); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Answers() = %v, want %v", got, tt.want)
			}
		})
	}
	DB.Delete(p)
	DB.Delete(p2)
}

func TestCheckDupProblem(t *testing.T) {
	type args struct {
		problem string
		user    string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"", args{"24+12-24", "123456"}, false},
		{"", args{"24*12", "123456"}, true},
	}
	p := &Problem{
		User:       "123456",
		Expression: "24*12",
	}
	DB.Create(p)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckDupProblem(tt.args.problem, tt.args.user); got != tt.want {
				t.Errorf("CheckDupProblem() = %v, want %v", got, tt.want)
			}
		})
	}
	DB.Delete(p)
}

func TestClearProblems(t *testing.T) {
	tests := []struct {
		name     string
		username string
	}{
		{"", "123456"},
	}
	p := &Problem{
		User:       "123456",
		Expression: "24*12",
		Answer:     288,
	}
	p2 := &Problem{
		User:       "123456",
		Expression: "24+12-24",
		Answer:     12,
	}
	DB.Create(p)
	DB.Create(p2)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ClearProblems(tt.username)
			if len(Answers("123456")) != 0 {
				t.Errorf("ClearProblem() Fail")
			}
		})
	}
}

func TestWriteProblemToDb(t *testing.T) {
	type args struct {
		problem string
		ans     float64
		usr     string
	}
	tests := []struct {
		name string
		args args
	}{
		{"", args{"12-3", 9, "123456"}},
		{"", args{"12*5.1", 25.5, "123456"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			WriteProblemToDb(tt.args.problem, tt.args.ans, tt.args.usr)
		})
	}
	if !reflect.DeepEqual(Answers("123456"), []float64{9, 25.5}) {
		t.Errorf("WriteProblemToDb() Fail")
	}
	ClearProblems("123456")
}
