package main

import (
	"testing"
)

func Test_doPackage(t *testing.T) {
	type args struct {
		deviceId string
		text     string
	}
	tests := []struct {
		name string
		args args
	}{
		{"testGPSMessage", args{deviceId: "868120170641689", text: "78:78:22:22:12:04:10:0f:3b:3b:cc:03:30:3b:b8:0c:e7:13:90:27:d4:52:01:cc:00:68:54:00:60:59:01:00:00:02:93:f3:b4:0d:0a"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			doPackage(tt.args.deviceId, tt.args.text)
		})
	}
}
