#!/usr/bin/env bash

# version printing feature
function runtest () {
  local HERE=$(cd $(dirname $BASH_SOURCE) && pwd -P)

  . "$HERE/../../etc/lib/lib.sh"

  logtest "test lock 2"

  # ===================================
  checkVerbose
  includeBinPath

  # ===================================
  varcgo testlock
}
