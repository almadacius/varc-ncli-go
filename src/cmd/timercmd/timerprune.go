package timercmd

import (
  "almadash/varc/controller/timercon"
)

// ================================================
type TimerPruneCmd struct {}

func (c *TimerPruneCmd) Declare() {
  // no flags
}

func (c *TimerPruneCmd) Run(options []string) {
  timercon.PruneOldStamps()
}
