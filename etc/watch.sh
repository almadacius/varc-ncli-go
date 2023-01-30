#!/usr/bin/env bash

set -e

function main () {
  local HERE=$(cd $(dirname $BASH_SOURCE) && pwd -P)
  local projRoot=$(cd "$HERE/.." && pwd -P)

  local srcDir="$projRoot/src"

  cd "$srcDir"

  nodemon --config "$HERE/nodemon.json" \
    main.go list -text "nice nice"
}

( main $@ )
