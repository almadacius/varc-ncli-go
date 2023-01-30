package cmd

import (
  "flag"
  "fmt"
  "os"
)

// ================================================
type ListCmd struct {
  flags *flag.FlagSet
  text *string
  metric *string
  unique *bool
}

func (c *ListCmd) Declare() {
  flags := flag.NewFlagSet("list", flag.ExitOnError)

  // List subcommand flag pointers
  c.text = flags.String("text", "", "Text to parse. (Required)")
  c.metric = flags.String("metric", "chars", "Metric <chars|words|lines>. (Required)")
  c.unique = flags.Bool("unique", false, "Measure unique values of a metric.")

  c.flags = flags
}

func (c *ListCmd) Run(options []string) {
  c.flags.Parse(options)

  // Check which subcommand was Parsed using the FlagSet.Parsed() function. Handle each case accordingly.
  // FlagSet.Parse() will evaluate to false if no flags were parsed (i.e. the user did not provide any flags)
  if c.flags.Parsed() {
    // Required Flags
    if *c.text == "" {
      c.flags.PrintDefaults()
      os.Exit(1)
    }
    //Choice flag
    metricChoices := map[string]bool{"chars": true,
      "words": true,
      "lines": true}
    if _, validChoice := metricChoices[*c.metric]; !validChoice {
      c.flags.PrintDefaults()
      os.Exit(1)
    }
    // Print
    fmt.Printf("textPtr: %s, metricPtr: %s, uniquePtr: %t\n",
      *c.text, *c.metric, *c.unique)
  }
}
