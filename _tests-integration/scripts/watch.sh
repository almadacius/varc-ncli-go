#!/usr/bin/env bash

function main () {
  set -e

  local HERE=$(cd $(dirname $BASH_SOURCE) && pwd -P)

  . "$HERE/../../etc/lib/lib.sh"

  local projRoot=$(getvar "projRoot")

  cd "$projRoot"

  nodemon \
    --config "$HERE/nodemon.json" \
    "$HERE/run.sh"
}

( main $@ )
