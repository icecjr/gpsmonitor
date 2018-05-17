package prepareKnow

import "testing"

func Test_timetry(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"first"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			timetry()
		})
	}
}
