package fabric

import (
	"context"
	"fmt"
	"log"

	"github.com/go-resty/resty/v2"
	"gopkg.in/yaml.v2"
)

type UpgradeAction struct {
	Steps []UpgradeStep `yaml:"upgrade"`
}

type UpgradeStep struct {
	UpgradeArguments `yaml:"fabric"`
}

type UpgradeArguments struct {
	WorkspaceId string                 `yaml:"workspace"`
	Token       string                 `yaml:"token"`
	ItemId      string                 `yaml:"item"`
	Args        map[string]interface{} `yaml:"arguments"`
}

func (m *Mixin) upgrade(workspace_id string, token string, item_id string, jsonRequest map[string]interface{}) error {

	client := resty.New()
	log.Println("Post called")

	response, _ := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(jsonRequest).
		SetAuthToken(token).
		Patch("https://api.fabric.microsoft.com/v1/workspaces/" + workspace_id + "/items/" + item_id)
	log.Println(response)

	return nil
}

func (m *Mixin) Upgrade(ctx context.Context) error {
	fmt.Println("Hi")

	fmt.Println("Hi")

	fmt.Fprintln(m.Out, "Starting upgrade operations...")

	payload, err := m.getPayloadData()
	if err != nil {
		log.Println(err)
		return err
	}

	var action UpgradeAction
	err = yaml.Unmarshal(payload, &action)
	if err != nil {
		log.Println(err)
		return err
	}

	log.Println(action.Steps[0].WorkspaceId, action.Steps[0].ItemId, action.Steps[0].Args)

	err = m.upgrade(action.Steps[0].WorkspaceId, action.Steps[0].Token, action.Steps[0].ItemId, action.Steps[0].Args)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
