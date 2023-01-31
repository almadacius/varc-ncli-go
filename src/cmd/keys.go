package cmd

import (
  "fmt"
  "os"
  "errors"
  "strings"
  "almadash/varc/utils"
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
    utils.LogErrorAndPanic(errors.New("scopeName not provided"))
  }

  scope := utils.NewScope(scopeName)

  values := scope.GetKeys()
  joined := strings.Join(values, "\n")
  bytes := []byte(joined)

  os.Stdout.Write(bytes)
}
