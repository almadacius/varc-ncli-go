#!/usr/bin/env bash

# version printing feature
function runtest () {
  local HERE=$(cd $(dirname $BASH_SOURCE) && pwd -P)

  . "$HERE/../../etc/lib/lib.sh"

  logtest "print version"

  # ===================================
  checkVerbose
  includeBinPath

  # ===================================
  varcgo version
}
