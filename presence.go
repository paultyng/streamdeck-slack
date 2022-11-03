package main

import (
	"fmt"
	"streamdeck-slack/internal/sdk"
	"sync"

	"github.com/slack-go/slack"
	"github.com/valyala/fastjson"
)

type presenceAction struct {
	// action instance info
	action  string
	context string

	// streamdeck properties
	apiKey string

	client *slack.Client
	// TODO: we need to make sure not to sync status during a key event
	// also probably need to check this periodically
	init sync.Once
	user string
}

func (a *presenceAction) Initialize(action, context, deviceID string) error {
	a.action = action
	a.context = context

	err := sdk.GetSettings(a.context)
	if err != nil {
		return fmt.Errorf("unable to get settings: %w", err)
	}

	return nil
}

func (a *presenceAction) UpdateSettings(settings *fastjson.Object) error {
	changes := false
	settings.Visit(func(key []byte, value *fastjson.Value) {
		switch key := string(key); key {
		case "api-token":
			newAPIKey := sdk.JsonStringValue(value)
			if a.apiKey != newAPIKey {
				a.client = nil
				a.user = ""
			}
			a.apiKey = newAPIKey
		default:
			sdk.Log(fmt.Sprintf("unknown setting: %q", key))
		}
	})

	if a.client == nil && a.apiKey != "" {
		sdk.Log("Creating Slack client")
		client := slack.New(a.apiKey)
		auth, err := client.AuthTest()
		if err != nil {
			return fmt.Errorf("unable to auth test: %w", err)
		}
		a.client = client
		a.user = auth.UserID
		changes = true
	}

	if !changes || a.client == nil {
		return nil
	}

	sdk.Log("Checking presence state")

	presence, err := a.client.GetUserPresence(a.user)
	if err != nil {
		return fmt.Errorf("unable to get user presence %q: %w", a.user, err)
	}

	sdk.Log(fmt.Sprintf("Current presence state: %#v", presence))

	a.init.Do(func() {
		if !presence.ManualAway {
			sdk.Log("Found auto away")
			err := sdk.SetState(a.context, stateInactive)
			if err != nil {
				handleError(a.context, fmt.Errorf("unable to set state: %w", err))
			}
		} else {
			sdk.Log("Found manual away")
			err := sdk.SetState(a.context, stateActive)
			if err != nil {
				handleError(a.context, fmt.Errorf("unable to set state: %w", err))
			}
		}
	})

	return nil
}

func (a *presenceAction) KeyDown(payload *fastjson.Object) error {
	return nil
}
func (a *presenceAction) KeyUp(payload *fastjson.Object) error {
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
		sdk.Log("Setting presence to auto away")

		err := a.client.SetUserPresence("auto")
		if err != nil {
			return fmt.Errorf("unable to set presence to auto: %w", err)
		}
	case stateActive:
		sdk.Log("Setting presence to manual away")

		err := a.client.SetUserPresence("away")
		if err != nil {
			return fmt.Errorf("unable to set presence to away: %w", err)
		}
	default:
		return fmt.Errorf("unknown state: %d", toState)
	}

	return nil
}
func (a *presenceAction) DidReceiveSettings(payload *fastjson.Object) error {
	return nil
}
func (a *presenceAction) SendToPlugin(payload *fastjson.Object) error {
	err := a.UpdateSettings(payload)
	if err != nil {
		return fmt.Errorf("unable to update settings in send to plugin: %w", err)
	}

	err = sdk.SetSettings(a.context, map[string]string{
		"api-token": a.apiKey,
	})
	if err != nil {
		return fmt.Errorf("unable to set settings: %w", err)
	}
	return nil
}
func (a *presenceAction) WillAppear(*fastjson.Object) error {
	return nil
}
func (a *presenceAction) WillDisappear(*fastjson.Object) error {
	return nil
}
