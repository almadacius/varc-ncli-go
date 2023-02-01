#!/usr/bin/env bash

function test1 () {
  local HERE=$(cd $(dirname $BASH_SOURCE) && pwd -P)

  . "$HERE/../lib.sh"

  logtest "test 1"  

  buildApp

  # ===================================
  local outDir=$(getvar "DIST_DIR")

  cd "$outDir"

  ./varcgo set "far.bring" "new-val"
  ./varcgo set "far.it" "incredible"
  ./varcgo set "far.on" "incredible"
  ./varcgo set "ogre.shrek" "green"

  local out=$(./varcgo get "far.bring")
  echo "get: $out"

  local keys=$(./varcgo keys "far")
  echo "keys: $keys"

  ./varcgo unset "far.bring"
  out=$(./varcgo get "far.bring")
  echo "get after unset: $out"

  keys=$(./varcgo keys "far")
  echo "keys after unset: $keys"
}
