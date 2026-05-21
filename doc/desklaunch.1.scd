DESKLAUNCH(1)

# NAME

desklaunch - launch a .desktop file and update history file

# SYNOPSIS

*desklaunch* [*-p*] <path>

# DESCRIPTION

Reads the .desktop file at _path_, strips field codes from the Exec
line, and runs the command in the background.

With the *-p* flag the command is printed to stdout instead of
executed. This is useful for inspecting the command or piping it to
another tool.

Both modes update the shared history file, which affects ordering in
*lsdesktop*(1).

# OPTIONS

*-p*
	Print the cleaned exec command. Do not run it.

# FILES

_$HOME/.cache/lsdesktop/history_
	Shared launch history. Written by *desklaunch*, read by
	*lsdesktop*(1).

# EXIT STATUS

*0*
	Success.

*1*
	Usage error, missing argument, or command failure.

# SEE ALSO

*lsdesktop*(1), *desktop-file-validate*(1)

# AUTHORS

Written by Francois Lachese.
