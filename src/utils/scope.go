package utils

import (
  // "os"
  // "path"
  "fmt"
  "regexp"
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
}

func NewScope(name string) Scope {
  scope := Scope{}
  scope.name = name

  dir := GetStorageDir()
  path := dir + "/" + scope.name + ".json"
  scope.path = path

  return scope
}
