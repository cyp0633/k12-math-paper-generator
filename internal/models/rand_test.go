package models

import "testing"

func TestGenNum(t *testing.T) {
	type args struct {
		min int
		max int
	}
	tests := []struct {
		name string
		args args
	}{
		{"++", args{3, 120}},
		{"0+", args{0, 100}},
		{"-+", args{-3, 20}},
		{"-0", args{-18, 0}},
		{"--", args{-70, -20}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GenNum(tt.args.min, tt.args.max)
			if got < tt.args.min || got >= tt.args.max {
				t.Errorf("GenNum() = %v, want [%v, %v)", got, tt.args.min, tt.args.max)
			}
		})
	}
}

func Test_genOp(t *testing.T) {
	tests := []struct {
		name  string
		level int
		want  int
	}{
		{"PrimarySchool", UserTypePrimarySchool, 0},
		{"JuniorHigh", UserTypeJuniorHigh, 0},
		{"SeniorHigh", UserTypeSeniorHigh, 0},
		{"default", 3, -1},
		{"default", -5, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := genOp(tt.level)
			switch tt.level {
			case UserTypePrimarySchool:
				if got < 0 || got >= 4 {
					t.Errorf("genOp() = %v, want [0,4)", got)
				}
			case UserTypeJuniorHigh:
				if got < 0 || got >= 6 {
					t.Errorf("genOp() = %v, want [0,6)", got)
				}
			case UserTypeSeniorHigh:
				if got < 0 || got >= 9 {
					t.Errorf("genOp() = %v, want [0,9)", got)
				}
			default:
				if tt.want != -1 {
					t.Errorf("genOp() = %v, want -1", got)
				}
			}
		})
	}
}
