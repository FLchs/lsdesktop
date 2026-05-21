package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/FLchs/lsdesktop/internal/desktop"
	"github.com/FLchs/lsdesktop/internal/history"
)

func usage() {
	fmt.Fprintln(os.Stderr, "Usage: desklaunch [-p] <path>")
	flag.PrintDefaults()
}

func updateHistory(name, path string) error {
	entry := desktop.EntryString(name, path)

	hist, err := history.ReadHistory()
	if err != nil {
		return fmt.Errorf("read history: %w", err)
	}

	lineCount := strings.Count(hist, "\n")
	out := make([]string, 0, lineCount+2)
	out = append(out, entry)

	for x := range strings.SplitSeq(hist, "\n") {
		if x == entry {
			continue
		}
		out = append(out, x)
	}

	historyContent := strings.Join(out, "\n")
	if err := history.WriteHistory(historyContent); err != nil {
		return fmt.Errorf("write history: %w", err)
	}

	return nil
}

func main() {
	printOnly := flag.Bool("p", false, "print the exec command instead of running it")
	flag.Usage = usage
	flag.Parse()

	args := flag.Args()
	if len(args) < 1 {
		usage()
		os.Exit(1)
	}

	path := args[0]

	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: read desktop entry: %v\n", err)
		os.Exit(1)
	}

	var name, execLine, tryExec, workDir string

	for line := range strings.Lines(string(data)) {
		line = strings.TrimSpace(line)

		if name == "" {
			if v, ok := strings.CutPrefix(line, "Name="); ok {
				name = strings.TrimSpace(v)
			}
		}
		if execLine == "" {
			if v, ok := strings.CutPrefix(line, "Exec="); ok {
				execLine = strings.TrimSpace(v)
			}
		}
		if tryExec == "" {
			if v, ok := strings.CutPrefix(line, "TryExec="); ok {
				tryExec = strings.TrimSpace(v)
			}
		}
		if workDir == "" {
			if v, ok := strings.CutPrefix(line, "Path="); ok {
				workDir = strings.TrimSpace(v)
			}
		}
	}

	if execLine == "" {
		fmt.Fprintf(os.Stderr, "Error: no Exec line found in %s\n", path)
		os.Exit(1)
	}

	if err := updateHistory(name, path); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	if *printOnly {
		fmt.Println(desktop.CleanExec(execLine))
		return
	}

	cmd, err := desktop.BuildCmd(execLine, workDir, tryExec)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: build command: %v\n", err)
		os.Exit(1)
	}

	if err := cmd.Start(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: start command: %v\n", err)
		os.Exit(1)
	}
}
