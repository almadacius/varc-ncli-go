package convert

import (
  "time"
  "strconv"
  "almadash/varc/utils/logger"
)

// ================================================
func StrToTime(key string) time.Time {
  num, err := strconv.ParseInt(key, 10, 64)
  logger.LogErrorAndPanic(err)
  timeinst := time.UnixMilli(num)

  return timeinst
}

func StrToInt(val string) int64 {
  num, err := strconv.ParseInt(val, 10, 0)
  logger.LogErrorAndPanic(err)
  return num
}

func IntToStr(val int64) string {
  return strconv.FormatInt(val, 10)
}
