package varcmd

import (
  "fmt"
  "os"
  "errors"
  "almadash/varc/utils/logger"
  "almadash/varc/controller/scopecon"
)

// ================================================
type KeysCmd struct {}

func (c *KeysCmd) Declare() {
  // no flags
}

func (c *KeysCmd) Run(options []string) {
  if len(options) < 1 {
    fmt.Println("keys <scopeName>")
    os.Exit(1)
  }

  scopeName := options[0]

  if scopeName == "" {
    logger.LogErrorAndPanic(errors.New("scopeName not provided"))
  }

  scope := scopecon.NewScope(scopeName)

  values := scope.GetKeys()
  bytes := logger.StringListToBytes(values)

  os.Stdout.Write(bytes)
}
