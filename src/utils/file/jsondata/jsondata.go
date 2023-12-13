package jsondata

import (
  "encoding/json"
  "almadash/varc/utils/logger"
  "almadash/varc/utils/obj"
  "almadash/varc/utils/reflect"
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
  if(reflect.IsComparable(value)) {
    currVal := this.Get(key)
    // value did NOT change, do nothing
    if currVal == value {
      return
    }
  }

  this.data[key] = value
  this.dirty = true
}

func (this *JsonData) Unset(key string) {
  delete(this.data, key)
  this.dirty = true
}

func (this *JsonData) Get(key string) interface{} {
  return this.data[key]
}

func (this *JsonData) GetIfAvailable(key string) (interface{}, bool) {
  value, ok := this.data[key]
  return value, ok
}

func (this *JsonData) HasKey(key string) bool {
  data := this.data
  _, ok := data[key]
  return ok
}

// ================================================
func (this *JsonData) GetString(key string) string {
  return this.Get(key).(string)
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
func (this *JsonData) GetKeys() []string {
  data := this.data

  keys := []string{}

  for key, _ := range data {
    keys = append(keys, key)
  }

  return keys
}

func (this *JsonData) GetData() JsonType {
  return this.data
}

// ================================================
func (this *JsonData) GetIntArray(key string) []int {
  out := []int{}

  if !this.HasKey(key) {
    return out
  }

  listItem := this.Get(key)
  raw := listItem.([]interface{})

  for _, v := range raw {
    num := int(v.(float64))
    out = append(out, num)
  }
  return out
}

// ================================================
func (this *JsonData) ForEachString(fn func(string, string)) {
  keys := this.GetKeys()
  for _, key := range keys {
    value := this.GetString(key)
    fn(key, value)
  }
}
