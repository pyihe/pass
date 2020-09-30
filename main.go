package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	flags := flag.NewFlagSet("pass", flag.ExitOnError)
	flags.Usage = func() {
		showHelp()
	}

	flags.StringVar(&c.command, "c", "", "executable tag")
	flags.StringVar(&c.key, "k", "", "key")
	flags.StringVar(&c.pass, "p", "", "pass")
	if err := flags.Parse(os.Args[1:]); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		flags.Usage()
		return
	}

	if err := initFile(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		flags.Usage()
		return
	}

	if err := c.execCommand(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		flags.Usage()
		return
	}
}
