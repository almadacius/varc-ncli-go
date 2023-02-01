package cmd

import (
  "os"
  "almadash/varc/utils"
)

// ================================================
type ScopesCmd struct {}

func (c *ScopesCmd) Declare() {
  // no flags
}

func (c *ScopesCmd) Run(options []string) {
  scopes := utils.ListScopes()

  bytes := utils.StringListToBytes(scopes)

  os.Stdout.Write(bytes)
}
