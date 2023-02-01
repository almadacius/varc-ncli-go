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
func strToTime(key string) time.Time {
  num, err := strconv.ParseInt(key, 10, 0)
  utils.LogErrorAndPanic(err)
  timeinst := time.Unix(num, 0)

  return timeinst
}

func load() map[string] string {
  path := utils.GetTimerFile()

  data := make(map[string] string)
  if utils.FileExists(path) {
    data = jsonfile.Load(path)
  }

  return data
}

// ================================================
func PruneOldStamps() {
  path := utils.GetTimerFile()

  now := time.Now()
  day := time.Hour * 24

  yesterday := now.Add(-day)

  data := load()

  for key, _ := range data {
    stamp := strToTime(key)
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
  return strToTime(t.key)
}

func (t *Timer) GetKey() string {
  return t.key
}

// ================================================
func (t *Timer) Step() string {
  t1 := t.getKeyAsTime()
  t2 := time.Now()

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

  data[t.key] = ""

  jsonfile.Save(t.path, data)
}

func (t *Timer) clearFromFile() {
  data := load()

  delete(data, t.key)

  jsonfile.Save(t.path, data)
}

func (t *Timer) keyExists() bool {
  data := load()

  _, exists := data[t.key]

  return exists
}
