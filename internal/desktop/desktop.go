package desktop

import (
	"regexp"
	"strings"
)

func ParseEntry(data string) (DesktopEntry, bool) {
	var name string
	visible := true
	var entry DesktopEntry

	for line := range strings.SplitSeq(data, "\n") {
		line = strings.TrimSpace(line)

		if v, ok := strings.CutPrefix(line, "Name="); ok && name == "" {
			name = v
		}
		if v, ok := strings.CutPrefix(line, "NoDisplay="); ok {
			if strings.TrimSpace(v) == "true" {
				visible = false
			}
		}
		if v, ok := strings.CutPrefix(line, "Hidden="); ok {
			if strings.TrimSpace(v) == "true" {
				visible = false
			}
		}
		if v, ok := strings.CutPrefix(line, "Terminal="); ok {
			if strings.TrimSpace(v) == "true" {
				visible = false
			}
		}
		if v, ok := strings.CutPrefix(line, "OnlyShowIn="); ok {
			for s := range strings.SplitSeq(v, ";") {
				if s != "" {
					entry.OnlyShowIn = append(entry.OnlyShowIn, s)
				}
			}
		}
		if v, ok := strings.CutPrefix(line, "NotShowIn="); ok {
			for s := range strings.SplitSeq(v, ";") {
				if s != "" {
					entry.NotShowIn = append(entry.NotShowIn, s)
				}
			}
		}
	}

	return DesktopEntry{Name: name}, visible && name != ""
}

func CleanExec(execLine string) string {
	// regex of exec arguments : https://specifications.freedesktop.org/desktop-entry/latest/exec-variables.html
	reFieldCodes := regexp.MustCompile(`['"]?%[fFuUdDnNiickvm]['"]?`)
	cleaned := reFieldCodes.ReplaceAllString(execLine, "")

	cleaned = strings.ReplaceAll(cleaned, "%%", "%")

	reSpaces := regexp.MustCompile(`\s+`)
	cleaned = reSpaces.ReplaceAllString(cleaned, " ")

	return strings.TrimSpace(cleaned)
}

type DesktopEntry struct {
	Name       string
	Path       string
	OnlyShowIn []string
	NotShowIn  []string
}

func EntryString(name, path string) string {
	return name + ":" + path
}
