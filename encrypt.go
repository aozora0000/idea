package idea

import (
	"errors"
	"fmt"
	"github.com/urfave/cli/v2"
	"os"
	"path"
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
	if context.Bool("with-command") {
		fmt.Printf("%s decrypt --psyche %s %s", path.Base(os.Args[0]), context.String("psyche"), encrypted)
		return nil
	}
	fmt.Println(encrypted)
	return nil
}
