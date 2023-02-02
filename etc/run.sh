#!/usr/bin/env bash

set -e

# ===================================
function main() {
  local HERE=$(cd $(dirname $BASH_SOURCE) && pwd -P)

  . "$HERE/lib.sh"

  buildApp

  # ===================================
  . "$HERE/tests/runlocal.sh"
  . "$HERE/tests/test1.sh"
  . "$HERE/tests/testscope.sh"
  . "$HERE/tests/testtimer.sh"

  # runlocal
  # test1
  testscope
  # testtimer
}

# ===================================
# exit 1 is to avoid error code 2, which stops nodmons
# this happens on syntax error at compilation, for example
function onExit() {
  local exitCode="$?"
  if [ "$exitCode" = "2" ]; then
    exit 1
  fi
}

trap onExit EXIT

( main $@ )
