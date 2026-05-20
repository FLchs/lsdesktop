package commands

import (
	"flc/lsdesktop/internal/desktop"
	"flc/lsdesktop/internal/history"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

func List() {
	rootDir := "/usr/share/applications/"

	hist, _ := history.ReadHistory()

	// show history before listing remaining applications
	if hist != "" {
		fmt.Println(hist)
	}

	// build a set of applications from the history file
	seen := make(map[string]struct{})
	for entry := range strings.SplitSeq(hist, "\n") {
		if n := strings.TrimSpace(entry); n != "" {
			namePart, _, _ := strings.Cut(n, ":")
			seen[namePart] = struct{}{}
		}
	}

	dir := os.DirFS(rootDir)
	var apps []desktop.DesktopEntry

	// walk the applications directory
	fs.WalkDir(dir, ".", func(filePath string, d fs.DirEntry, err error) error {

		// skip directories and non-desktop files
		if err != nil || d.IsDir() || !strings.HasSuffix(filePath, ".desktop") {
			return nil
		}

		data, err := fs.ReadFile(dir, filePath)
		if err != nil {
			return nil
		}

		entry, ok := desktop.ParseEntry(string(data))
		if !ok {
			return nil
		}

		// skip apps that appear in the history
		if _, ok := seen[entry.Name]; ok {
			return nil
		}

		entry.Path = filepath.Join(rootDir, filePath)
		apps = append(apps, entry)

		return nil
	})

	// output the result minus the ones from history
	for _, app := range apps {
		fmt.Println(desktop.EntryString(app.Name, app.Path))
	}
}
