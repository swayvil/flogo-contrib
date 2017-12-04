package main

import (
	mqtt "flogo-contrib/activity/mqtt-client"
)

func main() {
	mqttClient := mqtt.NewMQTTClient("tcp://localhost:1883", "iot-client", "0")
	mqttClient.Subscribe("test/topic")
	mqttClient.Publish("test/topic", "Hello World!")
	mqttClient.Disconnect()
}