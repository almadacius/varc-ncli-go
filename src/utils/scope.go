package utils

import (
  "os"
  "fmt"
  "regexp"
  "encoding/json"
  "io/ioutil"
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

func ListScopes() []string {
  reExt := regexp.MustCompile(`\..{3,4}$`)

  scopeDir := GetStorageDir()

  files, err := ioutil.ReadDir(scopeDir)
  LogErrorAndPanic(err)

  filenames := []string{}
  for _, f := range files {
    name := reExt.ReplaceAllString(f.Name(), "")
    filenames = append(filenames, name)
  }

  return filenames
}

// ================================================
func GetScopeFile(name string) string {
  dir := GetStorageDir()
  path := dir + "/" + name + ".json"
  return path
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

  path := GetScopeFile(scope.name)
  scope.path = path

  scope.data = make(map[string] string)

  scope.tryLoad()

  return scope
}

func (s *Scope) SetVar(key string, value string) {
  s.data[key] = value
  s.save()
}

func (s *Scope) UnsetVar(key string) {
  delete(s.data, key)
  s.save()
}

func (s *Scope) GetData() map[string] string {
  return s.data
}

func (s *Scope) GetVar(key string) string {
  return s.data[key]
}

func (s *Scope) GetKeys() []string {
  keys := []string{}

  for key, _ := range s.data {
    keys = append(keys, key)
  }

  return keys
}

func (s *Scope) Delete() {
  filepath := s.path

  if FileExists(filepath) {
    os.Remove(filepath)
  }
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
