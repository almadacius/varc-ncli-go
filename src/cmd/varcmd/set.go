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
type SetCmd struct {
  cmdlib.Command
}

func (c *SetCmd) Run(options []string) {
  if len(options) < 2 {
    fmt.Println("set <scopePath> <value>")
    os.Exit(1)
  }

  scopePath := options[0]
  newValue := options[1]

  if scopePath == "" {
    logger.LogErrorAndPanic(errors.New("[set var]: scopePath not provided"))
  }
  if newValue == "" {
    message := fmt.Sprintf("[set var '%s']: value not provided", scopePath)
    logger.LogErrorAndPanic(errors.New(message))
  }

  scopeName, varName := scopecon.ParseScope(scopePath)

  scope := scopecon.NewScope(scopeName)

  scope.SetVar(varName, newValue)
}
