package main

import (
  "flag"
  "fmt"
  "os"
  "almadash/varc/cmd/varcmd"
  "almadash/varc/cmd/timercmd"
)

// ================================================
func main() {
  // Verify that a subcommand has been provided
  // os.Arg[0] is the main command
  // os.Arg[1] will be the subcommand
  if len(os.Args) < 2 {
    fmt.Println("get subcommand is required")
    os.Exit(1)
  }

  cmdName := os.Args[1]
  options := os.Args[2:]

  switch cmdName {
  // variable management
  case "set": cmd := varcmd.SetCmd{}
  case "unset": cmd := varcmd.UnsetCmd{}
  case "get": cmd := varcmd.GetCmd{}
  case "keys": cmd := varcmd.KeysCmd{}
  case "scopes": cmd := varcmd.ScopesCmd{}
  case "del": cmd := varcmd.DelCmd{}

  case "timercreate": cmd := timercmd.TimerCreateCmd{}
  case "timerstep": cmd := timercmd.TimerStepCmd{}
  case "timerend": cmd := timercmd.TimerEndCmd{}
  case "timerprune": cmd := timercmd.TimerPruneCmd{}

  default:
    fmt.Println("command is not supported: ", cmdName)
    flag.PrintDefaults()
    os.Exit(1)
  }

  cmd.Run(options)
  os.Exit(0)
}
