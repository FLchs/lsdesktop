package history

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func historyFilePath() (string, error) {
	cacheDir, err := os.UserCacheDir()
	if err != nil {
		return "", fmt.Errorf("get user cache dir: %w", err)
	}
	dir := filepath.Join(cacheDir, "lsdesktop")
	if err := os.MkdirAll(dir, 0755); err != nil {
		return "", fmt.Errorf("create cache dir: %w", err)
	}
	return filepath.Join(dir, "history"), nil
}

func ReadHistory() (string, error) {
	path, err := historyFilePath()
	if err != nil {
		return "", fmt.Errorf("get history file path: %w", err)
	}
	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return "", nil
		}
		return "", fmt.Errorf("read history file: %w", err)
	}
	content := strings.TrimSpace(string(data))
	if content == "" {
		return "", nil
	}
	return content, nil
}

func WriteHistory(content string) error {
	path, err := historyFilePath()
	if err != nil {
		return fmt.Errorf("get history file path: %w", err)
	}
	if err := os.WriteFile(path, []byte(content), 0644); err != nil {
		return fmt.Errorf("write history file: %w", err)
	}
	return nil
}
