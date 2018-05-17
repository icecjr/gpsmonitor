package main

import (
	"fmt"
	"log"
	"os"

	"github.com/eclipse/paho.mqtt.golang"
	//"strconv"
	//"strings"
	"time"
)

var f mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("TOPIC: %s\n", msg.Topic())
	fmt.Printf("MSG: %s\n", msg.Payload())
}

func main() {
	//mqtt.DEBUG = log.New(os.Stdout, "", 0)
	mqtt.ERROR = log.New(os.Stdout, "", 0)
	// opts := mqtt.NewClientOptions().AddBroker("tcp://iot.eclipse.org:1883").SetClientID("gotrivial")
	opts := mqtt.NewClientOptions().AddBroker("tcp://172.16.2.3:1883").SetClientID("gotrivial")

	opts.SetKeepAlive(2 * time.Second)
	opts.SetDefaultPublishHandler(f)
	opts.SetPingTimeout(1 * time.Second)

	c := mqtt.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	// if token := c.Subscribe("/v1/d/+/m", 0, nil); token.Wait() && token.Error() != nil {
	// 	fmt.Println(token.Error())
	// 	os.Exit(1)
	// }

	for i := 0; i < 5; i++ {
		//text := fmt.Sprintf("this is msg #%d!", i)
		text := "78:78:22:22:12:04:10:0f:3b:3b:cc:03:30:3b:b8:0c:e7:13:90:27:d4:52:01:cc:00:68:54:00:60:59:01:00:00:02:93:f3:b4:0d:0a"
		// var msg []byte
		// for i, str := range strings.Split(text, ":") {
		// 	msg[i] = strconv.ParseInt(str, 16, 2)
		// 	i++
		// }
		token := c.Publish("/v1/d/868120174576584/m", 0, false, text)
		token.Wait()
		time.Sleep(6 * time.Second)
	}

	time.Sleep(6 * time.Second)

	// if token := c.Unsubscribe("/v1/d/+/m"); token.Wait() && token.Error() != nil {
	// 	fmt.Println(token.Error())
	// 	os.Exit(1)
	// }

	c.Disconnect(250)

	time.Sleep(1 * time.Second)
}
