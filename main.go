package main

import (
	"flc/lsdesktop/commands"
	"fmt"
	"os"
)

func usage() {
	fmt.Fprintln(os.Stderr, "Usage: lsdesktop [list|launch <path>|print <path>]")
}

func main() {
	args := os.Args[1:]

	command := ""

	if len(args) > 0 {
		command = args[0]
	}

	switch command {
	case "list", "":
		commands.List()
	case "launch":
		if len(args) < 2 {
			usage()
			os.Exit(1)
		}
		if err := commands.Launch(args[1]); err != nil {
			fmt.Fprintln(os.Stderr, "Error:", err)
			os.Exit(1)
		}
	case "print":
		if len(args) < 2 {
			usage()
			os.Exit(1)
		}
		if err := commands.Print(args[1]); err != nil {
			fmt.Fprintln(os.Stderr, "Error:", err)
			os.Exit(1)
		}
	case "-h", "--help":
		usage()
	default:
		fmt.Fprintf(os.Stderr, "Unknown command: %q\n", args[0])
		usage()
		os.Exit(1)
	}
}
