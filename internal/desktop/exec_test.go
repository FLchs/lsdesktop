package desktop

import (
	"testing"
)

// ai generated test usin dex as a model: https://github.com/jceb/dex
func TestParseExec(t *testing.T) {
	tests := []struct {
		input string
		want  []string
	}{
		// basic commands
		{"gvim", []string{"gvim"}},
		{"gvim test", []string{"gvim", "test"}},

		// double quotes
		{`"gvim" test`, []string{"gvim", "test"}},
		{`"gvim test"`, []string{"gvim test"}},
		{`"gvim test" test2 "test \" 3"`, []string{"gvim test", "test2", `test " 3`}},

		// single quotes
		{"vim '~/.vimrc test'", []string{"vim", "~/.vimrc test"}},

		// mixed quotes
		{"sh -c 'vim ~/.vimrc \" test\"'", []string{"sh", "-c", `vim ~/.vimrc " test"`}},
		{`sh -c "sh -c 'vim ~/.vimrc \" test\"'"`, []string{"sh", "-c", `sh -c 'vim ~/.vimrc " test"'`}},

		// field codes stripped
		{"vim %u", []string{"vim"}},
		{"vim ~/.vimrc %u", []string{"vim", "~/.vimrc"}},
		{"vim %u ~/.vimrc", []string{"vim", "~/.vimrc"}},
		{"vim /%u/.vimrc", []string{"vim", "//.vimrc"}},
		{"vim %u/.vimrc", []string{"vim", "/.vimrc"}},
		{"vim %U/.vimrc", []string{"vim", "/.vimrc"}},
		{"vim /%U/.vimrc", []string{"vim", "//.vimrc"}},
		{"vim %U .vimrc", []string{"vim", ".vimrc"}},

		// field codes preserved inside quotes
		{"vim '%u' ~/.vimrc", []string{"vim", "%u", "~/.vimrc"}},

		// escaped field code preserved
		{`vim \%u ~/.vimrc`, []string{"vim", "%u", "~/.vimrc"}},

		// unknown field codes preserved
		{"vim %x .vimrc", []string{"vim", "%x", ".vimrc"}},
		{"vim %x/.vimrc", []string{"vim", "%x/.vimrc"}},

		// literal %
		{"vim %%", []string{"vim", "%"}},
		{"vim %% ~/.vimrc", []string{"vim", "%", "~/.vimrc"}},

		// backslash pre-processing: \\\\ -> \\ so \\%u is not escaped
		{`vim \\%u ~/.vimrc`, []string{"vim", "%u", "~/.vimrc"}},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got, err := parseExec(tt.input)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if len(got) != len(tt.want) {
				t.Fatalf("len mismatch: got %v, want %v", got, tt.want)
			}
			for i := range got {
				if got[i] != tt.want[i] {
					t.Fatalf("arg %d: got %q, want %q", i, got[i], tt.want[i])
				}
			}
		})
	}
}
