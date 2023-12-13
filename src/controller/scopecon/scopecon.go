package scopecon

import (
  "fmt"
  "regexp"
  "almadash/varc/utils/fs"
  "almadash/varc/utils/file/jsondata"
  "almadash/varc/utils/file/file"
)

// ================================================
type JsonData = jsondata.JsonData
type File = file.File

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

  files := fs.ReadDir(scopeDir)

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
  data JsonData
  file File
}

func NewScope(name string) Scope {
  out := Scope{}
  out.name = name

  path := GetScopeFile(out.name)
  out.path = path

  out.file = file.New(path)
  out.data = jsondata.New()

  out.tryLoad()

  return out
}

// ================================================
func (this *Scope) GetData() JsonData {
  return this.data
}

func (this *Scope) GetVar(key string) string {
  rawVal := this.data.Get(key)
  return rawVal.(string)
}

func (this *Scope) GetKeys() []string {
  keys := this.data.GetKeys()
  return keys
}

func (this *Scope) SetVar(key string, value string) {
  this.data.Set(key, value)
  this.save()
}

func (this *Scope) UnsetVar(key string) {
  this.data.Unset(key)
  this.save()
}

func (this *Scope) Delete() {
  this.file.Delete()
}

// ================================================
func (this *Scope) save() {
  // data did NOT change, do NOT waste disk IO
  if(!this.data.IsDirty()) {
    return
  }
  rawData := this.data.ToBytes()
  this.file.Write(rawData)
  this.data.ClearDirty()
}

func (this *Scope) load() {
  rawData := this.file.Read()
  this.data.SetBytes(rawData)
}

func (this *Scope) tryLoad() {
  if this.file.Exists() {
    this.load()
  }
}
