package main

import (
	"errors"

	"github.com/slack-go/slack"
)

func isMissingScopeError(err error) bool {
	var slackError *slack.SlackErrorResponse
	return errors.As(err, &slackError) && slackError.Err == "missing_scope"
}
