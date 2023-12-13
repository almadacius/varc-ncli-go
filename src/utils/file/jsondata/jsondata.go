package jsondata

import (
  "encoding/json"
  "almadash/varc/utils/logger"
  "almadash/varc/utils/obj"
)

// ================================================
type JsonType = map[string] interface{}

// ================================================
/*
  - manages a JSON structure
  - marshall/unmarshall
  - isDirty
*/
type JsonData struct {
  data JsonType
  dirty bool
}

func New() JsonData {
  out := JsonData{}
  out.dirty = false
  out.ResetData()

  return out
}

// ================================================
func (this *JsonData) ResetData() {
  this.data = obj.EmptyMap()
}

func (this *JsonData) Set(key string, value interface{}) {
  this.data[key] = value
}

func (this *JsonData) Get(key string) interface{} {
  return this.data[key]
}

// ================================================
func (this *JsonData) IsDirty() bool {
  return this.dirty
}

func (this *JsonData) ClearDirty() {
  this.dirty = false
}

// ================================================
func (this *JsonData) SetBytes(rawData []byte) {
  var data JsonType
  err := json.Unmarshal(rawData, &data)
  logger.LogErrorAndPanic(err)
  this.data = data
}

func (this *JsonData) ToBytes() []byte {
  rawData, err := json.Marshal(this.data)
  logger.LogErrorAndPanic(err)
  return rawData
}

// ================================================
func (this *JsonData) GetIntArray(key string) []int {
  listItem := this.Get(key)
  raw := listItem.([]interface{})
  out := []int{}
  for _, v := range raw {
    num := int(v.(float64))
    out = append(out, num)
  }
  return out
}
