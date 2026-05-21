package history

import (
	"os"
	"path/filepath"
	"testing"
)

// ai generated test

func TestRoundTrip(t *testing.T) {
	tmp := t.TempDir()
	t.Setenv("XDG_CACHE_HOME", tmp)

	want := "Firefox:/usr/share/applications/firefox.desktop"
	if err := WriteHistory(want); err != nil {
		t.Fatalf("write history: %v", err)
	}

	got, err := ReadHistory()
	if err != nil {
		t.Fatalf("read history: %v", err)
	}
	if got != want {
		t.Fatalf("got %q, want %q", got, want)
	}
}

func TestReadHistoryEmptyFile(t *testing.T) {
	tmp := t.TempDir()
	t.Setenv("XDG_CACHE_HOME", tmp)

	// create an empty history file
	path := filepath.Join(tmp, "lsdesktop", "history")
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		t.Fatalf("mkdir: %v", err)
	}
	if err := os.WriteFile(path, []byte(""), 0644); err != nil {
		t.Fatalf("write empty file: %v", err)
	}

	got, err := ReadHistory()
	if err != nil {
		t.Fatalf("read history: %v", err)
	}
	if got != "" {
		t.Fatalf("got %q, want empty string", got)
	}
}

func TestReadHistoryMissingFile(t *testing.T) {
	tmp := t.TempDir()
	t.Setenv("XDG_CACHE_HOME", tmp)

	got, err := ReadHistory()
	if err != nil {
		t.Fatalf("read history: %v", err)
	}
	if got != "" {
		t.Fatalf("got %q, want empty string", got)
	}
}

func TestWriteHistoryCreatesDir(t *testing.T) {
	tmp := t.TempDir()
	t.Setenv("XDG_CACHE_HOME", tmp)

	// the lsdesktop subdir should not exist yet
	if err := WriteHistory("test"); err != nil {
		t.Fatalf("write history: %v", err)
	}

	path := filepath.Join(tmp, "lsdesktop", "history")
	if _, err := os.Stat(path); err != nil {
		t.Fatalf("history file not created: %v", err)
	}
}
