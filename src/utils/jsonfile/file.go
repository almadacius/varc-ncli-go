package jsonfile

import (
  "os"
  "syscall"
  "almadash/varc/utils/logger"
)

// ================================================
func writeFileWithLock(path string, data []byte, mode os.FileMode) {
  file := CreateFile(path)
  file.OpenWrite(mode)
  defer file.Close()
  file.Lock()
  defer file.Unlock()

  _, err := file.Write(data)
  logger.LogErrorAndPanic(err)
}

func readFileWithLock(path string) []byte {
  file := CreateFile(path)
  file.OpenRead()
  defer file.Close()
  file.Lock()
  defer file.Unlock()

  // not really sure this is respecting the lock
  // still getting some eventual read errors on the tests
  data, err := os.ReadFile(path)
  logger.LogErrorAndPanic(err)

  return data
}

// ================================================
/*
  - file class with os level syscall based Lock
  - still getting some file corruption with this approach
*/
type File struct {
  path string
  instance *os.File
}

func CreateFile(path string) File {
  file := File{}
  file.path = path
  return file
}

// ================================================
func (this *File) OpenWrite(mode os.FileMode) {
  flags := os.O_WRONLY | os.O_CREATE | os.O_TRUNC
  file, err := os.OpenFile(this.path, flags, mode)
  logger.LogErrorAndPanic(err)
  this.instance = file
}

func (this *File) OpenRead() {
  flags := os.O_RDONLY
  file, err := os.OpenFile(this.path, flags, 0444)
  logger.LogErrorAndPanic(err)
  this.instance = file
}

func (this *File) GetFd() int {
  file := this.instance
  fd := int(file.Fd())
  return fd
}

func (this *File) Lock() {
  fd := this.GetFd()
  // exclusive file lock
  err := syscall.Flock(fd, syscall.LOCK_EX)
  logger.LogErrorAndPanic(err)
}

func (this *File) Unlock() {
  fd := this.GetFd()
  syscall.Flock(fd, syscall.LOCK_UN)
}

// ================================================
func (this *File) Close() {
  file := this.instance
  file.Close()
}

func (this *File) Write(data []byte) (int, error) {
  file := this.instance
  return file.Write(data)
}

func (this *File) Read() []byte {
  file := this.instance
  buffer := make([]byte, 1024)
  // first is amount of bytes read
  _, err := file.Read(buffer)
  logger.LogErrorAndPanic(err)
  return buffer
}
