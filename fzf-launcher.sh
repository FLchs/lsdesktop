#!/bin/sh
#
# fzf-launcher.sh -- interactive desktop entry launcher using lsdesktop and fzf
#
# Lists desktop entries ordered by recency and launches the selected one.

export FZF_DEFAULT_OPTS=$FZF_DEFAULT_OPTS'
  --color=fg:#aeafb0,fg+:#d0d0d0,bg:#021727,bg+:#021727
  --color=hl:#5f87af,hl+:#577f63,info:#dbc074,marker:#87ff00
  --color=prompt:#d7005f,spinner:#af5fff,pointer:#021727,header:#87afaf
  --color=gutter:#021727,border:#cdcecf,label:#aeaeae,query:#d9d9d9
  --margin="0" --padding="0" --prompt="" --marker=" "
  --pointer=" " --separator="─" --scrollbar="│" --layout="reverse"
  --info="right"'

APP=$(lsdesktop | fzf --with-nth={1} --accept-nth=2 -d :)

if [ -n "$APP" ]; then
    exec desklaunch "$APP"
fi
