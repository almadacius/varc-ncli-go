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
  local scriptDir=$(getvar "scriptDir")

  logheader "build app"
  time source "$scriptDir/build.sh"
  logbold
  logempty
}
