#!/usr/bin/env bash

set -e

function main () {
  local HERE=$(cd $(dirname $BASH_SOURCE) && pwd -P)
  local projRoot=$(cd "$HERE/.." && pwd -P)

  local outDir="$projRoot/dist"

  cd "$outDir"

  ./varcgo get "far"
}

( main $@ )
