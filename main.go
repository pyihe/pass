package main

import (
	"flag"
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
	var err error
	if err = flags.Parse(os.Args[1:]); err != nil {
		flags.Usage()
		return
	}

	if err = initFile(); err != nil {
		flags.Usage()
		return
	}

	if err = c.execCommand(); err != nil {
		flags.Usage()
		return
	}
}
