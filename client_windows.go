package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

var (
	cmd      = "url.dll,FileProtocolHandler"
	runDll32 = filepath.Join(os.Getenv("SYSTEMROOT"), "System32", "rundll32.exe")
)

func openDM(team, user string) error {
	cmd := exec.Command(runDll32, cmd, fmt.Sprintf("slack://user?team=%s&id=%s", team, user))
	return cmd.Start()
}
