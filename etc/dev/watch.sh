#!/usr/bin/env bash

set -e

function main () {
  local HERE=$(cd $(dirname $BASH_SOURCE) && pwd -P)
  local projRoot=$(cd "$HERE/.." && pwd -P)

  cd "$projRoot"

  nodemon \
    --config "$HERE/nodemon.json" \
    etc/run.sh
}

( main $@ )
