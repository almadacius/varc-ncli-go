package jsonfile

import (
  "os"
  "encoding/json"
  "almadash/varc/utils"
)

// ================================================
func Save(path string, data map[string]string) {
  jsonBytes, err := json.Marshal(data)
  utils.LogErrorAndPanic(err)

  writeFileWithLock(path, jsonBytes, 0644)
}

func Reset(path string) map[string] string {
  emptyData := emptyMap()
  Save(path, emptyData)
  return emptyData
}

func Load(path string) map[string]string {
  jsonBytes := readFileWithLock(path)

  var data map[string] string

  err := json.Unmarshal(jsonBytes, &data)
  utils.LogError(err)
  // file is broken, reset it for now
  if err != nil {
    data = Reset(path)
  }

  return data
}

// ================================================
func emptyMap() map[string] string {
  out := make(map[string] string)
  return out
}

func writeFileWithLock(path string, data []byte, mode os.FileMode) {
  file := CreateFile(path)
  file.OpenWrite(mode)
  defer file.Close()
  file.Lock()
  defer file.Unlock()

  _, err := file.Write(data)
  utils.LogErrorAndPanic(err)
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
  utils.LogErrorAndPanic(err)

  return data
}
