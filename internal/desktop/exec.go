package desktop

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"syscall"
)

// https://specifications.freedesktop.org/desktop-entry-spec/latest/exec-variables.html
var fieldCodes = map[byte]bool{
	'f': true, 'F': true,
	'u': true, 'U': true,
	'd': true, 'D': true,
	'n': true, 'N': true,
	'i': true,
	'c': true,
	'k': true,
	'v': true,
	'm': true,
}

func BuildCmd(execLine, path, tryExec string) (*exec.Cmd, error) {
	if tryExec != "" {
		if err := checkTryExec(tryExec); err != nil {
			return nil, err
		}
	}

	args, err := parseExec(execLine)
	if err != nil {
		return nil, fmt.Errorf("parse exec line: %w", err)
	}

	if len(args) == 0 {
		return nil, fmt.Errorf("empty command after parsing exec line")
	}

	cmd := exec.Command(args[0], args[1:]...)
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Setsid: true,
	}
	cmd.Stdin = nil
	cmd.Stdout = nil
	cmd.Stderr = nil
	if path != "" {
		cmd.Dir = path
	}

	return cmd, nil
}

// probably overkill in most case but should be close enough to the freedesktop specs
func parseExec(execLine string) ([]string, error) {
	execLine = strings.ReplaceAll(execLine, "%%", "\x00")
	execLine = strings.ReplaceAll(execLine, `\\`, `\`)

	var args []string
	var arg strings.Builder
	inEsc := false
	inQuote := byte(0)

	for i := 0; i < len(execLine); i++ {
		c := execLine[i]

		if inEsc {
			arg.WriteByte(c)
			inEsc = false
			continue
		}

		if c == '\\' {
			inEsc = true
			continue
		}

		if inQuote != 0 {
			if c == inQuote {
				inQuote = 0
				args = append(args, arg.String())
				arg.Reset()
			} else {
				arg.WriteByte(c)
			}
			continue
		}

		if c == '"' || c == '\'' {
			if arg.Len() > 0 {
				args = append(args, arg.String())
				arg.Reset()
			}
			inQuote = c
			continue
		}

		if c == '%' && i+1 < len(execLine) && fieldCodes[execLine[i+1]] {
			i++
			continue
		}

		if c == ' ' || c == '\t' || c == '\n' || c == '\r' {
			if arg.Len() > 0 {
				args = append(args, arg.String())
				arg.Reset()
			}
			continue
		}

		arg.WriteByte(c)
	}

	if inEsc || inQuote != 0 {
		return nil, fmt.Errorf("unbalanced quotes or escapes in exec line")
	}

	if arg.Len() > 0 {
		args = append(args, arg.String())
	}

	for i := range args {
		args[i] = strings.ReplaceAll(args[i], "\x00", "%")
	}

	return args, nil
}

const xOK = 1 // syscall.X_OK is not exported on all platforms; POSIX defines it as 1.

func checkTryExec(name string) error {
	// Use syscall.Access with xOK instead of checking mode bits.
	// Mode bits only tell us that *someone* can execute the file,
	// but Access checks whether the *current user* actually can.
	// This matters on multi-user systems or when permissions are tight.
	if strings.Contains(name, string(os.PathSeparator)) {
		if err := syscall.Access(name, xOK); err != nil {
			return fmt.Errorf("tryexec not found or not executable: %s", name)
		}
		return nil
	}

	pathEnv := os.Getenv("PATH")
	for dir := range strings.SplitSeq(pathEnv, string(os.PathListSeparator)) {
		if dir == "" {
			continue
		}
		full := dir + string(os.PathSeparator) + name
		if err := syscall.Access(full, xOK); err == nil {
			return nil
		}
	}

	return fmt.Errorf("tryexec not found in PATH: %s", name)
}
