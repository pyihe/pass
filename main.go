package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	switch len(os.Args) {
	case 1:
		showHelp()
		return
	case 2:
		h := os.Args[1]
		h = strings.ToLower(h)
		if h == "-h" || h == "-help" {
			showHelp()
			return
		}
		c.command = os.Args[1]
	case 3:
		c.command = os.Args[1]
		c.key = os.Args[2]
	case 4:
		c.command = os.Args[1]
		c.key = os.Args[2]
		c.pass = os.Args[3]
	default:
		return
	}

	if err := initFile(); err != nil {
		fmt.Fprintf(os.Stdout, "%v\n", err)
		return
	}

	if err := c.execCommand(); err != nil {
		fmt.Fprintf(os.Stdout, "%v\n", err)
		return
	}
}
