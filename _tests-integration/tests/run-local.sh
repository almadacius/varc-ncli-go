#!/usr/bin/env bash

# this is currently @BROKEN, `test-vars` is the main dev scenario
function runtest () {
  local HERE=$(cd $(dirname $BASH_SOURCE) && pwd -P)

  . "$HERE/../../etc/lib/lib.sh"

  logtest "run local"

  local srcDir=$(getvar "srcDir")

  cd "$srcDir"

  go run main.go set "far.a" "new-val"
  go run main.go set "far.b" "incredible"

  local out=$(go run main.go get "far.a")
  echo "get: $out"

  local keys=$(go run main.go keys "far")
  echo "keys: $keys"
}
