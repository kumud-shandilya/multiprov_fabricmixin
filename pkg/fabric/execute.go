package fabric

import (
	"context"
	"fmt"

	"get.porter.sh/porter/pkg/exec/builder"
	"gopkg.in/yaml.v2"
)

func (m *Mixin) loadAction(ctx context.Context) (*Action, error) {
	var action Action
	fmt.Fprintln(m.Out, action)
	err := builder.LoadAction(ctx, m.RuntimeConfig, "", func(contents []byte) (interface{}, error) {
		fmt.Fprintln(m.Out, contents)
		err := yaml.Unmarshal(contents, &action)
		fmt.Fprintln(m.Out, action)
		fmt.Fprintln(m.Out, contents)
		return &action, err
	})
	return &action, err
}

func (m *Mixin) Execute(ctx context.Context) error {

	fmt.Fprintln(m.Out, "Starting deployment operations...")

	action, err := m.loadAction(ctx)
	if err != nil {
		return err
	}

	fmt.Fprintln(m.Out, action)
	//_, err = builder.ExecuteSingleStepAction(ctx, m.RuntimeConfig, action)
	return nil
}
