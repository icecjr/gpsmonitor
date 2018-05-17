package mutils

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/golang/glog"
)

func CheckError(err error) {
	if err != nil {
		glog.Errorf("Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

// //TODO 判断数据和其校验码是否一致
// func CheckCRCItu(message []string, crc []string) bool {
// 	return true
// }
func StrToI(str string) int {
	//fmt.Printf("convert string is %s\n", str)
	a, err := strconv.ParseInt(str, 16, 0)
	CheckError(err)
	return int(a)
}

func TranferByteToTime(message []string) time.Time {
	the_time := time.Date(2000+StrToI(message[0]), time.Month(StrToI(message[1])), StrToI(message[2]), StrToI(message[3]), StrToI(message[4]), StrToI(message[5]), 0, time.Local)
	return the_time
}

//byte[] -> string
func TranferBytearrayToString(payload []byte) string {
	// n := bytes.IndexByte(payload, 0)
	// for eachOne := range payload {
	// 	eachOne = eachOne & 0xFF
	// }
	// return string(payload[:n])
	var sa = make([]string, 0)
	for _, v := range payload {
		sa = append(sa, fmt.Sprintf("%02X", v))
	}
	ss := strings.Join(sa, ":")
	return ss
}
