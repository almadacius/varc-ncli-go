#!/usr/bin/env bash

function testscope () {
  local HERE=$(cd $(dirname $BASH_SOURCE) && pwd -P)

  . "$HERE/../lib.sh"

  logtest "test scope"

  buildApp

  # ===================================
  local outDir=$(getvar "DIST_DIR")

  cd "$outDir"

  time (
    local out

    logheader "set far.bring"
    ./varcgo set "far.bring" "new-val"

    out=$(./varcgo scopes)
    logbold
    echo "scopes"
    logline
    echo "$out"
    logbold

    logheader "delete scope: far"
    out=$(./varcgo del "far")

    out=$(./varcgo scopes)
    logbold
    echo "scopes"
    logline
    echo "$out"
    logbold
  )
}