#!/usr/bin/env bash

set -e

function main () {
  local HERE=$(cd $(dirname $BASH_SOURCE) && pwd -P)

  . "$HERE/lib.sh"

  local srcDir=$(getvar "SRC_DIR")
  local distDir=$(getvar "DIST_DIR")

  export GOOS="linux" GOARCH="arm"

  distDir="$distDir/$GOOS-$GOARCH"
  local output="$distDir/varcgo"
  mkdir -p "$distDir"

  # ===================================
  cd "$srcDir"
  go build -o "$output" main.go
}

( main $@ )
