#!/usr/bin/env bash

# copy build version to deploy dir for actual usage
function main () {
  set -e

  local HERE=$(cd $(dirname $BASH_SOURCE) && pwd -P)

  . "$HERE/../lib/lib.sh"

  local buildDir=$(getvar "buildDir")
  local distDir=$(getvar "distDir")

  mkdir -p "$distDir"

  cp "$buildDir/varcgo" "$distDir/varcgo"
}

( main $@ )
