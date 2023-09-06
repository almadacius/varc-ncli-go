#!/usr/bin/env bash

# ===================================
function getTest () {
  # echo "run-local.sh"
  # echo "test-vars.sh"
  # echo "test-scope.sh"
  # echo "test-timer.sh"
  echo "file-lock.sh"
}

# ===================================
function main() {
  set -e

  local HERE=$(cd $(dirname $BASH_SOURCE) && pwd -P)
  local testDir="$HERE/../tests"

  . "$HERE/../../etc/lib/lib.sh"

  trap forceExitCode1 EXIT

  buildApp

  # ===================================
  local testName=$(getTest)
  local testFile="$testDir/$testName"

  if [ ! -f "$testFile" ]; then
    echo "[Error]: test file not found: $testName"
    return 1
  fi

  . "$testFile"

  runtest
}

# ===================================
( main $@ )
