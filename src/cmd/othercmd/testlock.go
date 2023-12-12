package othercmd

import (
  "fmt"
  "os"
  "almadash/varc/utils/fs"
  "almadash/varc/utils/file/lockfile"
)

// ================================================
type TestLockCmd struct {}

func (c *TestLockCmd) Declare() { /* virtual */ }

func (c *TestLockCmd) Run(options []string) {
  path := fs.GetTestLockFile()
  file := lockfile.New(path)
  file.Load()
  pid := os.Getpid()

  file.AddPid(pid)
  pids := file.GetList()
  fmt.Println("added: ", pids)

  file.RemovePid(pid)
  pids = file.GetList()
  fmt.Println("removed: ", pids)
}
