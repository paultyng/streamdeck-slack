package main

import (
	"fmt"
	"os/exec"
)

func openDM(team, user string) error {
	return exec.Command("open", fmt.Sprintf("slack://user?team=%s&id=%s", team, user)).Start()
}
