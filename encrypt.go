package idea

import (
	"errors"
	"fmt"
	"github.com/urfave/cli/v2"
)

func Encrypt(context *cli.Context) error {
	if context.Args().Len() != 1 {
		return errors.New("input invalid")
	}
	pool, err := LoadPsyche(context.String("psyche"))
	if err != nil {
		return err
	}
	encrypted, err := NewCipher(pool.Key).Encrypt([]byte(context.Args().Get(0)))
	if err != nil {
		return err
	}
	fmt.Println(encrypted)
	return nil
}
