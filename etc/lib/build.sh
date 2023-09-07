#!/dev/null
# libray file, you should SOURCE it, not RUN it

# ===================================
function setTargetLinuxArm () {
  export GOOS="linux"
  export GOARCH="arm"
}

function setTargetLinux64 () {
  export GOOS="linux"
  export GOARCH="amd64"
}

function setTargetDarwin64 () {
  export GOOS="darwin"
  export GOARCH="amd64"
}

# ===================================
function includeBinPath () {
  local buildDir=$(getvar "buildDir")
  export PATH="$buildDir:$PATH"
}

# ===================================
function buildGenerateMeta () {
  local buildScriptDir=$(getvar "buildScriptDir")

  logheader "generate meta"
  time node "$buildScriptDir/genVersionSource.js"
}

function buildApp () {
  local buildScriptDir=$(getvar "buildScriptDir")

  buildGenerateMeta

  logheader "build app"
  time source "$buildScriptDir/build-local.sh"

  logbold
  logempty
}
