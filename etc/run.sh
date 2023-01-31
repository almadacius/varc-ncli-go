#!/usr/bin/env bash

set -e

function runlocal () {
  local HERE=$(cd $(dirname $BASH_SOURCE) && pwd -P)
  local projRoot=$(cd "$HERE/.." && pwd -P)

  local srcDir="$projRoot/src"

  cd "$srcDir"

  go run main.go set "far.a" "new-val"
  go run main.go set "far.b" "incredible"

  local out=$(go run main.go get "far.a")
  echo "get: $out"

  local keys=$(go run main.go keys "far")
  echo "keys: $keys"
}

function test1 () {
  local HERE=$(cd $(dirname $BASH_SOURCE) && pwd -P)
  local projRoot=$(cd "$HERE/.." && pwd -P)

  . "$HERE/build.sh"

  local outDir="$projRoot/dist"

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

# ===================================
function main() {
  # runlocal
  test1
}

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
