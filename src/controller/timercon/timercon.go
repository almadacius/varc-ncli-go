package timercon

import (
  "fmt"
  "errors"
  "time"
  "strconv"
  "almadash/varc/utils"
  "almadash/varc/utils/logger"
  "almadash/varc/utils/jsonfile"
)

// ================================================
func getNow() time.Time {
  return time.Now()
}

func getTimestamp() string {
  number := getNow().UnixMilli()
  stamp := strconv.FormatInt(number, 10)
  return stamp
}

func load() map[string] string {
  path := utils.GetTimerFile()

  data := make(map[string] string)
  if utils.FileExists(path) {
    data = jsonfile.Load(path)
  }

  return data
}

func strToTime(key string) time.Time {
  num, err := strconv.ParseInt(key, 10, 64)
  logger.LogErrorAndPanic(err)
  timeinst := time.UnixMilli(num)

  return timeinst
}

func strToInt(val string) int64 {
  num, err := strconv.ParseInt(val, 10, 0)
  logger.LogErrorAndPanic(err)
  return num
}

func intToStr(val int64) string {
  return strconv.FormatInt(val, 10)
}

func getNext() string {
  path := utils.GetTimerFile()

  data := load()
  next := int64(0)

  currNext, hasNext := data["_next"]
  if hasNext {
    next = strToInt(currNext)
  }

  newNext := next + 1
  if newNext > 100000 {
    newNext = 0
  }

  data["_next"] = intToStr(newNext)
  jsonfile.Save(path, data)

  return intToStr(next)
}

// ================================================
func PruneOldStamps() {
  path := utils.GetTimerFile()

  now := time.Now()
  day := time.Hour * 24

  yesterday := now.Add(-day)

  data := load()

  for key, value := range data {
    stamp := strToTime(value)
    if stamp.Before(yesterday) {
      delete(data, key)
    }
  }

  jsonfile.Save(path, data)
}

// ================================================
/*
  @behav - automatically clears out any lingering timers that are older
    than a day to keep clean and simple
*/
type Timer struct {
  key string
  path string
  startTime string
}

// ================================================
func createBaseTimer() Timer {
  path := utils.GetTimerFile()

  timer := Timer{}
  timer.path = path

  return timer
}

func CreateNew() Timer {
  timer := createBaseTimer()
  timer.key = getNext()
  timer.startTime = getTimestamp()

  timer.persist()

  return timer
}

func FromKey(key string) Timer {
  timer := createBaseTimer()
  timer.key = key

  timer.loadStartTime()

  return timer
}

// ================================================
func (t *Timer) getStartTimeAsTime() time.Time {
  return strToTime(t.startTime)
}

func (t *Timer) GetKey() string {
  return t.key
}

// ================================================
func (t *Timer) Step() string {
  t1 := t.getStartTimeAsTime()
  t2 := getNow()

  diff := t2.Sub(t1)

  return diff.String()
}

func (t *Timer) End() string {
  t.clearFromFile()

  return t.Step()
}

// ================================================
func (t *Timer) persist() {
  data := load()

  data[t.key] = t.startTime

  jsonfile.Save(t.path, data)
}

func (t *Timer) clearFromFile() {
  data := load()

  delete(data, t.key)

  jsonfile.Save(t.path, data)
}

func (t *Timer) loadStartTime() {
  data := load()

  startTime, exists := data[t.key]

  if ! exists {
    message := fmt.Sprintf("timerkey does not exist: %s", t.key)
    logger.LogErrorAndPanic(errors.New(message))
  }

  t.startTime = startTime
}
