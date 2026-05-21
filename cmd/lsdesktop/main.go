package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"flc/lsdesktop/internal/desktop"
	"flc/lsdesktop/internal/history"
)

func appDirs() []string {
	dirs := make([]string, 0, 4)

	dataHome := os.Getenv("XDG_DATA_HOME")
	if dataHome == "" {
		if home, err := os.UserHomeDir(); err == nil {
			dataHome = filepath.Join(home, ".local", "share")
		}
	}
	if dataHome != "" {
		dirs = append(dirs, filepath.Join(dataHome, "applications"))
	}

	dataDirs := os.Getenv("XDG_DATA_DIRS")
	if dataDirs == "" {
		dataDirs = "/usr/local/share:/usr/share"
	}
	for d := range strings.SplitSeq(dataDirs, ":") {
		if d != "" {
			dirs = append(dirs, filepath.Join(d, "applications"))
		}
	}

	return dirs
}

func main() {
	hist, _ := history.ReadHistory()

	if hist != "" {
		fmt.Println(hist)
	}

	seen := make(map[string]struct{})
	for entry := range strings.SplitSeq(hist, "\n") {
		if n := strings.TrimSpace(entry); n != "" {
			namePart, _, _ := strings.Cut(n, ":")
			seen[namePart] = struct{}{}
		}
	}

	var apps []desktop.DesktopEntry

	for _, rootDir := range appDirs() {
		dir := os.DirFS(rootDir)

		fs.WalkDir(dir, ".", func(filePath string, d fs.DirEntry, err error) error {
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

			if _, ok := seen[entry.Name]; ok {
				return nil
			}

			entry.Path = filepath.Join(rootDir, filePath)
			apps = append(apps, entry)
			seen[entry.Name] = struct{}{}

			return nil
		})
	}

	for _, app := range apps {
		fmt.Println(desktop.EntryString(app.Name, app.Path))
	}
}
