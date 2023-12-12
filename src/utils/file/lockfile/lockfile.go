package lockfile

import (
  "almadash/varc/utils/file/jsonfile"
  "almadash/varc/utils/slice"
)

type JsonFile = jsonfile.JsonFile

// ================================================
type LockFile struct {
  JsonFile
  pids []int
}

func New(path string) LockFile {
  out := LockFile{}
  out.JsonFile = jsonfile.New(path)
  return out
}

func (this *LockFile) Reset() {
  data := jsonfile.EmptyMap()
  data["pids"] = []int{}
  this.Data = data
  this.Save()
}

func (this *LockFile) Load() {
  if !this.Exists() {
    this.Reset()
  }
  this.JsonFile.Load()
  pids := this.JsonFile.GetIntArray("pids")
  this.updateData()
  this.pids = pids
}

func (this *LockFile) updateData() {
  this.JsonFile.Data["pids"] = this.pids
}

func (this *LockFile) GetList() []int {
  return this.pids
}

func (this *LockFile) AddPid(pid int) {
  pids := this.GetList()
  this.pids = append(pids, pid)
  this.updateData()
  this.Save()
}

func (this *LockFile) RemovePid(pid int) {
  pids := this.GetList()
  this.pids = slice.RemoveFirst(pids, pid)
  this.updateData()
  this.Save()
}
