package cmd

import (
  "fmt"
  "os"
  "path"
)

// ================================================
type ListCmd struct {

}

func (c *ListCmd) Declare() {
  // no flags
}

func (c *ListCmd) Run(options []string) {
  scopeName := options[0]

  fmt.Println("scopeName ", scopeName)

  execPath, err := os.Executable()
  if err != nil {
    fmt.Println("Failed to get program path")
    os.Exit(1)
  }

  dir := path.Dir(execPath)
  fmt.Println("path: ", dir)
}
