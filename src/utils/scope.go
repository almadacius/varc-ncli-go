package utils

import (
  "os"
  "fmt"
  "regexp"
  "encoding/json"
)

// ================================================
func ParseScope(scopePath string) (string, string) {
  reScope := regexp.MustCompile(`^([0-9a-zA-Z_-]+)\.(\w+)$`)

  bytes := []byte(scopePath)
  m := reScope.Match(bytes)

  if ! m {
    fmt.Println("not a match")
  }

  subs := reScope.FindStringSubmatch(scopePath)
  scopeName := subs[1]
  varName := subs[2]

  return scopeName, varName
}

// ================================================
type Scope struct {
  name string
  path string
  data map[string] string
}

func NewScope(name string) Scope {
  scope := Scope{}
  scope.name = name

  dir := GetStorageDir()
  path := dir + "/" + scope.name + ".json"
  scope.path = path

  scope.data = make(map[string] string)

  scope.tryLoad()

  return scope
}

func (s *Scope) SetVar(key string, value string) {
  s.data[key] = value
  s.save()
}

func (s *Scope) GetData() map[string] string {
  return s.data
}

func (s *Scope) GetVar(key string) string {
  return s.data[key]
}

// ================================================
func (s *Scope) save() {
  jsonBytes, err := json.Marshal(s.data)
  LogError(err)

  err = os.WriteFile(s.path, jsonBytes, 0644)
  LogErrorAndPanic(err)
}

func (s *Scope) load() {
  jsonBytes, err := os.ReadFile(s.path)
  LogErrorAndPanic(err)

  var data map[string] string

  err = json.Unmarshal(jsonBytes, &data)
  LogErrorAndPanic(err)

  s.data = data
}

func (s *Scope) tryLoad() {
  if FileExists(s.path) {
    s.load()
  }
}
