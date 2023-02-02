package varcmd

import (
  "fmt"
  "os"
  "errors"
  "almadash/varc/utils"
  "almadash/varc/controller/scopecon"
)

// ================================================
type SetCmd struct {}

func (c *SetCmd) Declare() {
  // no flags
}

func (c *SetCmd) Run(options []string) {
  if len(options) < 2 {
    fmt.Println("set <scopePath> <value>")
    os.Exit(1)
  }

  scopePath := options[0]
  newValue := options[1]

  if scopePath == "" {
    utils.LogErrorAndPanic(errors.New("[set var]: scopePath not provided"))
  }
  if newValue == "" {
    message := fmt.Sprintf("[set var '%s']: value not provided", scopePath)
    utils.LogErrorAndPanic(errors.New(message))
  }

  scopeName, varName := scopecon.ParseScope(scopePath)

  scope := scopecon.NewScope(scopeName)

  scope.SetVar(varName, newValue)
}
