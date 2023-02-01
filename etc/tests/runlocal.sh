#!/usr/bin/env bash

function runlocal () {
  local HERE=$(cd $(dirname $BASH_SOURCE) && pwd -P)

  . "$HERE/../lib.sh"

  logtest "run local"

  local srcDir=$(getvar "SRC_DIR")

  cd "$srcDir"

  go run main.go set "far.a" "new-val"
  go run main.go set "far.b" "incredible"

  local out=$(go run main.go get "far.a")
  echo "get: $out"

  local keys=$(go run main.go keys "far")
  echo "keys: $keys"
}
