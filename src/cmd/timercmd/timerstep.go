package timercmd

import (
  "fmt"
  "os"
  "errors"
  "almadash/varc/utils/logger"
  "almadash/varc/cmd/cmdlib"
  "almadash/varc/controller/timercon"
)

// ================================================
type TimerStepCmd struct {
  cmdlib.Command
}

func (c *TimerStepCmd) Run(options []string) {
  if len(options) < 1 {
    fmt.Println("timerstep <timerkey>")
    os.Exit(1)
  }
  timerkey := options[0]
  if timerkey == "" {
    logger.LogErrorAndPanic(errors.New("timerkey NOT provided"))
  }

  timer := timercon.NewTimerFromKey(timerkey)
  diff := timer.Step()

  logger.Output(diff)
}
