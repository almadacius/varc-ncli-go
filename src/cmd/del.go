package cmd

import (
  "fmt"
  "os"
  "errors"
  "almadash/varc/utils"
)

// ================================================
type DelCmd struct {}

func (c *DelCmd) Declare() {
  // no flags
}

func (c *DelCmd) Run(options []string) {
  if len(options) < 1 {
    fmt.Println("del <scopeName>")
    os.Exit(1)
  }

  scopeName := options[0]

  if scopeName == "" {
    utils.LogErrorAndPanic(errors.New("scopeName not provided"))
  }

  scope := utils.NewScope(scopeName)
  scope.Delete()
}
