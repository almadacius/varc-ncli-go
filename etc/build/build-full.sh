#!/usr/bin/env bash

# ===================================
function buildSingle () {
  local srcDir=$(getvar "srcDir")
  local distDir=$(getvar "outputDirArch")
  local output="$distDir/varcgo"

  mkdir -p "$distDir"

  # ===================================
  cd "$srcDir"
  time (
    go build -o "$output" main.go
  )
}

# ===================================
function main () {
  set -e

  local HERE=$(cd $(dirname $BASH_SOURCE) && pwd -P)

  . "$HERE/../lib/lib.sh"

  (
    logheader "target: linux-arm"
    setTargetLinuxArm
    buildSingle
  )
  (
    logheader "target: linux-amd64"
    setTargetLinux64
    buildSingle
  )
  (
    logheader "target: darwin-amd64"
    setTargetDarwin64
    buildSingle
  )

  logheader "success"
}

( main $@ )
