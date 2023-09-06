#!/usr/bin/env bash

function runtest () {
  local HERE=$(cd $(dirname $BASH_SOURCE) && pwd -P)

  . "$HERE/../../etc/lib/lib.sh"

  logtest "test scope"

  # ===================================
  checkVerbose
  includeBinPath

  # ===================================
  time (
    local out

    logheader "set far.bring"
    varcgo set "far.bring" "new-val"

    out=$(varcgo scopes)
    logbold
    echo "scopes"
    logline
    echo "$out"
    logbold

    logheader "delete scope: far"
    out=$(varcgo del "far")

    out=$(varcgo scopes)
    logbold
    echo "scopes"
    logline
    echo "$out"
    logbold
  )
}
