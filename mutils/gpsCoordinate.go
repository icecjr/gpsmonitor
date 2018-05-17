package mutils

import (
	"strings"
)

func TranferStrToLatitudeOrLogitude(message []string) float64 {
	var degree float64
	degree = float64(StrToI(strings.Join(message[:4], ""))) / float64(30000*60)
	//second := float64(strToI(strings.Join(message[:4], "")))/float64(30000*60) - float64(degree)
	//fmt.Printf("Latitude : %d %2.4f \n", degree, second)
	return degree
}

func TranferStringToGpsDir(message []string) int {
	//tempByte := []byte(message[0])
	//fmt.Printf("string %s and %s", message[0], message[1])
	a := (StrToI(message[0])%4)*256 + StrToI(message[1])
	//fmt.Printf("output a:%d", a)
	return a
}
