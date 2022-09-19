package models

import "testing"

func TestSendSms(t *testing.T) {
	type args struct {
		phone string
		code  int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"", args{"1**********", 123456}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SendSms(tt.args.phone, tt.args.code); got != tt.want {
				t.Errorf("SendSms() = %v, want %v", got, tt.want)
			}
		})
	}
}
