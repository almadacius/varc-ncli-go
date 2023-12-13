package file

import (
  "os"
  "errors"
  "almadash/varc/utils/logger"
  "almadash/varc/utils/fs"
)

// ================================================
/*
  - a basic file class
  - no lock, just basic interface
*/
type File struct {
  path string
  instance *os.File
}

func New(path string) File {
  file := File{}
  file.path = path
  return file
}

// ================================================
func (this *File) OpenWrite(mode os.FileMode) {
  instance := fs.OpenWrite(this.path, mode)
  this.instance = instance
}

func (this *File) OpenRead() {
  instance := fs.OpenRead(this.path)
  this.instance = instance
}

func (this *File) GetFd() int {
  file := this.instance
  fd := int(file.Fd())
  return fd
}

func (this *File) EnsureOpen() {
  file := this.instance
  if file == nil {
    logger.LogErrorAndPanic(errors.New("file NOT open"))
  }
}

// ================================================
func (this *File) Exists() bool {
  return fs.FileExists(this.path)
}

func (this *File) Close() {
  file := this.instance
  file.Close()
}

func (this *File) Write(data []byte) (int, error) {
  this.EnsureOpen()
  file := this.instance
  return file.Write(data)
}

func (this *File) Read() []byte {
  this.EnsureOpen()
  file := this.instance
  buffer := make([]byte, 1024)
  amount, err := file.Read(buffer)
  logger.LogErrorAndPanic(err)
  out := buffer[0:amount]
  return out
}

func (this *File) Delete() {
  if this.Exists() {
    fs.FileDelete(this.path)
  }
}
