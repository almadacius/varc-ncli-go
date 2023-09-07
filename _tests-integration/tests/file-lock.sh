#!/usr/bin/env bash

function writeWave () {
  # simultaneous writes to the same file
  # should NOT corrupt
  varcgo set "far.bring" "new-val" &
  varcgo set "far.it" "incredible" &

  varcgo get "far.bring" > /dev/null &
  varcgo get "far.it" > /dev/null &
  varcgo get "far.on" > /dev/null &

  varcgo set "far.on" "nice" &
  varcgo set "far.again" "again" &
  varcgo set "far.andAgain" "other-value" &
  varcgo set "far.again" "again2" &
  varcgo set "far.key3" "again2" &
  varcgo set "far.key4" "again2" &
  varcgo set "far.key5" "again2" &
}

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

  writeWave
  writeWave
  writeWave
  writeWave
  writeWave
  writeWave
  writeWave
  writeWave

  # ===================================
  local out=$(varcgo get "far.bring")
  echo "get: $out"
}
