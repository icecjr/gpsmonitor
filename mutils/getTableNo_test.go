package mutils

import (
	"fmt"
	"testing"
)

func TestCrc16(t *testing.T) {

	gotCrc := Crc16("868120170641689")
	fmt.Printf("Crc16() = %v", gotCrc)

}
