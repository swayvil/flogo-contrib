package ds18b20

import (
	"github.com/TIBCOSoftware/flogo-lib/core/action"
	"github.com/TIBCOSoftware/flogo-lib/core/trigger"
	"github.com/TIBCOSoftware/flogo-lib/logger"
	"time"
	"strconv"
	"github.com/yryz/ds18b20"
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
	log.Info("Starting ds18b20 Trigger")
	
	sleepDuration, err := strconv.Atoi(t.config.GetSetting("sleepDuration"))
	if err != nil {
		log.Error("Error converting \"sleepDuration\" to an integer ", err.Error())
		return err
	}

	for true {
		t.RunAction(t.getTemperature())
		time.Sleep(time.Duration(sleepDuration) * time.Second)
	}
	return nil
}

// RunAction starts a new Process Instance
func (t *DS18b20Trigger) RunAction(temperature string) {
	req := t.constructStartRequest(temperature)
	startAttrs, _ := t.metadata.OutputsToAttrs(req.Data, false)

	actionId := singleton.config.Handlers[0].ActionId
	act := action.Get(actionId)
	ctx := trigger.NewInitialContext(startAttrs, singleton.config.Handlers[0])
	_, err := singleton.runner.RunAction(ctx, act, nil)

	if err != nil {
		log.Debugf("ds18b20 Trigger Error: %s", err.Error())
	}
}

// Stop implements trigger.Trigger.Start
func (t *DS18b20Trigger) Stop() error {
	// stop the trigger
	return nil
}

func (t *DS18b20Trigger) constructStartRequest(temperature string) *StartRequest {
	req := &StartRequest{}
	data := make(map[string]interface{})
	data["temperature"] = temperature
	req.Data = data
	return req
}

// StartRequest describes a request for starting a ProcessInstance
type StartRequest struct {
	ProcessURI string                 `json:"flowUri"`
	Data       map[string]interface{} `json:"data"`
}

// *************
// ds18b20 specific
// *************

func (t *DS18b20Trigger) getTemperature() string {
		sensors, err := ds18b20.Sensors()
		if err != nil {
			log.Error(err)
			panic(err)
		}
	
		for _, sensor := range sensors {
			t, err := ds18b20.Temperature(sensor)
			if err == nil {
				log.Debugf("Sensor: %s temperature: %.2f°C\n", sensor, t)
				return strconv.FormatFloat(t, 'E', 4, 64)
			} else {
				log.Error(err)
			}
		}
	return "0"
}
