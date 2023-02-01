package timercon

import (
  "fmt"
  "errors"
  "time"
  "strconv"
  "almadash/varc/utils"
  "almadash/varc/utils/jsonfile"
)

// ================================================
/*
  @behav - automatically clears out any lingering timers that are older
    than a day to keep clean and simple
*/
type Timer struct {
  key string
  path string
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
  timer.key = strconv.FormatInt(time.Now().Unix(), 10)

  timer.persist()

  return timer
}

func FromKey(key string) Timer {
  timer := createBaseTimer()
  timer.key = key

  exists := timer.keyExists()

  if ! exists {
    message := fmt.Sprintf("timerkey does not exist: %s", key)
    utils.LogErrorAndPanic(errors.New(message))
  }

  return timer
}

// ================================================
func (t *Timer) getKeyAsTime() time.Time {
  num, err := strconv.ParseInt(t.key, 10, 0)
  utils.LogErrorAndPanic(err)
  timeinst := time.Unix(num, 0)

  return timeinst
}

func (t *Timer) GetKey() string {
  return t.key
}

// ================================================
func (t *Timer) Step() time.Duration {
  t1 := t.getKeyAsTime()
  t2 := time.Now()

  diff := t2.Sub(t1)

  return diff
}

func (t *Timer) End() time.Duration {
  t.clearFromFile()

  return t.Step()
}

// ================================================
func (t *Timer) persist() {
  data := t.load()

  data[t.key] = ""

  jsonfile.Save(t.path, data)
}

func (t *Timer) clearFromFile() {
  data := t.load()

  delete(data, t.key)

  jsonfile.Save(t.path, data)
}

func (t *Timer) keyExists() bool {
  data := t.load()

  _, exists := data[t.key]

  return exists
}

// ================================================
func (t *Timer) load() map[string] string {
  data := make(map[string] string)
  if utils.FileExists(t.path) {
    data = jsonfile.Load(t.path)
  }

  return data
}
