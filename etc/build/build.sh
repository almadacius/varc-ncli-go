#!/usr/bin/env bash

set -e

function main () {
  local HERE=$(cd $(dirname $BASH_SOURCE) && pwd -P)

  . "$HERE/../lib.sh"

  local srcDir=$(getvar "SRC_DIR")
  local buildDir=$(getvar "BUILD_DIR")

  local output="$buildDir/varcgo"

  mkdir -p "$buildDir"

  cd "$srcDir"

  go build -o "$output" main.go
}

( main $@ )
