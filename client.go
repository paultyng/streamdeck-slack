//go:build !windows && !darwin

package main

import (
	"fmt"

	"streamdeck-slack/sdk"
)

func openDM(team, user string) error {
	return sdk.OpenURL(fmt.Sprintf("https://slack.com/app_redirect?channel=%s", user))
}
