package timercmd

import (
  "almadash/varc/utils/logger"
  "almadash/varc/controller/timercon"
)

// ================================================
type TimerCreateCmd struct {}

func (c *TimerCreateCmd) Declare() {
  // no flags
}

func (c *TimerCreateCmd) Run(options []string) {
  timer := timercon.NewTimer()
  logger.Output(timer.GetKey())
}
