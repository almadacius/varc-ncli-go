package logger

import (
  "os"
  "fmt"
  "log"
  "strings"
)

// ================================================
func LogInfo(message string) {
  fmt.Println("[info]: ", message)
}

func LogError(err error) {
  if err != nil {
    fmt.Fprintf(os.Stderr, "[error]: %s\n", err)
  }
}

func LogWarn(msg string) {
  fmt.Fprintf(os.Stderr, "[warn]: %s\n", msg)  
}

// ================================================
func LogErrorAndPanic(err error) {
  if err != nil {
    LogError(err)
    log.Fatal(err)
    panic(err)
  }
}

// ================================================
func StringListToBytes(strs []string) []byte {
  joined := strings.Join(strs, "\n")
  bytes := []byte(joined)
  return bytes
}

// ================================================
func Output(data string) {
  bytes := []byte(data)
  os.Stdout.Write(bytes)
}
