package jsonfile

import (
  "os"
  "time"
  "fmt"
  "errors"
  "almadash/varc/utils/logger"
  "almadash/varc/utils/fs"
)

// ================================================
func writeFile2WithLock(path string, data []byte, mode os.FileMode) {
  file := CreateFile(path)
  file.OpenWrite(mode)
  defer file.Close()
  file.Lock()
  defer file.Unlock()

  _, err := file.Write(data)
  logger.LogErrorAndPanic(err)
}

func readFile2WithLock(path string) []byte {
  file := CreateFile2(path)

  fs.UntilFileDoesNotExist(file.lockPathW)

  file.OpenRead()
  defer file.Close()
  data := file.Read()

  return data
}

func waitForFile2Lock(path string, maxTries int, interval time.Duration) {
  tries := 0
  for {
    if(!fs.FileExists(path)) { break }

    tries++
    if tries >= maxTries {
      msg := fmt.Sprintf("file did NOT unlock after %d tries, %s", tries, path)
      logger.LogErrorAndPanic(errors.New(msg))
    }

    time.Sleep(interval)
  }
}

// ================================================
/*
  - file class with manual lockfile based locking for write
  - simple await approach
*/
type File2 struct {
  path string
  lockPathW string
  lockPathR string
  instance *os.File
}

func CreateFile2(path string) File2 {
  file := File2{}
  file.path = path
  file.lockPathW = path + ".lockw"
  file.lockPathR = path + ".lockr"
  return file
}

// ================================================
func (this *File2) OpenWrite(mode os.FileMode) {
  instance := fs.OpenWrite(this.path, mode)
  this.instance = instance
}

func (this *File2) OpenRead() {
  instance := fs.OpenRead(this.path)
  this.instance = instance
}

func (this *File2) GetFd() int {
  file := this.instance
  fd := int(file.Fd())
  return fd
}

// ================================================
func (this *File2) IsLocked() bool {
  return fs.FileExists(this.lockPathW)
}

func (this *File2) Lock() {
  if this.IsLocked() {
    logger.LogErrorAndPanic(errors.New("file already locked"))
  }
  fs.TouchFile(this.lockPathW)
}

func (this *File2) Unlock() {
  fs.FileDelete(this.lockPathW)
}

// ================================================
func (this *File2) Close() {
  file := this.instance
  file.Close()
}

func (this *File2) Write(data []byte) (int, error) {
  file := this.instance
  return file.Write(data)
}

func (this *File2) Read() []byte {
  file := this.instance
  buffer := make([]byte, 1024)
  // first is amount of bytes read
  _, err := file.Read(buffer)
  logger.LogErrorAndPanic(err)
  return buffer
}
