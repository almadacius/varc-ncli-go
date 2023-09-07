#!/usr/bin/env bash

function main () {
  set -e

  local HERE=$(cd $(dirname $BASH_SOURCE) && pwd -P)

  . "$HERE/../../etc/lib/lib.sh"

  local projRoot=$(getvar "projRoot")

  cd "$projRoot"

  export APP_MODE_='dev'
  npx almonitor "$HERE/run.sh"
}

( main $@ )
