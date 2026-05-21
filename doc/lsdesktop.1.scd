LSDESKTOP(1)

# NAME

lsdesktop - list XDG desktop entries with recency-based ordering

# SYNOPSIS

*lsdesktop*

# DESCRIPTION

*lsdesktop* reads .desktop files from the standard XDG data directories
and prints visible entries as _name:path_ pairs.  Entries are ordered by
most recently launched first, using the history file maintained by
*desklaunch*(1).

Hidden, NoDisplay, and Terminal entries are omitted.

# FILES

- _$HOME/.cache/lsdesktop/history_
	Launch history.  Updated by *desklaunch*(1), read by *lsdesktop* to
	order output.

# ENVIRONMENT

- *XDG_DATA_HOME*
	Defaults to _$HOME/.local/share_.

- *XDG_DATA_DIRS*
	Defaults to _/usr/local/share:/usr/share_.

# EXIT STATUS

- 0
	Successful execution.

- 1
	An error occurred.

# SEE ALSO

*desklaunch*(1), *fzf*(1)

# AUTHORS

Maintained by the lsdesktop authors.
