#!/usr/bin/env bash

function testtimer () {
  local HERE=$(cd $(dirname $BASH_SOURCE) && pwd -P)

  . "$HERE/../lib.sh"

  logtest "test scope"

  checkverbose

  # ===================================
  local buildDir=$(getvar "BUILD_DIR")

  cd "$buildDir"

  time (
    local timerkey
    local out

    logheader "create timer"
    timerkey=$(./varcgo timercreate)
    logbold
    echo "timerkey: $timerkey"
    logbold

    out=$(./varcgo timerstep "$timerkey")
    logbold
    echo "step: $out"
    logbold

    sleep 2

    out=$(./varcgo timerstep "$timerkey")
    logbold
    echo "step: $out"
    logbold

    out=$(./varcgo timerend "$timerkey")
    logbold
    echo "end: $out"
    logbold

    ./varcgo timerprune
  )
}
