package mutils

import "testing"

func Test_convert(t *testing.T) {

	tests := []struct {
		name string
		args string
	}{
		{"first", "120.038916,29.280464"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			convertGD(120.038916, 29.280464)
		})
	}
}
