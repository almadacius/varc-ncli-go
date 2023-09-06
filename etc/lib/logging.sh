#!/dev/null
# libray file, you should SOURCE it, not RUN it

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

# ===================================
function logsep () {
  logline
  logempty
}

function logheader () {
  local text="$1"
  logbold
  echo "$text"
  logbold
}

# ===================================
function logtest () {
  local text="$1"
  logheader "[test]: $text"
}
