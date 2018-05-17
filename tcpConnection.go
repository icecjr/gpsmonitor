package main

import (
	"net"
	. "project/elangshen/mutils"

	"github.com/golang/glog"
)

func CreateTCPConnection() {
	netListen, err := net.Listen("tcp", "localhost:6060")
	CheckError(err)
	defer netListen.Close()
	glog.Info("Waiting for clients")
	for {
		conn, err := netListen.Accept()
		if err != nil {
			continue
		}

		glog.Info(conn.RemoteAddr().String(), " tcp connect success")
		go handleConnection(conn)

	}
}
func handleConnection(conn net.Conn) {
	// 缓冲区，存储被截断的数据
	//tmpBuffer := make([]byte, 0)
	defer conn.Close()
	//接收解包
	var deviceId string
	buffer := make([]byte, 1024)
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			glog.Errorf(conn.RemoteAddr().String()+" connection error: ", err)
			return
		}
		glog.Info(string(buffer[0:n]))
		if deviceId == "" && n >= 22 {
			if buffer[0] == 78 && buffer[1] == 78 && buffer[3] == 01 {
				deviceId = string(buffer[4:12])
			}
		}
		if deviceId != "" {
			handleMessage(buffer[0:n])
		}

	}

}

//TODO 这里需要修改参数，变成一个报文处理一次
func handleMessage(buffer []byte) {

}
