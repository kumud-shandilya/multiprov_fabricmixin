package fabric

import (
	"context"
	"fmt"
	"log"

	"github.com/go-resty/resty/v2"
	"gopkg.in/yaml.v2"
)

type InvokeAction struct {
	Steps []InvokeStep `yaml:"invoke"`
}

type InvokeStep struct {
	InvokeArguments `yaml:"fabric"`
}

type InvokeArguments struct {
	WorkspaceId string `yaml:"workspace"`
	Token       string `yaml:"token"`
	ItemId      string `yaml:"item"`
}

func (m *Mixin) get(workspace_id string, token string, item_id string) error {

	client := resty.New()
	log.Println("Get called")

	response, _ := client.R().
		SetHeader("Content-Type", "application/json").
		SetAuthToken(token).
		Get("https://api.fabric.microsoft.com/v1/workspaces/" + workspace_id + "/items/" + item_id)
	log.Println(response)

	return nil
}

func (m *Mixin) Invoke(ctx context.Context) error {
	fmt.Println("Hi")

	fmt.Fprintln(m.Out, "Starting uninstall operations...")

	payload, err := m.getPayloadData()
	if err != nil {
		log.Println(err)
		return err
	}

	var action InvokeAction
	err = yaml.Unmarshal(payload, &action)
	if err != nil {
		log.Println(err)
		return err
	}

	log.Println(action.Steps[0].WorkspaceId, action.Steps[0].ItemId)

	err = m.get(action.Steps[0].WorkspaceId, action.Steps[0].Token, action.Steps[0].ItemId)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
