package str

import (
  "fmt"
  "regexp"
  "errors"
  "almadash/varc/utils/logger"
)

// ================================================
func RegexMatch(str string, patt string) bool {
  re, err := regexp.Compile(patt)
  if err != nil {
    msg := fmt.Sprintf("regex compile FAIL: %s", patt)
    logger.LogErrorAndPanic(errors.New(msg))
  }

  return re.MatchString(str)
}
