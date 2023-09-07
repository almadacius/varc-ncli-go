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

# ===================================
function includeBinPath () {
  local buildDir=$(getvar "buildDir")
  export PATH="$buildDir:$PATH"
}

# ===================================
function buildApp () {
  local buildScriptDir=$(getvar "buildScriptDir")

  logheader "generate meta"
  time node "$buildScriptDir/genVersionSource.js"

  logheader "build app"
  time source "$buildScriptDir/build.sh"

  logbold
  logempty
}
