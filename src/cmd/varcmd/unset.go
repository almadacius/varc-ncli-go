package varcmd

import (
  "fmt"
  "os"
  "errors"
  "almadash/varc/utils/logger"
  "almadash/varc/cmd/cmdlib"
  "almadash/varc/controller/scopecon"
)

// ================================================
type UnsetCmd struct {
  cmdlib.Command
}

func (c *UnsetCmd) Run(options []string) {
  if len(options) < 1 {
    fmt.Println("unset <scopePath>")
    os.Exit(1)
  }

  scopePath := options[0]

  if scopePath == "" {
    logger.LogErrorAndPanic(errors.New("scopePath not provided"))
  }

  scopeName, varName := scopecon.ParseScope(scopePath)

  scope := scopecon.NewScope(scopeName)

  scope.UnsetVar(varName)
}
