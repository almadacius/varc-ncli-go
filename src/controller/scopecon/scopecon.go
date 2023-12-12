package scopecon

import (
  "os"
  "fmt"
  "regexp"
  "io/ioutil"
  "almadash/varc/utils/fs"
  "almadash/varc/utils/logger"
  "almadash/varc/utils/jsonfile"
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

  scopeDir := fs.GetStorageDir()

  files, err := ioutil.ReadDir(scopeDir)
  logger.LogErrorAndPanic(err)

  filenames := []string{}
  for _, f := range files {
    name := reExt.ReplaceAllString(f.Name(), "")
    filenames = append(filenames, name)
  }

  return filenames
}

// ================================================
func GetScopeFile(name string) string {
  dir := fs.GetStorageDir()
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

/*
  - use right after NewScope, as it has fresh loaded from disk
*/
func (s *Scope) SetVarIfChanged(key string, value string) {
  currVal := s.GetVar(key)
  // value didn't change, do nothing
  if currVal == value { return }
  s.SetVar(key, value)
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

  if fs.FileExists(filepath) {
    os.Remove(filepath)
  }
}

// ================================================
func (s *Scope) save() {
  jsonfile.Save(s.path, s.data)
}

func (s *Scope) load() {
  data := jsonfile.Load(s.path)
  s.data = data
}

func (s *Scope) tryLoad() {
  if fs.FileExists(s.path) {
    s.load()
  }
}
