package time

import (
  "time"
  "strconv"
)

// ================================================
func GetNow() time.Time {
  return time.Now()
}

func GetTimestamp() string {
  number := GetNow().UnixMilli()
  stamp := strconv.FormatInt(number, 10)
  return stamp
}
