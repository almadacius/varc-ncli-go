package lock

import (
  "fmt"
  "os"
  "errors"
  gopath "path"
  "almadash/varc/utils/logger"
  "almadash/varc/utils/fs"
  "almadash/varc/utils/str"
)

// ================================================
func GetSuffix(typeName string) string {
  var suffix string
  if typeName == "write" {
    suffix = ".lockw."
  } else if typeName == "read" {
    suffix = ".lockr."
  } else {
    msg := fmt.Sprintf("INVALID type: %s", typeName)
    logger.LogErrorAndPanic(errors.New(msg))
  }

  return suffix
}

func ListLocksOfType(path string, typeName string) []string {
  suffix := GetSuffix(typeName)

  dir := fs.GetLockDir()
  basename := gopath.Base(path)

  search := basename + suffix

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

func ListWriteLocks(path string) []string {
  matches := ListLocksOfType(path, "write")
  return matches
}

func ListReadLocks(path string) []string {
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
  typeName string
}

func New(typeName string, basePath string) Lock {
  out := Lock{}
  out.locked = false
  out.typeName = ""

  pid := os.Getpid()
  suffix := GetSuffix(typeName)

  out.basePath = basePath

  dir := fs.GetLockDir()
  basename := gopath.Base(basePath)
  out.path = dir + "/" + basename + suffix + string(pid)

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
    msg := fmt.Sprintf("[error]: %s Lock EXISTS", this.typeName)
    logger.LogErrorAndPanic(errors.New(msg))
  }
  fs.TouchFile(this.path)
}

func (this *Lock) Unlock() {
  fs.FileDelete(this.path)
}
