package timercon

import (
  "fmt"
  "errors"
  "time"
  "almadash/varc/utils/fs"
  "almadash/varc/utils/logger"
  varctime "almadash/varc/utils/time"
  "almadash/varc/utils/convert"
  "almadash/varc/utils/file/file"
  "almadash/varc/utils/file/jsondata"
)

// ================================================
type JsonData = jsondata.JsonData
type File = file.File

// ================================================
/*
  @behav - automatically clears out any lingering timers that are older
    than a day to keep clean and simple
*/
type Timer struct {
  key string
  path string
  startTime string
  file File
  data JsonData
}

// ================================================
func newTimerBase() Timer {
  path := fs.GetTimerFile()

  out := Timer{}
  out.path = path
  out.file = file.New(path)
  out.data = jsondata.New()

  return out
}

func NewTimer() Timer {
  out := newTimerBase()

  out.key = out.getNext()
  out.startTime = varctime.GetTimestamp()

  out.persist()

  return out
}

func NewTimerFromKey(key string) Timer {
  out := newTimerBase()
  out.key = key

  out.loadStartTime()

  return out
}

// ================================================
func (this *Timer) getNext() string {
  data := &this.data

  this.load()

  next := int64(0)

  if data.HasKey("_next") {
    next = convert.StrToInt(data.GetString("_next"))
  }

  newNext := next + 1
  if newNext > 100000 {
    newNext = 0
  }

  data.Set("_next", convert.IntToStr(newNext))

  this.save()

  return convert.IntToStr(next)
}

func (this *Timer) getStartTimeAsTime() time.Time {
  return convert.StrToTime(this.startTime)
}

func (this *Timer) GetKey() string {
  return this.key
}

// ================================================
func (this *Timer) Step() string {
  t1 := this.getStartTimeAsTime()
  t2 := varctime.GetNow()

  diff := t2.Sub(t1)

  return diff.String()
}

func (this *Timer) End() string {
  this.clearFromFile()

  return this.Step()
}

// ================================================
func (this *Timer) PruneOldStamps() {
  now := time.Now()
  day := time.Hour * 24

  yesterday := now.Add(-day)

  this.load()

  this.data.ForEachString(func(key string, value string) {
    stamp := convert.StrToTime(value)
    if stamp.Before(yesterday) {
      this.data.Unset(key)
    }
  })

  this.save()
}

// ================================================
func (this *Timer) load() {
  file := &this.file
  data := &this.data

  if file.Exists() {
    file.OpenRead()
    defer file.Close()
    bytes := file.Read()
    data.SetBytes(bytes)
  }
}

func (this *Timer) save() {
  file := &this.file
  data := &this.data

  bytes := data.ToBytes()
  file.OpenWrite()
  defer file.Close()
  file.Write(bytes)
}

func (this *Timer) persist() {
  data := &this.data

  this.load()

  data.Set(this.key, this.startTime)

  this.save()
}

func (this *Timer) clearFromFile() {
  data := &this.data

  this.load()

  data.Unset(this.key)

  this.save()
}

func (this *Timer) loadStartTime() {
  data := &this.data

  this.load()

  if !data.HasKey(this.key) {
    msg := fmt.Sprintf("timerkey does not exist: %s", this.key)
    logger.LogErrorAndPanic(errors.New(msg))
  }

  startTime := data.GetString(this.key)
  this.startTime = startTime
}
