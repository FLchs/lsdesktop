package commands

import (
	"flc/lsdesktop/internal/desktop"
	"flc/lsdesktop/internal/history"
	"fmt"
	"os"
	"strings"
)

// Launch parses the desktop entry at path and starts it in the background.
func Launch(path string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("read desktop entry: %w", err)
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
		return fmt.Errorf("no Exec line found in %s", path)
	}

	cmd, err := desktop.BuildCmd(execLine, workDir, tryExec)
	if err != nil {
		return fmt.Errorf("build command: %w", err)
	}

	if err := cmd.Start(); err != nil {
		return fmt.Errorf("start command: %w", err)
	}

	// Update history using the same logic as Print.
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
