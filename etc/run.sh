#!/usr/bin/env bash

set -e

function main () {
  local HERE=$(cd $(dirname $BASH_SOURCE) && pwd -P)
  local projRoot=$(cd "$HERE/.." && pwd -P)

  local srcDir="$projRoot/src"

  cd "$srcDir"

  # exit 1 is to avoid error code 2, which stops nodmons
  # this happens on syntax error at compilation, for example
  go run main.go set "far.a" "new-val" || exit 1
}

( main $@ )
