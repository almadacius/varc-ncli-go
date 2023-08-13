#!/usr/bin/env bash

# copy build version to deploy dir for actual usage
function main () {
  set -e

  local HERE=$(cd $(dirname $BASH_SOURCE) && pwd -P)

  . "$HERE/../lib.sh"

  local buildDir=$(getvar "BUILD_DIR")
  local distDir=$(getvar "DIST_DIR")

  mkdir -p "$distDir"

  cp "$buildDir/varcgo" "$distDir/varcgo"
}

( main $@ )
