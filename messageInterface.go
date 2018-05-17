package main

type messageMqtt interface {
	dealPackage(deviceId string, str []string)
}
