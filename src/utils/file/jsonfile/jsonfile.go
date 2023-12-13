package jsonfile

import (
  "almadash/varc/utils/file/file"
  "almadash/varc/utils/file/jsondata"
)

type File = file.File
type JsonData = jsondata.JsonData

// ================================================
type JsonFile struct {
  Data JsonData
  File
}

func New(path string) JsonFile {
  out := JsonFile{}
  out.File = file.New(path)
  out.Data = jsondata.New()
  return out
}

func (this *JsonFile) Save() {
  jsonBytes := this.Data.ToBytes()

  this.OpenWrite()
  defer this.Close()
  this.Write(jsonBytes)
}

func (this *JsonFile) Reset() {
  this.Data.ResetData()
  this.Save()
}

func (this *JsonFile) Load() JsonData {
  this.OpenRead()
  defer this.Close()

  jsonBytes := this.Read()

  this.Data.SetBytes(jsonBytes)
  return this.Data
}
