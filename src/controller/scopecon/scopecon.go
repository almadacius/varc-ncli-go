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
  data := &this.data
  if !data.HasKey(key) {
    return ""
  }
  return data.GetString(key)
}

func (this *Scope) GetKeys() []string {
  keys := this.data.GetKeys()
  return keys
}

func (this *Scope) SetVar(key string, value string) {
  data := &this.data
  data.Set(key, value)
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
  file := &this.file
  data := &this.data

  // data did NOT change, do NOT waste disk IO
  if(!data.IsDirty()) {
    return
  }

  rawData := data.ToBytes()
  file.OpenWrite()
  defer file.Close()
  file.Write(rawData)
  data.ClearDirty()
}

func (this *Scope) load() {
  file := &this.file
  data := &this.data

  file.OpenRead()
  defer file.Close()
  rawData := file.Read()
  data.SetBytes(rawData)
}

func (this *Scope) tryLoad() {
  file := this.file
  if file.Exists() {
    this.load()
  }
}
