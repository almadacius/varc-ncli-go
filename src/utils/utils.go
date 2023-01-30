package utils

import (
  "os"
  "path"
  "fmt"
)

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
