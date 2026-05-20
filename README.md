# lsdesktop

List XDG desktop entries with recency-based ordering.

## Building

    go build

## Commands

    list          Print visible desktop entries as name:path.
                  Most recently used entries appear first.

    run path      Print the exec command for the .desktop file at path.
                  Updates the history file.

## Files

    ~/.cache/lsdesktop/history
                  Run history.  Updated on each run, used to order list.

## Environment

    XDG_DATA_HOME         Defaults to ~/.local/share
    XDG_DATA_DIRS         Defaults to /usr/local/share:/usr/share

## Examples

The mango.launcher.sh script provides a ready-made integration for mango WM.

## License

GPL v3.  See LICENSE.
