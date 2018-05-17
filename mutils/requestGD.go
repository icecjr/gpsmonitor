package mutils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type returnGD struct {
	Status    string `json:"status"`
	Info      string `json:"info"`
	Infocode  string `json:"infocode"`
	Locations string `json:"locations"`
}

func ConvertGD(gps_wd, gps_jd float64) (float64, float64) {
	//这里添加post的body内容
	str := strconv.FormatFloat(gps_wd, 'f', -1, 64) + "," + strconv.FormatFloat(gps_jd, 'f', -1, 64)
	//fmt.Println(gps_wd, gps_jd, str)
	data := make(url.Values)

	data["locations"] = []string{str}
	data["coordsys"] = []string{"gps"}
	data["output"] = []string{"json"}
	data["key"] = []string{"d3b89b31a7ca0bf078ada78fc7faf68d"}
	res, err := http.PostForm("http://restapi.amap.com/v3/assistant/coordinate/convert", data)
	if err != nil {
		fmt.Println(err.Error())
		return 0, 0
	}
	defer res.Body.Close()

	//fmt.Println("post send success")
	body, err := ioutil.ReadAll(res.Body)
	CheckError(err)
	tempStr := &returnGD{}
	err = json.Unmarshal(body, &tempStr)
	CheckError(err)
	//fmt.Println(tempStr)
	//fmt.Printf("return string %s \n", body)
	locations := strings.Split(tempStr.Locations, ",")
	wd, err := strconv.ParseFloat(locations[0], 64)
	jd, err := strconv.ParseFloat(locations[1], 64)
	return wd, jd
}
func Get_gps_from_lbs() (float64, float64) {
	data := make(url.Values)
	data["accesstype"] = []string{"0"}
	data["cmd"] = []string{"0"}
	data["output"] = []string{"json"}
	data["network"] = []string{"GPRS"}
	data["key"] = []string{"1638c7c31128b56a5e5ecc72836ca9a0"}
	res, err := http.PostForm("http://apilocate.amap.com/position", data)
	if err != nil {
		fmt.Println(err.Error())
		return 0, 0
	}
	defer res.Body.Close()

	//fmt.Println("post send success")
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		// handle error
	}
	var tempStr []string
	tempStr = strings.Split(string(body), " ")
	//fmt.Println(tempStr[0], tempStr[1])
	wd, err := strconv.ParseFloat(tempStr[0], 64)
	jd, err := strconv.ParseFloat(tempStr[1], 64)
	return wd, jd
}
