package idea

import (
	"errors"
	"fmt"
	"github.com/urfave/cli/v2"
)

func Create(context *cli.Context) error {
	if ExistPsyche(context.String("psyche")) {
		return errors.New("pool exist")
	}
	_, err := CreatePsyche(context.String("psyche"))
	if err == nil {
		fmt.Println("psyche create success")
		return nil
	}
	return err
}
