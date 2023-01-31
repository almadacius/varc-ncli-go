package main

import (
  "flag"
  "fmt"
  "os"
  "almadash/varc/cmd"
)

// ================================================
func count(options []string) {
  countCommand := flag.NewFlagSet("count", flag.ExitOnError)

  // Count subcommand flag pointers
  // Adding a new choice for --metric of 'substring' and a new --substring flag
  countTextPtr := countCommand.String("text", "", "Text to parse. (Required)")
  countMetricPtr := countCommand.String("metric", "chars", "Metric {chars|words|lines|substring}. (Required)")
  countSubstringPtr := countCommand.String("substring", "", "The substring to be counted. Required for --metric=substring")
  countUniquePtr := countCommand.Bool("unique", false, "Measure unique values of a metric.")

  countCommand.Parse(options)

  if countCommand.Parsed() {
    // Required Flags
    if *countTextPtr == "" {
      countCommand.PrintDefaults()
      os.Exit(1)
    }
    // If the metric flag is substring, the substring flag is required
    if *countMetricPtr == "substring" && *countSubstringPtr == "" {
      countCommand.PrintDefaults()
      os.Exit(1)
    }
    //If the metric flag is not substring, the substring flag must not be used
    if *countMetricPtr != "substring" && *countSubstringPtr != "" {
      fmt.Println("--substring may only be used with --metric=substring.")
      countCommand.PrintDefaults()
      os.Exit(1)
    }
    //Choice flag
    metricChoices := map[string]bool{"chars": true,
      "words": true, "lines": true, "substring": true}
    if _, validChoice := metricChoices[*countMetricPtr]; !validChoice {
      countCommand.PrintDefaults()
      os.Exit(1)
    }
    //Print
    fmt.Printf("textPtr: %s, metricPtr: %s, substringPtr: %v, uniquePtr: %t\n", *countTextPtr, *countMetricPtr, *countSubstringPtr, *countUniquePtr)
  }
}

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
    listCmd := cmd.SetCmd{}
    listCmd.Declare()
    listCmd.Run(options)
  default:
    fmt.Println("command is not supported: ", cmdName)
    flag.PrintDefaults()
    os.Exit(1)
  }
}
