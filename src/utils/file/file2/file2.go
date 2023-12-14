package file2

import (
  "time"
  "almadash/varc/utils/promise"
  "almadash/varc/utils/file/file"
  "almadash/varc/utils/file/lock"
)

type File = file.File
type Lock = lock.Lock

// ================================================
/*
  - supports manual file based locking
*/
type File2 struct {
  File
  lockW Lock
  lockR Lock
}

func New(path string) File2 {
  out := File2{}
  out.File = file.New(path)

  out.lockW = lock.New("write", path)
  out.lockR = lock.New("read", path)
  return out
}

// ================================================
func (this *File2) Save(data []byte) {
  maxTries := 100000
  interval := 200 * time.Millisecond

  path := this.GetPath()

  // cannot write if someone is writing or reading
  promise.WaitFor(func() bool {
    writeLocks := lock.ListWriteLocks(path)
    if writeLocks != nil {
      return false
    }

    readLocks := lock.ListReadLocks(path)
    if readLocks != nil {
      return false
    }

    return true
  }, maxTries, interval)

  wLock := this.lockW

  wLock.Lock()
  defer wLock.Unlock()

  this.OpenWrite()
  defer this.Close()

  this.Write(data)
}

func (this *File2) Load() []byte {
  maxTries := 100000
  interval := 200 * time.Millisecond

  path := this.GetPath()

  // cannot read if someone is writing
  promise.WaitFor(func() bool {
    writeLocks := lock.ListWriteLocks(path)
    if writeLocks != nil {
      return false
    }

    return true
  }, maxTries, interval)

  rLock := this.lockR

  rLock.Lock()
  defer rLock.Unlock()

  this.OpenRead()
  defer this.Close()

  data := this.Read()
  return data
}
