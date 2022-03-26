package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"oj/apps"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	var cmd string

	app := &cli.App{
		Name:  "oj",
		Usage: "convert the output of linux cli to JSON",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "command",
				Aliases:     []string{"cmd"},
				Value:       "env",
				Usage:       "data output by `COMMAND`",
				Destination: &cmd,
			},
		},
		Before: func(c *cli.Context) error {
			fi, err := os.Stdin.Stat()
			if err != nil {
				panic(err)
			}

			if fi.Mode()&os.ModeNamedPipe == 0 {
				return fmt.Errorf("You need to pipe information through oj.")
			}
			return nil
		},
		Action: func(c *cli.Context) error {

			var app *App
			if cmd == "env" {
				app = NewApp(apps.NewEnv())
			} else if cmd == "ps" {
				app = NewApp(apps.NewPS())
			} else {
				fmt.Println("Hello")
			}

			reader := bufio.NewReader(os.Stdin)
			var output []rune

			for {
				input, _, err := reader.ReadRune()
				if err != nil && err == io.EOF {
					break
				}
				output = append(output, input)
			}

			var jsonOutput = app.parse(string(output))
			fmt.Println(jsonOutput)

			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
