#!/dev/null
# libray file, you should SOURCE it, not RUN it

function loadlib () {
  local HERE=$(cd $(dirname $BASH_SOURCE) && pwd -P)

  . "$HERE/vars.sh"
  . "$HERE/logging.sh"
  . "$HERE/build.sh"
}

loadlib
