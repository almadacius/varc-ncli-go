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
    local timerkey
    local out

    logheader "create timer"
    timerkey=$(varcgo timercreate)
    logbold
    echo "timerkey: $timerkey"
    logbold

    out=$(varcgo timerstep "$timerkey")
    logbold
    echo "step: $out"
    logbold

    sleep 2

    out=$(varcgo timerstep "$timerkey")
    logbold
    echo "step: $out"
    logbold

    out=$(varcgo timerend "$timerkey")
    logbold
    echo "end: $out"
    logbold

    varcgo timerprune
  )
}
