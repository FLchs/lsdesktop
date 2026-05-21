LSDESKTOP(1)

# NAME

lsdesktop - list XDG desktop entries, recent first

# SYNOPSIS

*lsdesktop*

# DESCRIPTION

Lists .desktop files from the standard XDG application directories,
printing visible entries as _name:path_. Output is sorted by most
recently launched first, using the history file maintained by
*desklaunch*(1). Before anything has been launched, order follows the
filesystem.

Hidden, NoDisplay, and Terminal entries are omitted.

# FILES

_$HOME/.cache/lsdesktop/history_
	Launch history. Updated by *desklaunch*(1), read by *lsdesktop*
	to order output.

# ENVIRONMENT

_XDG_DATA_HOME_
	Defaults to _$HOME/.local/share_.

_XDG_DATA_DIRS_
	Defaults to _/usr/local/share:/usr/share_.

# EXIT STATUS

*0*
	Success.

*1*
	An error occurred.

# SEE ALSO

*desklaunch*(1), *fzf*(1)

# AUTHORS

Written by Francois Lachese.
