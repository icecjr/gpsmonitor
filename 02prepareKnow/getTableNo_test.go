package prepareKnow

import "testing"

func TestCrc16(t *testing.T) {

	tests := []struct {
		name    string
		args    string
		wantCrc uint16
	}{
		{"first", "868120174576584", 0x6E},
		{"second", "868120177653141", 0x6E},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotCrc := Crc16(tt.args); gotCrc != tt.wantCrc {
				t.Errorf("Crc16() = %v, want %v", gotCrc, tt.wantCrc)
			}
			// t.Logf("%s successful!", tt.name)

		})
	}
}
