package cmd

import (
  "fmt"
  "os"
  "errors"
  "almadash/varc/utils"
)

// ================================================
type UnsetCmd struct {}

func (c *UnsetCmd) Declare() {
  // no flags
}

func (c *UnsetCmd) Run(options []string) {
  if len(options) < 1 {
    fmt.Println("unset <scopePath>")
    os.Exit(1)
  }

  scopePath := options[0]

  if scopePath == "" {
    utils.LogErrorAndPanic(errors.New("scopePath not provided"))
  }

  scopeName, varName := utils.ParseScope(scopePath)

  scope := utils.NewScope(scopeName)

  scope.UnsetVar(varName)
}
