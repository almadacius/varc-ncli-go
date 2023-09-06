#!/usr/bin/env bash

function main () {
  set -e

  local HERE=$(cd $(dirname $BASH_SOURCE) && pwd -P)

  . "$HERE/../lib/lib.sh"

  setTargetLinux64

  local srcDir=$(getvar "srcDir")
  local distDir=$(getvar "outputDirArch")
  
  local output="$distDir/varcgo"
  mkdir -p "$distDir"

  # ===================================
  cd "$srcDir"
  go build -o "$output" main.go
}

( main $@ )
