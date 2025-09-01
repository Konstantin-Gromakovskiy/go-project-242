package main

import (
	"code"
	"context"
	"fmt"
	"os"

	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:  "human",
				Usage: "human-readable sizes (auto-select unit)",
			},
			&cli.BoolFlag{
				Name:    "all",
				Aliases: []string{"a"},
				Usage:   "include hidden files and directories",
			},
			&cli.BoolFlag{
				Name:    "recursive",
				Aliases: []string{"r"},
				Usage:   "recursive size of directories",
			},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			if cmd.NArg() == 0 {
				return cli.ShowAppHelp(cmd)
			}
			humanFlag := cmd.Bool("human")
			hiddenFlag := cmd.Bool("all")
			recursiveFlag := cmd.Bool("recursive")
			path := cmd.Args().First()
			result, err := code.GetSize(path, hiddenFlag, recursiveFlag)

			if err != nil {
				fmt.Print(err)
			}

			fmt.Printf("%s	%s", code.FormatSize(result, humanFlag), path)
			return nil
		},
	}

	err := cmd.Run(context.Background(), os.Args)
	if err != nil {
		fmt.Println(err)
	}
}
