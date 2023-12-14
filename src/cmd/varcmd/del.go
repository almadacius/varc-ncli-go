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
type DelCmd struct {
  cmdlib.Command
}

func (c *DelCmd) Run(options []string) {
  if len(options) < 1 {
    fmt.Println("del <scopeName>")
    os.Exit(1)
  }

  scopeName := options[0]

  if scopeName == "" {
    logger.LogErrorAndPanic(errors.New("scopeName not provided"))
  }

  scope := scopecon.NewScope(scopeName)
  scope.Delete()
}
