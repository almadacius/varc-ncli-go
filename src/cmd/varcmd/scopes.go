package varcmd

import (
  "os"
  "almadash/varc/utils/logger"
  "almadash/varc/cmd/cmdlib"
  "almadash/varc/controller/scopecon"
)

// ================================================
type ScopesCmd struct {
  cmdlib.Command
}

func (c *ScopesCmd) Run(options []string) {
  scopes := scopecon.ListScopes()

  bytes := logger.StringListToBytes(scopes)

  os.Stdout.Write(bytes)
}
