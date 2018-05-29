package main

import (
	"flag"
	"fmt"
	"fun/cli"
	"os"
	"time"
)

const (
	minRevistiDelay = 10 * time.Second
)

var (
	revisitEvery = flag.Duration("revisit-every", -1, "")
	showLicense  = flag.Bool("license", false, "")
	showL        = flag.Bool("l", false, "")
	showHelp     = flag.Bool("help", false, "")
	showH        = flag.Bool("h", false, "")
	urls         []string
)

func init() {
	if os.Getuid() == 0 {
		fmt.Println("Do not run as ROOT")
		os.Exit(2)
	}
	handleCommand()
}

func handleCommand() {
	flag.Parse()

	if *showH || *showHelp {
		usage()
		os.Exit(0)
	}

	if *showL || *showLicense {
		cli.PrintLogo("")
		cli.PrintLicense("CoderMe.com", "")
		os.Exit(0)
	}

	if flag.NArg() == 0 {
		usage()
		os.Exit(2)
	}

	urls = flag.Args()

	if *revisitEvery < minRevistiDelay {
		*revisitEvery = 0
	}

}

func usage() {
	fmt.Printf(`
Usage: %s [-h | --help] [-l | -license] [-revisit-every duration] url...

flags:
========
 -h | --help
  Display this help and exit

 -l | --license
  Display license and exit.


Options:
========
 -revisit-every duration
  Revisit page every certain duration of time, minimum duration = %s, -1 means don't revisit (default).

Args:
========
 url...
  URL(s) to be fetch.



`,
		os.Args[0],
		minRevistiDelay,
	)
}
