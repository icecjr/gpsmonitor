package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var ch chan int
var i int
var f mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	//fmt.Printf("TOPIC: %s\n", msg.Topic())
	// fmt.Printf("receive %d", i)
	// i++
	//fmt.Printf("MSG: %s\n", msg.Payload())
	deviceId := pickDeviceID(msg.Topic())
	// if i%100 == 0 {
	// 	fmt.Println(i)
	// }

	// i++
	//fmt.Println(deviceId)
	//os.Exit(1)
	// //更新session管理列表timeUpdate.go
	//updateSessonTime(deviceId)
	//byte[] -> string
	// n := bytes.IndexByte(msg.Payload(), 0)
	// s := string(msg.Payload()[:n])
	doPackage(deviceId, msg.Payload())
}

func ConnectMqtt() {
	opts := mqtt.NewClientOptions().AddBroker(mqttConfig.Mqtt_broker).SetClientID(mqttConfig.ClientID)
	opts.SetWill(mqttConfig.Topic, "Goodbye", 1, true)
	//	opts.SetKeepAlive(20 * time.Second)
	opts.SetDefaultPublishHandler(f)
	//	opts.SetPingTimeout(1 * time.Second)
	opts.SetProtocolVersion(4)
	opts.SetCleanSession(false) //这里必须设置，否则接收会阻塞

	c := mqtt.NewClient(opts)
	if !c.IsConnected() {
		if token := c.Connect(); token.Wait() && token.Error() != nil {
			panic(token.Error())
		}

	}
	defer c.Disconnect(250)

	if token := c.Subscribe(mqttConfig.Topic, 0, nil); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}
	//<-ch
	for {
		//checkSessionAlive()
		time.Sleep(6 * time.Second)
	} //循环处理报文，不退出
}

func pickDeviceID(str string) string {
	var strArray []string
	strArray = strings.Split(str, "/")
	//fmt.Println(strArray[3])
	return strArray[3]
}
