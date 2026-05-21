# lsdesktop

Two small utilities for working with XDG desktop entries.

`lsdesktop` lists visible .desktop files as name:path, with the most
recently launched shown first.

`desklaunch` parses a single .desktop file and runs it, or prints the
cleaned exec command with `-p`. Both update the shared history so
`lsdesktop` knows what you launched last.

This is mostly *tradcoded*, but some parts were AI-assisted (mostly tests and regexes :))  and carefully
reviewed. It is a learning project. (Ironically, this paragraph was
partly LLM-generated.)

## Installation

With Go:

    go install github.com/FLchs/lsdesktop/cmd/...@latest

From source:

    git clone https://github.com/FLchs/lsdesktop.git
    cd lsdesktop
    make
    sudo make install

`make install` respects `PREFIX` (defaults to `/usr/local`) and puts
man pages where they belong.

## Usage

List applications:

    lsdesktop

Launch one:

    desklaunch /usr/share/applications/firefox.desktop

Print the command without running it:

    desklaunch -p /usr/share/applications/firefox.desktop

Use with fzf for an interactive picker:

    desklaunch $(lsdesktop | fzf --with-nth={1} --accept-nth=2 -d : --tiebreak=begin)

## License

GPL v3. See LICENSE.
