package cmd

import (
  "fmt"
  "os"
  "errors"
  "almadash/varc/utils"
)

// ================================================
type GetCmd struct {}

func (c *GetCmd) Declare() {
  // no flags
}

func (c *GetCmd) Run(options []string) {
  if len(options) < 1 {
    fmt.Println("get <scopePath>")
    os.Exit(1)
  }

  scopePath := options[0]

  if scopePath == "" {
    utils.LogErrorAndPanic(errors.New("scopePath not provided"))
  }

  scopeName, varName := utils.ParseScope(scopePath)

  scope := utils.NewScope(scopeName)

  value := scope.GetVar(varName)

  os.Stdout.Write([]byte(value))
}
