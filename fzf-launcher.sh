#!/bin/sh
#
# fzf-launcher.sh -- interactive desktop entry launcher using lsdesktop and fzf
#
# Lists desktop entries ordered by recency and launches the selected one.

APP=$(lsdesktop | fzf --with-nth={1} --accept-nth=2 -d : --tiebreak=begin)

if [ -n "$APP" ]; then
    desklaunch "$APP"
fi
