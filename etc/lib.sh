#!/usr/bin/env bash

# ===================================
function logbold () {
  echo "================================================="
}

function logline () {
  echo "-------------------------------------------------"
}

function logempty () {
  echo ""
}

function logsep () {
  logline
  logempty
}

function logheader () {
  local text="$1"
  echo "================================================="
  echo "$text"
  echo "================================================="
}

function logtest () {
  local text="$1"
  logheader "[test]: $text"
}

# ===================================
function getvar () {
  local varname="$1"
  if [ -z "$varname" ]; then
    echo "[Error]: variable name not provided"
    return 1
  fi

  # ===================================
  local HERE=$(cd $(dirname $BASH_SOURCE) && pwd -P)
  local PROJECT_ROOT=$(cd "$HERE/.." && pwd -P)

  local SRC_DIR="$PROJECT_ROOT/src"
  local BUILD_DIR="$PROJECT_ROOT/build"
  local DIST_DIR="$PROJECT_ROOT/dist"
  local ETC_DIR="$PROJECT_ROOT/etc"

  # ===================================
  if [ "${!varname-unset}" = "unset" ]; then
    echo "[Error]: variable '$varname' is not set"
    return 1
  fi

  # ===================================
  # Bash-specific
  local value=${!varname}

  echo "$value"
}

# ===================================
function checkverbose () {
  local verbose="true"
  if [ "$verbose" = "true" ]; then
    set -x
  fi
}

function buildApp () {
  local HERE=$(cd $(dirname $BASH_SOURCE) && pwd -P)

  local etcDir=$(getvar "ETC_DIR")

  logheader "build app"
  time source "$etcDir/build.sh"
  logbold
  logempty
}

# ===================================
echo "load lib"
