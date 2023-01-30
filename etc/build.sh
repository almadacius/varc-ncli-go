#!/usr/bin/env bash

set -e

function main () {
  local HERE=$(cd $(dirname $BASH_SOURCE) && pwd -P)
  local projRoot=$(cd "$HERE/.." && pwd -P)

  local srcDir="$projRoot/src"
  local outDir="$projRoot/dist"

  local output="$outDir/varcgo"

  mkdir -p "$outDir"

  cd "$srcDir"

  go build -o "$output" main.go
}

( main $@ )
