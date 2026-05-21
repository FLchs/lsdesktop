# lsdesktop

Two small utilities for listing and launching XDG desktop entries that are meant to be used together.

`lsdesktop` prints visible .desktop files as name:path, ordered by most
recently launched first.

`desklaunch` parses a single .desktop file and launches it,
or prints its cleaned Exec command with `-p`.  
Both update the shared history file (to display the most recent items first on lsdesktop).

## Building

    go build ./cmd/...

## Quick start

list .desktop applications:
`lsdesktop`

Launch firefox:
`desklaunch /usr/share/applications/firefox.desktop`

Print the cleaned exec command to start firefox with dex, gtk-launch or whatever suit you while still keeping the history updated:
`desklaunch -p /usr/share/applications/firefox.desktop`

Interactive picker with fzf:
`desklaunch $(lsdesktop | fzf --with-nth={1} --accept-nth=2 -d : --tiebreak=begin)`

## Documentation

Manual pages are in `doc/`

## License

GPL v3.  See LICENSE.
