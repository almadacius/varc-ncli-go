package slice

import (
  "fmt"
  "errors"
  "almadash/varc/utils/logger"
)

// ================================================
func RemoveIndex(s []int, index int) []int {
  out := make([]int, 0)
  out = append(out, s[0:index]...)
  out = append(out, s[index+1:]...)
  return out
}

func RemoveFirst(s []int, item int) []int {
  for i, v := range s {
    if v == item {
      out := RemoveIndex(s, i)
      return out
    }
  }
  msg := fmt.Sprintf("item NOT found: %i", item)
  logger.LogErrorAndPanic(errors.New(msg))
  return nil
}
