package jsonfile

import (
  "encoding/json"
  "almadash/varc/utils/logger"
  "almadash/varc/utils/file/file"
)

type File = file.File
type JsonData map[string] interface{}

// ================================================
type JsonFile struct {
  Data JsonData
  File
}

func New(path string) JsonFile {
  jsonFile := JsonFile{}
  jsonFile.File = file.New(path)
  return jsonFile
}

func (this *JsonFile) Save() {
  jsonBytes, err := json.Marshal(this.Data)
  logger.LogErrorAndPanic(err)

  this.OpenWrite(0644)
  defer this.Close()
  this.Write(jsonBytes)
}

func (this *JsonFile) Reset() {
  this.Data = EmptyMap()
  this.Save()
}

func (this *JsonFile) Load() JsonData {
  this.OpenRead()
  defer this.Close()

  jsonBytes := this.Read()

  var data JsonData

  err := json.Unmarshal(jsonBytes, &data)
  logger.LogErrorAndPanic(err)

  this.Data = data
  return data
}

// ================================================
func EmptyMap() JsonData {
  out := make(JsonData)
  return out
}

// ================================================
func (this *JsonFile) GetIntArray(key string) []int {
  data := this.Data
  raw := data[key].([]interface{})
  out := []int{}
  for _, v := range raw {
    num := int(v.(float64))
    out = append(out, num)
  }
  return out
}
