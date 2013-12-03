package main

import (
  "fmt"
  "strings"
  "os/exec"
  "syscall"
)

type CheckConfig struct {
  handler     string
  command     string
  interval    int
  standalone  bool
}

type CheckResult struct {
  handler []string
  command string
  interval int
  subscribers []string
  standalone bool
  name string
  issued int64
  executed int64
  output string
  status int
  duration float32
}

type ResultMsg struct {
  client string
  check CheckResult
}

func buildConfig() CheckConfig {
  check := CheckConfig{
    handler: "default",
    command: "/checks/random.sh 150 199",
    interval: 10,
    standalone: true,
  }

  return check
}

func runCheck(check CheckConfig) ResultMsg {
  pathAndArgs := strings.Fields(check.command)

  cmd := exec.Command(pathAndArgs[0], pathAndArgs[1:]...)

  output, _ := cmd.Output()
  exitCode := cmd.ProcessState.Sys().(syscall.WaitStatus).ExitStatus()

  fmt.Println(string(output))
  fmt.Println(exitCode)


  return ResultMsg{}
}

func main() {
  c := buildConfig()
  runCheck(c)
}
