package commands

import (
	"flc/lsdesktop/internal/desktop"
	"flc/lsdesktop/internal/history"
	"fmt"
	"os"
	"strings"
)

// Print outputs the exec command to stdout and updates history.
func Print(path string) error {

	// read desktop entry
	data, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("read desktop entry: %w", err)
	}

	var name string
	var exec string

	for line := range strings.Lines(string(data)) {
		line = strings.TrimSpace(line)

		if name == "" {
			if v, ok := strings.CutPrefix(line, "Name="); ok {
				name = strings.TrimSpace(v)
			}
		}
		if exec == "" {
			if v, ok := strings.CutPrefix(line, "Exec="); ok {
				exec = desktop.CleanExec(strings.TrimSpace(v))
				// TODO: wrap in something if Terminal=true
				fmt.Println(exec)
			}
		}

	}

	entry := desktop.EntryString(name, path)

	// read history (ignore missing file)
	hist, err := history.ReadHistory()
	if err != nil {
		return fmt.Errorf("read history: %w", err)
	}

	lineCount := strings.Count(hist, "\n")
	out := make([]string, 0, lineCount+2)
	out = append(out, entry)

	for x := range strings.SplitSeq(hist, "\n") {
		// only append if not a duplicate of the last added line
		if x == entry {
			continue
		}
		out = append(out, x)
	}

	historyContent := strings.Join(out, "\n")

	// rewrite history file
	if err := history.WriteHistory(historyContent); err != nil {
		return fmt.Errorf("write history: %w", err)
	}

	return nil
}
