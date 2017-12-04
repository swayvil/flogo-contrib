package mqttclient

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type MQTTClient struct {
	client    mqtt.Client
	brokerUrl string
	clientId  string
	qos       byte
	topic     string
}

// Default message handler function
//func msgHandler(client mqtt.Client, msg mqtt.Message) {
//	log.Infof("TOPIC: %s\n", msg.Topic())
//	log.Infof("MSG: %s\n", msg.Payload())
//}

func NewMQTTClient(brokerUrl string, clientId string, qos byte) *MQTTClient {
	opts := mqtt.NewClientOptions().AddBroker(brokerUrl)
	opts.SetClientID(clientId)
//	opts.SetDefaultPublishHandler(msgHandler)

	// Create and start a client using the above ClientOptions
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	return &MQTTClient{client, brokerUrl, brokerUrl, qos, ""}
}

//func (mqttClient *MQTTClient) Subscribe(topic string) {
//	if token := mqttClient.client.Subscribe(topic, mqttClient.qos, nil); token.Wait() && token.Error() != nil {
//		log.Error(token.Error())
//	}
//}

func (mqttClient *MQTTClient) Publish(topic string, msg string) {
	mqttClient.topic = topic
	token := mqttClient.client.Publish(topic, mqttClient.qos, false, msg)
	token.Wait()
}

func (mqttClient *MQTTClient) Disconnect() {
	if mqttClient.topic != "" {
		if token := mqttClient.client.Unsubscribe(mqttClient.topic); token.Wait() && token.Error() != nil {
			log.Error(token.Error())
		}
	}
	mqttClient.client.Disconnect(250)
}
