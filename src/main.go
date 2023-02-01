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
  case "set":
    cmd := varcmd.SetCmd{}
    cmd.Run(options)
  case "unset":
    cmd := varcmd.UnsetCmd{}
    cmd.Run(options)
  case "get":
    cmd := varcmd.GetCmd{}
    cmd.Run(options)
  case "keys":
    cmd := varcmd.KeysCmd{}
    cmd.Run(options)
  case "scopes":
    cmd := varcmd.ScopesCmd{}
    cmd.Run(options)
  case "del":
    cmd := varcmd.DelCmd{}
    cmd.Run(options)

  case "timercreate":
    cmd := timercmd.TimerCreateCmd{}
    cmd.Run(options)
  case "timerstep":
    cmd := timercmd.TimerStepCmd{}
    cmd.Run(options)
  case "timerend":
    cmd := timercmd.TimerEndCmd{}
    cmd.Run(options)
  case "timerprune":
    cmd := timercmd.TimerPruneCmd{}
    cmd.Run(options)

  default:
    fmt.Println("command is not supported: ", cmdName)
    flag.PrintDefaults()
    os.Exit(1)
  }
}
