package main

import (
	code "code/src"
	"fmt"
)

func main() {
	// (&cli.Command{}).Run(context.Background(), os.Args)
	path := "/Users/konstantin/Programming/go-project-242/src"
	result, err := code.GetSize(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%d	%s", result, path)
}
