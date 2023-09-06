#!/dev/null
# libray file, you should SOURCE it, not RUN it

# ===================================
function getvar () {
  local varname="$1"
  if [ -z "$varname" ]; then
    echo "[Error]: variable name not provided"
    return 1
  fi

  # ===================================
  local HERE=$(cd $(dirname $BASH_SOURCE) && pwd -P)
  local projRoot=$(cd "$HERE/../.." && pwd -P)

  local srcDir="$projRoot/src"
  local buildDir="$projRoot/build"
  local distDir="$projRoot/dist"
  local etcDir="$projRoot/etc"
  local libDir="$projRoot/etc/lib"

  local scriptDir="$projRoot/_tests-integration/scripts"

  local outputDirArch="$distDir/$GOOS-$GOARCH"

  # ===================================
  local verbose="false"
  # verbose="true"

  # ===================================
  if [ "${!varname-unset}" = "unset" ]; then
    echo "[Error]: variable '$varname' is not set"
    return 1
  fi

  # ===================================
  # bash-specific syntax
  local value=${!varname}

  echo "$value"
}

# ===================================
function checkVerbose () {
  local verbose=$(getvar "verbose")
  if [ "$verbose" = "true" ]; then
    set -x
  fi
}

# ===================================
# - exit code 2 stops nodemon
# - happens, for example, on syntax error at compilation
# - force exit code 1, so that nodemon can continue running
# - @usage: trap forceExitCode1 EXIT
function forceExitCode1 () {
  local exitCode="$?"
  if [ "$exitCode" = "2" ]; then
    exit 1
  fi
}
