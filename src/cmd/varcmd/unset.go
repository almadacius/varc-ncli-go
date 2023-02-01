package varcmd

import (
  "fmt"
  "os"
  "errors"
  "almadash/varc/utils"
  "almadash/varc/controller/scopecon"
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

  scopeName, varName := scopecon.ParseScope(scopePath)

  scope := scopecon.NewScope(scopeName)

  scope.UnsetVar(varName)
}
