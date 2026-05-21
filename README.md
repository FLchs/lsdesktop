LSDESKTOP(1)                 General Commands Manual                LSDESKTOP(1)

NAME
       lsdesktop -- list and launch XDG desktop entries

SYNOPSIS
       lsdesktop [list]
       lsdesktop launch <path>
       lsdesktop print <path>

DESCRIPTION
       lsdesktop reads .desktop files from the standard XDG data directories
       and prints visible entries as name:path pairs.  Entries are ordered by
       most recently launched first.

       launch parses the Exec line, strips field codes per the Desktop Entry
       specification, and starts the command in the background.  print outputs
       the cleaned Exec command to stdout without running it.  Both update the
       history file.

COMMANDS
       list          Print visible desktop entries as name:path.  Most
                     recently used entries appear first.  This is the default
                     when no command is given.

       launch path   Parse the .desktop file at path and execute it.  Updates
                     the history file.

       print path    Print the exec command for the .desktop file at path.
                     Updates the history file.

FILES
       ~/.cache/lsdesktop/history
                     Launch history.  Updated on each launch, used to order
                     list output.

ENVIRONMENT
       XDG_DATA_HOME   Defaults to ~/.local/share.
       XDG_DATA_DIRS   Defaults to /usr/local/share:/usr/share.

EXIT STATUS
       0       Successful execution.
       1       Usage error, missing argument, or command failure.

EXAMPLES
       List entries ordered by recency:

              lsdesktop list

       Launch a specific application:

              lsdesktop launch /usr/share/applications/firefox.desktop

       Interactive launcher with fzf:

              APP=$(lsdesktop list | fzf --with-nth=1 --accept-nth=2 -d : \
                      --tiebreak=begin)
              [ -n "$APP" ] && lsdesktop launch "$APP"

LICENSE
       GPL v3.  See LICENSE.
