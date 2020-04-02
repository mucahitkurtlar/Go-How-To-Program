package main

import (
  "fmt"
  "os"
  "os/exec"
  "bufio"
)

func main() {
  cmd := exec.Command("go", "version")
  pipe, err := cmd.StdoutPipe()
  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
  out := bufio.NewScanner(pipe)
  go func() {
    for out.Scan() {
      fmt.Println(out.Text())
    }
  }()
  err = cmd.Start()
  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
  err = cmd.Wait()
  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
}
