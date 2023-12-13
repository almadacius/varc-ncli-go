package timercmd

import (
  "fmt"
  "os"
  "errors"
  "almadash/varc/cmd/cmdlib"
  "almadash/varc/utils/logger"
  "almadash/varc/controller/timercon"
)

// ================================================
type TimerEndCmd struct {
  cmdlib.Command
}

func (c *TimerEndCmd) Run(options []string) {
  if len(options) < 1 {
    fmt.Println("timerend <timerkey>")
    os.Exit(1)
  }
  timerkey := options[0]
  if timerkey == "" {
    logger.LogErrorAndPanic(errors.New("timerkey NOT provided"))
  }

  timer := timercon.NewTimerFromKey(timerkey)
  diff := timer.End()

  logger.Output(diff)
}
