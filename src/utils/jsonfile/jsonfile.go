package jsonfile

import (
  "encoding/json"
  "almadash/varc/utils/logger"
)

// ================================================
func Save(path string, data map[string]string) {
  jsonBytes, err := json.Marshal(data)
  logger.LogErrorAndPanic(err)

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
  logger.LogError(err)
  // file is broken, reset it for now
  if err != nil {
    logger.LogInfo("error reading data, reset file")
    data = Reset(path)
  }

  return data
}

// ================================================
func emptyMap() map[string] string {
  out := make(map[string] string)
  return out
}
