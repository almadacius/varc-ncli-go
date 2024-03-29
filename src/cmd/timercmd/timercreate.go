package timercmd

import (
  "almadash/varc/utils/logger"
  "almadash/varc/cmd/cmdlib"
  "almadash/varc/controller/timercon"
)

// ================================================
type TimerCreateCmd struct {
  cmdlib.Command
}

func (c *TimerCreateCmd) Run(options []string) {
  timer := timercon.NewTimer()
  logger.Output(timer.GetKey())
}
