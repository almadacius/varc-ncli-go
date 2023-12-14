package timercmd

import (
  "almadash/varc/cmd/cmdlib"
  "almadash/varc/controller/timercon"
)

// ================================================
type TimerPruneCmd struct {
  cmdlib.Command
}

func (c *TimerPruneCmd) Run(options []string) {
  timer := timercon.NewTimer()
  timer.PruneOldStamps()
}
