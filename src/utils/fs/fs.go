package fs

import (
  "os"
  "path"
  "path/filepath"
  "fmt"
  gofs "io/fs"
  "io/ioutil"
  "time"
  "errors"
  "almadash/varc/utils/logger"
  "almadash/varc/utils/promise"
)

// ================================================
func DirExists(path string) bool {
  _, err := os.Stat(path)
  if os.IsNotExist(err) {
    return false
  }
  return true
}

func FileExists(path string) bool {
  _, err := os.Stat(path)
  if os.IsNotExist(err) {
    return false
  }
  return true
}

func isDir(path string) bool {
  stat, err := os.Lstat(path)
  if os.IsNotExist(err) {
    return false
  }
  mode := stat.Mode()
  if ! mode.IsDir() {
    return false
  }
  return true
}

func EnsureDir(path string) {
  exists := DirExists(path)
  if ! exists {
    os.Mkdir(path, 0755)
  }
}

func ReadDir(path string) []gofs.DirEntry {
  entries, err := os.ReadDir(path)
  if err != nil {
    msg := fmt.Sprintf("read dir FAILED: %s", path)
    logger.LogErrorAndPanic(errors.New(msg))
  }
  return entries
}

func ReadDir2(path string) []gofs.FileInfo {
  entries, err := ioutil.ReadDir(path)
  logger.LogErrorAndPanic(err)
  return entries
}

// ================================================
func GetDirname() string {
  execPath, err := os.Executable()
  if err != nil {
    fmt.Println("Failed to get program path")
    os.Exit(1)
  }

  execPath, err = filepath.EvalSymlinks(execPath)
  logger.LogErrorAndPanic(err)

  dir := path.Dir(execPath)

  return dir
}

func GetStorageDir() string {
  execDir := GetDirname()
  storageDir := execDir + "/varcgo-storage"

  EnsureDir(storageDir)

  return storageDir
}

func GetTimerFile() string {
  execDir := GetDirname()
  timerDir := execDir + "/varcgo-timer"
  timerFile := timerDir + "/timer.json"

  EnsureDir(timerDir)

  return timerFile
}

func GetTestLockFile() string {
  dir := GetStorageDir()
  filePath := dir + "/test-lock.json"

  EnsureDir(dir)

  return filePath
}

func GetLockDir() string {
  dir := GetDirname()
  dir = dir + "/locks"

  EnsureDir(dir)

  return dir
}

// ================================================
func OpenWrite(path string, mode os.FileMode) *os.File {
  flags := os.O_WRONLY | os.O_CREATE | os.O_TRUNC
  file, err := os.OpenFile(path, flags, mode)
  logger.LogErrorAndPanic(err)
  return file
}

func OpenRead(path string) *os.File {
  flags := os.O_RDONLY
  file, err := os.OpenFile(path, flags, 0444)
  logger.LogErrorAndPanic(err)
  return file
}

func FileDelete(path string) {
  err := os.Remove(path)
  logger.LogErrorAndPanic(err)
}

func TouchFile(path string) {
  file := OpenWrite(path, 0644)
  file.Close()
}

// ================================================
func UntilFileExists(path string) {
  maxTries := 100000
  interval := 200 * time.Millisecond
  promise.WaitFor(func() bool {
    return FileExists(path)
  }, maxTries, interval)
}

func UntilFileDoesNotExist(path string) {
  maxTries := 100000
  interval := 200 * time.Millisecond
  promise.WaitFor(func() bool {
    return !FileExists(path)
  }, maxTries, interval)
}
