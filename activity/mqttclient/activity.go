package mqttclient

import (
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"strconv"
)

var log = logger.GetLogger("activity-mqttclient")

// MQTTClientActivity is a stub for your Activity implementation
type MQTTClientActivity struct {
	metadata *activity.Metadata
}

// NewActivity creates a new activity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &MQTTClientActivity{metadata: metadata}
}

// Metadata implements activity.Activity.Metadata
func (a *MQTTClientActivity) Metadata() *activity.Metadata {
	return a.metadata
}

// Eval implements activity.Activity.Eval
func (a *MQTTClientActivity) Eval(context activity.Context) (done bool, err error) {
	// Get the activity data from the context
	brokerUrl := context.GetInput("brokerUrl").(string)
	clientId := context.GetInput("clientId").(string)
	topic := context.GetInput("topic").(string)
	msg := context.GetInput("message").(string)
	qos, err := strconv.Atoi(context.GetInput("qos").(string))
	if err != nil {
		log.Error("Error converting \"qos\" to an integer ", err.Error())
		return err
	}

	a.publish(brokerUrl, clientId, qos, topic, msg)

	log.Debugf("Message published on [%s] topic, [%s] MQTT broker", topic, brokerUrl)

	// Set the result as part of the context
	context.SetOutput("result", "OK")

	// Signal to the Flogo engine that the activity is completed
	return true, nil
}

// *************
// MQTT specific
// *************

func (a *MQTTClientActivity) publish(brokerUrl string, clientId string, qos int, topic string, msg string) {
	opts := mqtt.NewClientOptions().AddBroker(brokerUrl)
	opts.SetClientID(clientId)

	// Create and start a client using the above ClientOptions
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Error(token.Error())
	}

	token := client.Publish(topic, byte(qos), false, msg)
	token.Wait()

	// Disconnect
	if token := client.Unsubscribe(topic); token.Wait() && token.Error() != nil {
		log.Error(token.Error())
	}
	client.Disconnect(250)
}
