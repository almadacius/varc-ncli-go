package cmd

import (
  "fmt"
  "os"
  "errors"
  "almadash/varc/utils"
)

// ================================================
type SetCmd struct {}

func (c *SetCmd) Declare() {
  // no flags
}

func (c *SetCmd) Run(options []string) {
  if len(options) < 2 {
    fmt.Println("get <scopePath> <value>")
    os.Exit(1)
  }

  scopePath := options[0]
  newValue := options[1]

  if scopePath == "" {
    utils.LogErrorAndPanic(errors.New("scopePath not provided"))
  }
  if newValue == "" {
    utils.LogErrorAndPanic(errors.New("value not provided"))
  }

  scopeName, varName := utils.ParseScope(scopePath)

  scope := utils.NewScope(scopeName)

  scope.SetVar(varName, newValue)
}
