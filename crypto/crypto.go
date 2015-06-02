package crypto

import (
  "runtime"
  "os"
  "strings"
)

func Decrypt(str string) string {
  return str
}

func Encrypt(str string) string {
  return str
}

func getKey() string {
  host, _ := os.Hostname()
  cpuCount := string(runtime.NumCPU())
  keyParts := []string{cpuCount, runtime.GOARCH, runtime.GOOS, host}

  return strings.Join(keyParts, "-")
}
