#!/usr/bin/env bash

function runtest () {
  local HERE=$(cd $(dirname $BASH_SOURCE) && pwd -P)

  . "$HERE/../../etc/lib/lib.sh"

  logtest "file lock"

  # ===================================
  checkVerbose
  includeBinPath

  # ===================================
  # clear scope
  varcgo del "far"

  # simultaneous writes to the same file
  # should NOT corrupt
  varcgo set "far.bring" "new-val" &
  varcgo set "far.it" "incredible" &
  varcgo set "far.on" "nice" &

  # ===================================
  local out=$(varcgo get "far.bring")
  echo "get: $out"
}
