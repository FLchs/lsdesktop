DESKLAUNCH(1)

# NAME

desklaunch - launch or print the parsed exec line from a single XDG desktop entry and keep an history of started commands

# SYNOPSIS

*desklaunch* [*-p*] <path>

# DESCRIPTION

*desklaunch* parses the .desktop file at _path_, extracts the _Exec=_
line, strips field codes per the Desktop Entry specification, and starts
the command in the background.

With the *-p* flag the command is printed to stdout instead of executed.

Both modes update the launch history file so that *lsdesktop*(1) ordering
stays current.

# OPTIONS

- *-p*
	Print the cleaned exec command to stdout without running it.

# FILES

- _$HOME/.cache/lsdesktop/history_
	Launch history.  Updated by *desklaunch*, read by *lsdesktop*(1).

# EXIT STATUS

- 0
	Successful execution.

- 1
	Usage error, missing argument, or command failure.

# SEE ALSO

*lsdesktop*(1), *desktop-file-validate*(1)

# AUTHORS

Maintained by the lsdesktop authors.
