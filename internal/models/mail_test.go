package models

import "testing"

func TestSendMail(t *testing.T) {
	type args struct {
		dest string
		code int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"QQ", args{"*******@qq.com", 123456}, true},
		{"163", args{"*******@163.com", 123456}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SendMail(tt.args.dest, tt.args.code); got != tt.want {
				t.Errorf("SendMail() = %v, want %v", got, tt.want)
			}
		})
	}
}
