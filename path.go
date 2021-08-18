package idea

import (
	"fmt"
	"github.com/urfave/cli/v2"
)

func Path(context *cli.Context) error {
	fmt.Println(getShelfDir())
	return nil
}
