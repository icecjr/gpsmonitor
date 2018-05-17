package main

import (
	"errors"
	"fmt"
	. "project/elangshen/mutils"

	"strings"

	"github.com/golang/glog"
)

func checkMsg(str []string, load []byte) bool {
	_, err3 := checkStopBits(str)
	if err3 != nil {
		glog.Info(err3)
		glog.Flush()
		return false
	}

	_, err1 := checksumMsg(load)
	if err1 != nil {
		glog.Info(err1)
		glog.Flush()
		return false
	}

	_, err2 := checkPackageLength(str)
	if err2 != nil {
		glog.Info(err2)
		glog.Flush()
		return false
	}

	return true

}

//TODO checksum message 出错 不知道原因
//[78 78 0A 13 04 06 04 00 01 00 A8 AD 7D 0D 0A]
//crc is 31705（算出来的） checksum is 95269（AD 7D）
func checksumMsg(str []byte) (bool, error) {
	inLen := len(str)
	var crc uint16
	//fmt.Println(fmt.Sprintf("%X", str[0]), fmt.Sprintf("%X", str[1]))
	if fmt.Sprintf("%X", str[0]) == "78" && fmt.Sprintf("%X", str[1]) == "78" {
		crc = Crc16byte(str[2 : inLen-4])

	} else if fmt.Sprintf("%X", str[0]) == "79" && fmt.Sprintf("%X", str[1]) == "79" {
		//crc = Crc16(strings.Join(str[4:inLen-4], ""))
	} else {
		return false, errors.New("Package head is not right")
	}
	crcString := fmt.Sprintf("crc is %d", crc)
	//fmt.Printf("%s checksum is %d %X %X\n", crcString, strToI(fmt.Sprintf("%X", str[inLen-4]))*256+strToI(fmt.Sprintf("%X", str[inLen-3])), str[inLen-4], str[inLen-3])
	if int(crc) != StrToI(fmt.Sprintf("%X", str[inLen-4]))*256+StrToI(fmt.Sprintf("%X", str[inLen-3])) {
		return false, errors.New("Package checkSum error " + crcString)
	}
	return true, nil
}

//检查报文长度 长度=协议号+信息内容+信息序列号+错误校验
func checkPackageLength(str []string) (bool, error) {
	messaglen := len(str)
	var inLen int
	if str[0] == "78" && str[1] == "78" {
		inLen = StrToI(str[2])
	} else if str[1] == "79" && str[0] == "79" {
		inLen = StrToI(str[2])*256 + StrToI(str[3])
		messaglen--
	} else {
		return false, errors.New("Package head is not right")
	}
	if inLen != messaglen-5 {
		return false, errors.New("Package length is not right")
	}

	return true, nil
}

//检查报文停止位
func checkStopBits(str []string) (bool, error) {
	len := len(str)
	if len <= 10 {
		return false, errors.New("Package length is not enough")
	}
	if str[len-2] != "0D" || str[len-1] != "0A" {
		return false, errors.New("Package stop byte should be [0d 0a]")
	}
	return true, nil
}

//报文头格式
// 起始位 2
// 包长度 1(2)
// 协议号 1
// 信息内容 N
// 信息序列号 2
// 错误校验 2
// 停止位 2
// 十六进制 0x78 0x78(包长度 1 位) 或 0x79 0x79(包长度 2 位)
func doPackage(deviceId string, load []byte) {

	var text string = TranferBytearrayToString(load)
	var message []string
	//把字符串除掉：，然后按字节分配到int或者byte中
	//text := "78:78:22:22:12:04:10:0f:3b:3b:cc:03:30:3b:b8:0c:e7:13:90:27:d4:52:01:cc:00:68:54:00:60:59:01:00:00:02:93:f3:b4:0d:0a"
	glog.Info(text)
	glog.Flush()
	// os.Exit(1)
	for _, str := range strings.Split(text, ":") {
		// some, err := strconv.ParseInt(str, 16, 0)
		// if err != nil {
		// 	panic("结束")
		// }
		message = append(message, str)

	}
	//fmt.Println(message)
	//TODO 校验报文 checksumMsg 如果有错误，就丢弃报文
	if !checkMsg(message, load) {
		return
	}

	//fmt.Println(message)
	var tempStr []string
	if message[0] == "78" {
		tempStr = message[3:]

	} else if message[0] == "79" {
		tempStr = message[4:]
	} else {
		tempStr = message
	}
	// if tempStr[0] != "22" && tempStr[0] != "13" {
	// 	fmt.Println(deviceId, " ", tempStr[0])
	// }

	var singleMessage messageMqtt
	switch tempStr[0] {
	case "01": //登陆包 GV25 & GT740 &S17 &G17
		//TODO 78:78:11:01:08:68:12:01:77:75:51:69:20:04:32:02:00:00:DB:6A:0D:0A 这个类型码是13 14字节 20 04
		//20 04 类型码不知道是那种设备？？
		singleMessage = new(package01) //完成
		// glog.Info("enter do_login,device_id=" + deviceId + ",device_model=" + strToI(tempStr[9:11]) + ",utf_lang=" + tempStr[11:13])
		//TODO 完善日志信息
		glog.Info("enter do_login,device_id=" + deviceId + ",device_model=")

	case "13", "23": //GV25 心跳包 13 //GT740 心跳包 23
		singleMessage = new(package13And23) //完成
		//[78 78 0A 13 04 06 03 00 01 01 B4 5E 94 0D 0A]

	case "19": //LBS、状态合并包 G740.pdf P23
		singleMessage = new(package19)
	case "21": //GV25  //终端指令回复
		//TODO 不是很理解这种包的分析
		singleMessage = new(package21)
		singleMessage.dealPackage(deviceId, tempStr)
	case "22": //定位数据 GV25 & GT740
		singleMessage = new(package22) //完成

	// // [78 78 22 22 12 05 04 09 2A 22 CF 02 E5 E1 A8 0B 36 3E D0 37 15 14 01 CC 00 85 70 00 0A 50 00 02 00 07 C4 BD C5 0D 0A]
	// case "23":
	// 	singleMessage = new(heartBeatMessage) //完成 和case 13完全一样 TODO修改
	// 	singleMessage.dealPackage(deviceId, tempStr)
	case "26": //报警数据 GT740 & GV25
		singleMessage = new(package26)
	//fmt.Println(text)
	case "28": //LBS 多基站扩展信息包 GT740
		singleMessage = new(package28)
		fmt.Println("28  ", text)
	// case "29": // 未知类型
	// case "2C": //
	// 	singleMessage = new(stationAndWifiMessage)
	// singleMessage.dealPackage(deviceId, tempStr)
	// fmt.Printf("There is a message type is 2C %s \n", text)
	//TODO 没发现这种类型的消息

	case "8A": //校时包 GT740
		//TODO 需要发送时间字节 6各字节数据 yy:MM:dd:HH:mm:ss
	case "80": //	0x80  GT740 & GV25 //服务器向终端发送指令信息
	// case "94": //	暂时没具体内容   //信息传输(终端向服务器上传数据)
	default:
	}
	// singleMessage.dealPackage(deviceId, tempStr)
}
