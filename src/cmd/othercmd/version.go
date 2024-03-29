package othercmd

import (
  "fmt"
  "almadash/varc/cmd/cmdlib"
  "almadash/varc/auto"
)

// ================================================
type VersionCmd struct {
  cmdlib.Command
}

func (c *VersionCmd) Run(options []string) {
  fmt.Println("varc-ncli-go")
  fmt.Println("version: ", auto.Version)
}
