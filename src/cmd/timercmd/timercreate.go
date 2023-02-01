package timercmd

import (
  "almadash/varc/utils"
  "almadash/varc/controller/timercon"
)

// ================================================
type TimerCreateCmd struct {}

func (c *TimerCreateCmd) Declare() {
  // no flags
}

func (c *TimerCreateCmd) Run(options []string) {
  timer := timercon.CreateNew()

  utils.Output(timer.GetKey())
}
