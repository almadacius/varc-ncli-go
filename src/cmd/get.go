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
  scopeName := options[0]

  fmt.Println("scopeName ", scopeName)

  dir := utils.GetDirname()

  fmt.Println("path: ", dir)
}
