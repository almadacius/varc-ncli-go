package main

import (
  "flag"
  "fmt"
  "os"
  "almadash/varc/cmd"
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
  case "set":
    cmd := cmd.SetCmd{}
    cmd.Run(options)
  case "unset":
    cmd := cmd.UnsetCmd{}
    cmd.Run(options)
  case "get":
    cmd := cmd.GetCmd{}
    cmd.Run(options)
  case "keys":
    cmd := cmd.KeysCmd{}
    cmd.Run(options)
  case "scopes":
    cmd := cmd.ScopesCmd{}
    cmd.Run(options)
  case "del":
    cmd := cmd.DelCmd{}
    cmd.Run(options)
  default:
    fmt.Println("command is not supported: ", cmdName)
    flag.PrintDefaults()
    os.Exit(1)
  }
}
