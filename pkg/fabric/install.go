package fabric

import (
	"bufio"
	"context"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/go-resty/resty/v2"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

type InstallAction struct {
	Steps []InstallStep `yaml:"install"`
}

type InstallStep struct {
	InstallArguments `yaml:"fabric"`
}

type InstallArguments struct {
	Name  string                 `yaml:"workspace"`
	Token string                 `yaml:"token"`
	Args  map[string]interface{} `yaml:"arguments"`
}

type Arguments struct {
	Type        string `json:"type"`
	DisplayName string `json:"displayName"`
}

func (m *Mixin) getPayloadData() ([]byte, error) {
	reader := bufio.NewReader(m.In)
	data, err := ioutil.ReadAll(reader)
	return data, errors.Wrap(err, "could not read the payload from STDIN")
}

func (m *Mixin) post(workspace_id string, token string, jsonRequest map[string]interface{}) error {

	client := resty.New()
	log.Println("Post called")

	response, _ := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(jsonRequest).
		SetAuthToken(token).
		Post("https://api.fabric.microsoft.com/v1/workspaces/" + workspace_id + "/items")
	log.Println(response)

	return nil
}

func (m *Mixin) Install(ctx context.Context) error {

	fmt.Fprintln(m.Out, "Starting deployment operations...")

	payload, err := m.getPayloadData()
	if err != nil {
		log.Println(err)
		return err
	}

	var action InstallAction
	err = yaml.Unmarshal(payload, &action)
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println(action.Steps[0].Name, action.Steps[0].Args)

	// for index, a := range action.Steps[0].Args {
	// 	fmt.Println(index, a)
	// 	err = m.post(action.Steps[0].Name, action.Steps[0].Token, a)
	// 	if err != nil {
	// 		log.Println(err)
	// 		return err
	// 	}
	// }

	err = m.post(action.Steps[0].Name, action.Steps[0].Token, action.Steps[0].Args)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

// if action.Steps[0].Name == "" || action.Steps[0].Token == "" {
// 	log.Println("Workspace id and access token must be provided")
// 	return errors.New("Wrong input")
// }

// check if the arguments are valid
// var args Arguments
// args_, err := json.Marshal(action.Steps[0].Args)
// err = json.Unmarshal(args_, &args)
// if err != nil {
// 	log.Println(err)
// 	return err
// }

// if args.DisplayName == "" || args.Type == "" {
// 	log.Println("Name and resource type must be provided")
// 	return errors.New("Wrong input")
// }
