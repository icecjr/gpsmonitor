package main

import (
	"github.com/antlinker/go-mqtt/client"
	"github.com/antlinker/go-mqtt/event"
)

//连接事件接口
type MqttDisConnListener interface {
	event.Listener
	OnDisconning(event *client.MqttEvent)
	OnDisconned(event *client.MqttEvent)
	OnLostConn(event *client.MqttEvent, err error)
}

//默认连接事件实现，可以打印事件信息
//事件接口
type MqttSubListener interface {
	event.Listener
	OnSubStart(event *client.MqttEvent, sub []client.SubFilter)
	OnSubSuccess(event *client.MqttEvent, sub []client.SubFilter, result []client.QoS)
}

//默认事件实现，可以打印事件信息
type DefaultPrintSubscribeListen struct {
}

func (*DefaultPrintSubscribeListen) OnSubStart(event *client.MqttEvent, sub []client.SubFilter) {
	client.Mlog.Debugf("OnSubStart:%v", sub)
}
func (*DefaultPrintSubscribeListen) OnSubSuccess(event *client.MqttEvent, sub []client.SubFilter, result []client.QoS) {
	client.Mlog.Debugf("OnSubSuccess:%v:%v", sub, result)
}

//默认事件实现，不做任何事情
type DefaultSubscribeListen struct {
}

func (*DefaultSubscribeListen) OnSubStart(event *client.MqttEvent, sub []client.SubFilter) {
}
func (*DefaultSubscribeListen) OnSubSuccess(event *client.MqttEvent, sub []client.SubFilter, result []client.QoS) {
}

// func main() {
// 	//开始日志　false则关闭日志显示
// 	client.Mlog.SetEnabled(true)
// 	client1, err := client.CreateClient(client.MqttOption{
// 		Addr: "tcp://172.16.2.3:1883",
// 		//断开连接１秒后自动连接，０不自动重连
// 		ReconnTimeInterval: 1,
// 	})
// 	if err != nil {
// 		//配置文件解析失败
// 		panic("配置文件解析失败")
// 	}
// 	listener := &DefaultPrintSubscribeListen{}
// 	//注册订阅事件
// 	client.AddSubListener(listener)
// 	//建立连接
// 	err := client.Connect()
// 	if err != nil {
// 		//连接失败，不会进入自动重连状态
// 		panic(fmt.Errorf("连接失败:%v", err))
// 	}
// 	mq, err := client.Subscribe("Test/1", client.QoS1)
// 	if err != nil {
// 		//订阅失败
// 		panic(fmt.Errorf("订阅失败:%v", err))
// 	}

// 	//等待订阅成功
// 	mq.Wait()
// 	if mq.Err() != nil {
// 		//订阅失败
// 		panic(fmt.Errorf("订阅失败:%v", client.mqt.Err()))
// 	}
// 	mq, err = client.Subscribes(client.CreateSubFilter("test/1", client.QoS1),
// 		client.CreateSubFilter("test/2", client.QoS2),
// 		client.CreateSubFilter("test/0", client.QoS0))
// 	if err != nil {
// 		//订阅失败
// 		panic(fmt.Errorf("订阅失败:%v", err))
// 	}
// 	//等待订阅成功
// 	mq.Wait()
// 	if mq.Err() != nil {
// 		//订阅失败
// 		panic(fmt.Errorf("订阅失败:%v", mqt.Err()))
// 	}

// 	//移除事件
// 	client.RemoveSubListener(listener)
// 	fmt.Println("Create connect successfully")

// }
