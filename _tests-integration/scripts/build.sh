#!/usr/bin/env bash

function main () {
  set -e

  local HERE=$(cd $(dirname $BASH_SOURCE) && pwd -P)

  . "$HERE/../../etc/lib/lib.sh"

  local srcDir=$(getvar "srcDir")
  local buildDir=$(getvar "buildDir")

  local output="$buildDir/varcgo"

  mkdir -p "$buildDir"

  cd "$srcDir"

  go build -o "$output" main.go
}

( main $@ )
