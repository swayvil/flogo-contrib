package ds18b20

import (
	"github.com/TIBCOSoftware/flogo-lib/core/action"
	"github.com/TIBCOSoftware/flogo-lib/core/trigger"
	"github.com/TIBCOSoftware/flogo-lib/logger"
	//"github.com/yryz/ds18b20"
)

var log = logger.GetLogger("trigger-ds18b20")
var singleton *DS18b20Trigger

// DS18b20TriggerFactory My Trigger factory
type DS18b20TriggerFactory struct {
	metadata *trigger.Metadata
}

//NewFactory create a new Trigger factory
func NewFactory(md *trigger.Metadata) trigger.Factory {
	return &DS18b20TriggerFactory{metadata: md}
}

//New Creates a new trigger instance for a given id
func (t *DS18b20TriggerFactory) New(config *trigger.Config) trigger.Trigger {
	singleton = &DS18b20Trigger{metadata: t.metadata, config: config}
	return singleton
}

// DS18b20Trigger is a stub for your Trigger implementation
type DS18b20Trigger struct {
	metadata *trigger.Metadata
	runner   action.Runner
	config   *trigger.Config
}

// Init implements trigger.Trigger.Init
func (t *DS18b20Trigger) Init(runner action.Runner) {
	t.runner = runner
}

// Metadata implements trigger.Trigger.Metadata
func (t *DS18b20Trigger) Metadata() *trigger.Metadata {
	return t.metadata
}

// Start implements trigger.Trigger.Start
func (t *DS18b20Trigger) Start() error {
	return nil
}

// Stop implements trigger.Trigger.Start
func (t *DS18b20Trigger) Stop() error {
	// stop the trigger
	return nil
}

func Invoke() (interface{}, error) {
	log.Info("Starting ds18b20 Trigger")

	temperature := getTemperature()
	data := map[string]interface{}{
		"temperature": temperature,
	}

	startAttrs, err := singleton.metadata.OutputsToAttrs(data, false)
	if err != nil {
		log.Errorf("After run error' %s'\n", err)
		return nil, err
	}

	actionId := singleton.config.Handlers[0].ActionId
	log.Debugf("Calling actionid: '%s'\n", actionId)
	act := action.Get(actionId)

	ctx := trigger.NewInitialContext(startAttrs, singleton.config.Handlers[0])
	results, err := singleton.runner.RunAction(ctx, act, nil)

	var replyData interface{}

	if len(results) != 0 {
		dataAttr, ok := results["data"]
		if ok {
			replyData = dataAttr.Value
		}
	}

	if err != nil {
		log.Debugf("ds18b20 Trigger Error: %s", err.Error())
		return nil, err
	}

	return replyData, err
}

// *************
// ds18b20 specific
// *************

func getTemperature() string {
//	sensors, err := ds18b20.Sensors()
//	if err != nil {
//		log.Error(err)
//		panic(err)
//	}
//
//	for _, sensor := range sensors {
//		t, err := ds18b20.Temperature(sensor)
//		if err == nil {
//			log.Debugf("Sensor: %s temperature: %.2f°C\n", sensor, t)
//			return t
//		} else {
//			log.Error(err)
//		}
//	}
	return "42"
}
