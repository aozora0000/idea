package idea

import (
	"errors"
	"fmt"
	"github.com/urfave/cli/v2"
)

func Decrypt(context *cli.Context) error {
	if context.Args().Len() != 1 {
		return errors.New("input invalid")
	}
	pool, err := LoadPsyche(context.String("psyche"))
	if err != nil {
		return err
	}
	raw, err := NewCipher(pool.Key).Decrypt(context.Args().Get(0))
	if err != nil {
		return err
	}
	fmt.Println(raw)
	return nil
}
