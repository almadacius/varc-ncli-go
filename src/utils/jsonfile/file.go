package jsonfile

import (
  "os"
  "syscall"
  "almadash/varc/utils"
)

// ================================================
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
  utils.LogErrorAndPanic(err)
  this.instance = file
}

func (this *File) OpenRead() {
  flags := os.O_RDONLY
  file, err := os.OpenFile(this.path, flags, 0444)
  utils.LogErrorAndPanic(err)
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
  utils.LogErrorAndPanic(err)
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
  utils.LogErrorAndPanic(err)
  return buffer
}
