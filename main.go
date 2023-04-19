package main

import (
	"fmt"
	"os"

	commands "github.com/AF250329/tfs2pdf/cmd"
)

func main() {
	err := commands.Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error occurred while running application ! Error is: %s", err)
		os.Exit(1)
	}
}
