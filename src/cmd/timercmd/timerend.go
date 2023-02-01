package timercmd

import (
  "fmt"
  "os"
  "errors"  
  "almadash/varc/utils"
  "almadash/varc/controller/timercon"
)

// ================================================
type TimerEndCmd struct {}

func (c *TimerEndCmd) Declare() {
  // no flags
}

func (c *TimerEndCmd) Run(options []string) {
  if len(options) < 1 {
    fmt.Println("timerend <timerkey>")
    os.Exit(1)
  }
  timerkey := options[0]
  if timerkey == "" {
    utils.LogErrorAndPanic(errors.New("timerkey not provided"))
  }

  timer := timercon.FromKey(timerkey)
  diff := timer.End()

  utils.Output(diff)
}
