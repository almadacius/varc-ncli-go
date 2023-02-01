package varcmd

import (
  "os"
  "almadash/varc/utils"
  "almadash/varc/controller/scopecon"
)

// ================================================
type ScopesCmd struct {}

func (c *ScopesCmd) Declare() {
  // no flags
}

func (c *ScopesCmd) Run(options []string) {
  scopes := scopecon.ListScopes()

  bytes := utils.StringListToBytes(scopes)

  os.Stdout.Write(bytes)
}
