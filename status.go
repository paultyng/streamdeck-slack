package main

import (
	"fmt"
	"sync"

	"github.com/slack-go/slack"
	"github.com/valyala/fastjson"

	"streamdeck-slack/internal/sdk"
)

const (
	stateInactive = 0
	stateActive   = 1
)

type statusAction struct {
	// action instance info
	action  string
	context string

	// streamdeck properties
	apiKey     string
	emoji      string
	statusText string

	client *slack.Client
	// TODO: we need to make sure not to sync status during a key event
	// also probably need to check this periodically
	init sync.Once
}

func (a *statusAction) Initialize(action, context, deviceID string) error {
	a.action = action
	a.context = context

	err := sdk.GetSettings(a.context)
	if err != nil {
		return fmt.Errorf("unable to get settings: %w", err)
	}

	return nil
}

func (a *statusAction) UpdateSettings(settings *fastjson.Object) error {
	changes := false
	settings.Visit(func(key []byte, value *fastjson.Value) {
		switch key := string(key); key {
		case "api-token":
			newAPIKey := sdk.JsonStringValue(value)
			if a.apiKey != newAPIKey {
				a.client = nil
			}
			a.apiKey = newAPIKey
		case "status-text":
			a.statusText = sdk.JsonStringValue(value)
			changes = true
		case "status-emoji":
			a.emoji = sdk.JsonStringValue(value)
			changes = true
		default:
			sdk.Log(fmt.Sprintf("unknown setting: %q", key))
		}
	})

	if a.client == nil && a.apiKey != "" {
		sdk.Log("Creating Slack client")
		a.client = slack.New(a.apiKey)
		changes = true
	}

	if !changes || a.client == nil {
		return nil
	}

	sdk.Log("Checking User profile")

	profile, err := a.client.GetUserProfile(&slack.GetUserProfileParameters{})
	if err != nil {
		return fmt.Errorf("unable to get user profile: %w", err)
	}

	sdk.Log(fmt.Sprintf("Current status emoji: %q", profile.StatusEmoji))

	a.init.Do(func() {
		// TODO: pull emoji images for team if this is a custom/alias
		// hard code all the unicode emoji in here since no public API to get them

		if profile.StatusEmoji != a.emoji {
			sdk.Log("Setting inactive state")
			err := sdk.SetState(a.context, stateInactive)
			if err != nil {
				handleError(a.context, fmt.Errorf("unable to set state: %w", err))
			}
		} else {
			sdk.Log("Setting active state")
			err := sdk.SetState(a.context, stateActive)
			if err != nil {
				handleError(a.context, fmt.Errorf("unable to set state: %w", err))
			}
		}
	})

	return nil
}

func (a *statusAction) KeyDown(payload *fastjson.Object) error {
	return nil
}
func (a *statusAction) KeyUp(payload *fastjson.Object) error {
	if a.client == nil {
		return nil
	}

	toState := 0
	if desired := payload.Get("userDesiredState"); desired != nil && desired.Type() != fastjson.TypeNull {
		toState = desired.GetInt()
	} else {
		// not a multiaction with a specific desired state, so just read current state and toggle it
		state, err := payload.Get("state").Int()
		if err != nil {
			return fmt.Errorf("unable to get state from payload: %w", err)
		}
		toState = (state + 1) % 2
	}

	sdk.Log(fmt.Sprintf("Setting state to %d", toState))

	// TODO: compare state to whatever is set in slack periodically?

	switch toState {
	case stateInactive:
		sdk.Log("Clearing status")

		err := a.client.SetUserCustomStatus("", "", 0)
		if err != nil {
			return fmt.Errorf("unable to clear status: %w", err)
		}
	case stateActive:
		sdk.Log(fmt.Sprintf("Setting status %q %q", a.emoji, a.statusText))

		err := a.client.SetUserCustomStatus(a.statusText, a.emoji, 0)
		if err != nil {
			return fmt.Errorf("unable to set status: %w", err)
		}
	default:
		return fmt.Errorf("unknown state: %d", toState)
	}

	return nil
}
func (a *statusAction) DidReceiveSettings(payload *fastjson.Object) error {
	return nil
}
func (a *statusAction) SendToPlugin(payload *fastjson.Object) error {
	err := a.UpdateSettings(payload)
	if err != nil {
		return fmt.Errorf("unable to update settings in send to plugin: %w", err)
	}

	err = sdk.SetSettings(a.context, map[string]string{
		"api-token":    a.apiKey,
		"status-text":  a.statusText,
		"status-emoji": a.emoji,
	})
	if err != nil {
		return fmt.Errorf("unable to set settings: %w", err)
	}
	return nil
}
func (a *statusAction) WillAppear(*fastjson.Object) error {
	return nil
}
func (a *statusAction) WillDisappear(*fastjson.Object) error {
	return nil
}
