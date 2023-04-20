package main

import (
	"fmt"
	"os"

	commands "github.com/AF250329/tfs2pdf/cmd"
)

func main() {
	err := commands.Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error occurred while running application !\nError is:\t%s", err)
		os.Exit(1)
	}
}
