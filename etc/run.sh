#!/usr/bin/env bash

set -e

function main () {
  local HERE=$(cd $(dirname $BASH_SOURCE) && pwd -P)
  local projRoot=$(cd "$HERE/.." && pwd -P)

  local srcDir="$projRoot/src"

  cd "$srcDir"

  go run main.go
}

( main $@ )
