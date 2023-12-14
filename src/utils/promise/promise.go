package promise

import (
  "fmt"
  "time"
  "errors"
  "almadash/varc/utils/logger"
)

// ================================================
func WaitFor(fn func() bool, maxTries int, interval time.Duration) {
  tries := 0
  for {
    if(fn()) {
      break
    }

    tries++
    if tries >= maxTries {
      msg := fmt.Sprintf("condition NOT met after %d tries", tries)
      logger.LogErrorAndPanic(errors.New(msg))
      break
    }

    time.Sleep(interval)
  }
}
