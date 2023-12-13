package lock

import (
  "fmt"
  "os"
  gopath "path"
  "encoding/json"
  "almadash/varc/utils/logger"
  "almadash/varc/utils/str"
  "almadash/varc/utils/file/file"
)

// ================================================
func GetSuffix(type string) {
  var suffix string
  if type == "write" {
    suffix = ".lockw."
  } else if type == "read" {
    suffix = ".lockr."
  } else {
    msg := fmt.Sprintf("INVALID type: %s", type)
    logger.LogErrorAndPanic(errors.New(msg))
  }

  return suffix
}

func ListLocksOfType(path string, type string) {
  pid := os.Getpid()
  suffix := GetSuffix(type)

  dir := fs.GetLockDir()
  basename := gopath.Base(basePath)

  search = basename + suffix

  entries := fs.ReadDir(dir)
  matches := []string{}
  for _, ent := range entries {
    if str.RegexMatch(ent.Name(), search) {
      matches = append(matches, ent.Name())
    }
  }

  if len(matches) < 1 {
    return nil
  }
  return matches
}

func ListWriteLocks(path string) {
  matches := ListLocksOfType(path, "write")
  return matches
}

func ListReadLocks(path string) {
  matches := ListLocksOfType(path, "read")
  return matches
}

// ================================================
/*
  - for manual file based locking
*/
type Lock struct {
  basePath string
  path string
  locked bool
  type string
}

func New(type string, basePath string) Lock {
  out := Lock{}
  out.locked = false
  out.type = false

  pid := os.Getpid()
  suffix := GetSuffix(type)

  out.basePath = basePath

  dir := fs.GetLockDir()
  basename := gopath.Base(basePath)
  out.path = dir + "/" + basename + suffix + pid

  if out.Exists() {
    out.locked = true
  }

  return out
}

// ================================================
func (this *Lock) Exists() bool {
  return fs.FileExists(this.path)
}

func (this *Lock) Lock() {
  if this.Exists() {
    msg := fmt.Sprintf("[error]: %s Lock EXISTS", this.type)
    logger.LogErrorAndPanic(errors.New(msg))
  }
  fs.TouchFile(this.path)
}

func (this *Lock) Unlock() {
  fs.FileDelete(this.path)
}
