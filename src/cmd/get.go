package cmd

import (
  "fmt"
  "almadash/varc/utils"
)

// ================================================
type ListCmd struct {

}

func (c *ListCmd) Declare() {
  // no flags
}

func (c *ListCmd) Run(options []string) {
  scopePath := options[0]

  scopeName, varName := utils.ParseScope(scopePath)
  fmt.Println("out", scopeName, varName)

  scope := utils.NewScope(scopeName)
  fmt.Println("aa", scope)
}
