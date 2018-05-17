package mutils

import (
	"fmt"
	"testing"
)

func Test_doprint(t *testing.T) {

}
func Test_tranferByteToTime(t *testing.T) {
	fmt.Println(TranferByteToTime([]string{"12", "04", "10", "0f", "3b", "3b"}))
}
func Test_tranferStrToLatitudeOrLogitude(t *testing.T) {
	TranferStrToLatitudeOrLogitude([]string{"03", "30", "3b", "b8"})
	TranferStrToLatitudeOrLogitude([]string{"0c", "e7", "13", "90"})
}
