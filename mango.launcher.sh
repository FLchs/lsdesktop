#!/bin/sh

APP=$(./lsdesktop list | fzf --with-nth={1} --accept-nth={2} -d : --tiebreak=begin)


if [ -n "$APP" ]; then
  EX=$(./lsdesktop run "$APP")
  mmsg -d spawn,"$EX"
fi
