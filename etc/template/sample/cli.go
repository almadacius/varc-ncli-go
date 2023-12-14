/*
  this sample includes basic structure for option handling
  might not be used on this program in particular.
*/
package main

import (
  "flag"
  "fmt"
  "os"
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
    metricChoices := map[string]bool{
      "chars": true, "words": true, "lines": true, "substring": true
    }
    if _, validChoice := metricChoices[*countMetricPtr]; !validChoice {
      countCommand.PrintDefaults()
      os.Exit(1)
    }
    //Print
    fmt.Printf("textPtr: %s, metricPtr: %s, substringPtr: %v, uniquePtr: %t\n", *countTextPtr, *countMetricPtr, *countSubstringPtr, *countUniquePtr)
  }
}

func list(options []string) {
  listCommand := flag.NewFlagSet("list", flag.ExitOnError)

  // List subcommand flag pointers
  listTextPtr := listCommand.String("text", "", "Text to parse. (Required)")
  listMetricPtr := listCommand.String("metric", "chars", "Metric <chars|words|lines>. (Required)")
  listUniquePtr := listCommand.Bool("unique", false, "Measure unique values of a metric.")

  listCommand.Parse(options)

  // Check which subcommand was Parsed using the FlagSet.Parsed() function. Handle each case accordingly.
  // FlagSet.Parse() will evaluate to false if no flags were parsed (i.e. the user did not provide any flags)
  if listCommand.Parsed() {
    // Required Flags
    if *listTextPtr == "" {
      listCommand.PrintDefaults()
      os.Exit(1)
    }
    //Choice flag
    metricChoices := map[string]bool{
      "chars": true, "words": true, "lines": true
    }
    if _, validChoice := metricChoices[*listMetricPtr]; !validChoice {
      listCommand.PrintDefaults()
      os.Exit(1)
    }
    // Print
    fmt.Printf("textPtr: %s, metricPtr: %s, uniquePtr: %t\n", *listTextPtr, *listMetricPtr, *listUniquePtr)
  }
}

// ================================================
func main() {
  // Verify that a subcommand has been provided
  // os.Arg[0] is the main command
  // os.Arg[1] will be the subcommand
  if len(os.Args) < 2 {
    fmt.Println("list or count subcommand is required")
    os.Exit(1)
  }

  options := os.Args[2:]

  // Switch on the subcommand
  // Parse the flags for appropriate FlagSet
  // FlagSet.Parse() requires a set of arguments to parse as input
  // os.Args[2:] will be all arguments starting after the subcommand at os.Args[1]
  switch os.Args[1] {
  case "list":
    list(options)
  case "count":
    count(options)
  default:
    flag.PrintDefaults()
    os.Exit(1)
  }
}
