package mqttclient

import (
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)

var log = logger.GetLogger("activity-mqtt-client")

// MQTTClientActivity is a stub for your Activity implementation
type MQTTClientActivity struct {
	metadata   *activity.Metadata
	mqttClient *MQTTClient
}

// NewActivity creates a new activity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &MQTTClientActivity{metadata, nil}
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
	qos := context.GetInput("qos").(string)
	topic := context.GetInput("topic").(string)
	msg := context.GetInput("message").(string)
	
	a.mqttClient = NewMQTTClient(brokerUrl, clientId, qos)
	a.mqttClient.Publish(topic, msg)

	log.Debugf("Message published on [%s] topic, [%s] MQTT broker", topic, brokerUrl)

	// Set the result as part of the context
	context.SetOutput("result", "OK")

	// Signal to the Flogo engine that the activity is completed
	return true, nil
}
