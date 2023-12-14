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
func (this *File) Save(data []byte) {
  this.OpenWrite()
  defer this.Close()

  this.Write(data)
}

func (this *File) Load() []byte {
  this.OpenRead()
  defer this.Close()

  data := this.Read()
  return data
}

// ================================================
func (this *File) OpenWrite() {
  this.OpenWriteMode(0644)
}

func (this *File) OpenWriteMode(mode os.FileMode) {
  instance := fs.OpenWrite(this.path, mode)
  this.instance = instance
}

func (this *File) OpenRead() {
  instance := fs.OpenRead(this.path)
  this.instance = instance
}

// ================================================
func (this *File) GetFd() int {
  file := this.instance
  fd := int(file.Fd())
  return fd
}

func (this *File) GetPath() string {
  return this.path
}

func (this *File) AssertOpen() {
  if !this.IsOpen() {
    logger.LogErrorAndPanic(errors.New("file NOT open"))
  }
}

func (this *File) IsOpen() bool {
  file := this.instance
  return file != nil
}

// ================================================
func (this *File) Exists() bool {
  return fs.FileExists(this.path)
}

func (this *File) Close() {
  file := this.instance
  file.Close()
  this.instance = nil
}

func (this *File) Write(data []byte) (int, error) {
  this.AssertOpen()
  file := this.instance
  return file.Write(data)
}

func (this *File) Read() []byte {
  this.AssertOpen()
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
