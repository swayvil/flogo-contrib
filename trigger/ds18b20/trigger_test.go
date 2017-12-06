package ds18b20

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/TIBCOSoftware/flogo-lib/core/action"
	"github.com/TIBCOSoftware/flogo-lib/core/trigger"
	"github.com/TIBCOSoftware/flogo-lib/core/data"
)

func getJsonMetadata() string {
	jsonMetadataBytes, err := ioutil.ReadFile("trigger.json")
	if err != nil {
		panic("No Json Metadata found for trigger.json path")
	}
	return string(jsonMetadataBytes)
}

type TestRunner struct {
}

// Run implements action.Runner.Run
func (tr *TestRunner) Run(context context.Context, action action.Action, uri string, options interface{}) (code int, data interface{}, err error) {
	return 0, nil, nil
}

func (tr *TestRunner) RunAction(ctx context.Context, act action.Action, options map[string]interface{}) (results map[string]*data.Attribute, err error) {
	return nil, nil
}

const testConfig string = `{
  "id": "ds18b20trigger",
  "settings": {
  },
  "handlers": [
    {
      "actionId": "test_action",
      "settings": {
        "handler_setting": "somevalue"
      }
    }
  ]
}`

func TestInit(t *testing.T) {
	// New factory
	md := trigger.NewMetadata(getJsonMetadata())
	f := NewFactory(md)

	// New Trigger
	config := trigger.Config{}
	json.Unmarshal([]byte(testConfig), config)
	tgr := f.New(&config)

	runner := &TestRunner{}

	tgr.Init(runner)
}
