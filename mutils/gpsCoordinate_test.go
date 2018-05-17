package mutils

import "testing"

func Test_tranferStringToGpsDir(t *testing.T) {

	tests := []struct {
		name string
		args []string
		//	want int
	}{
		{"first", []string{"d4", "52"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			TranferStringToGpsDir(tt.args)
			//if got := tranferStringToGpsDir(tt.args); got != tt.want {
			//t.Errorf("tranferStringToGpsDir() = %v, want %v", got, tt.want)

		})
	}
}
