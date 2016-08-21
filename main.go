package main

import (
	"os"

	"github.com/urfave/cli"
)

func main() {
	cli.NewApp().Run(os.Args)
}
