package utils

import (
  "os"
  "path"
  "fmt"
)

// ================================================
func DirExists(path string) bool {
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

// ================================================
func GetDirname() string {
  execPath, err := os.Executable()
  if err != nil {
    fmt.Println("Failed to get program path")
    os.Exit(1)
  }

  dir := path.Dir(execPath)

  return dir
}

func GetStorageDir() string {
  execDir := GetDirname()
  storageDir := execDir + "/varcgo-storage"

  EnsureDir(storageDir)

  return storageDir
}
