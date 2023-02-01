package jsonfile

import (
  "os"
  "encoding/json"
  "almadash/varc/utils"
)

// ================================================
func Save(path string, data map[string]string) {
  jsonBytes, err := json.Marshal(data)
  utils.LogError(err)

  err = os.WriteFile(path, jsonBytes, 0644)
  utils.LogErrorAndPanic(err)
}

func Load(path string) map[string]string {
  jsonBytes, err := os.ReadFile(path)
  utils.LogErrorAndPanic(err)

  var data map[string] string

  err = json.Unmarshal(jsonBytes, &data)
  utils.LogErrorAndPanic(err)

  return data
}
