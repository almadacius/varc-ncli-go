package othercmd

import (
  "fmt"
  "almadash/varc/auto"
)

// ================================================
type VersionCmd struct {}

func (c *VersionCmd) Declare() { /* virtual */ }

func (c *VersionCmd) Run(options []string) {
  fmt.Println("varc-ncli-go")
  fmt.Println("version: ", auto.Version)
}
