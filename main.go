package main

import (
	"fmt"
	"log"
	"os"

	"github.com/valyala/fastjson"

	"streamdeck-slack/internal/sdk"
)

var debugLog = os.Getenv("STREAMDECK_SLACK_DEBUG") != ""

func main() {
	if debugLog {
		f, err := os.OpenFile("../../logs/com.paultyng.slack.go.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			log.Fatalf("error opening file: %v", err)
		}
		defer f.Close()
		log.SetOutput(f)
	}

	sdk.Log("Starting plugin 1")
	defer func() {
		if debugLog {
			log.Printf("Finalizing plugin")
		}
		sdk.Log("Finalizing plugin")
		if err := recover(); err != nil {
			if debugLog {
				log.Printf("Recovering from panic: %s", err)
			}
			sdk.Log(fmt.Sprintf("Panic: %v", err))
		}
	}()

	if err := run(os.Args[1:]); err != nil {
		if debugLog {
			log.Printf("error from run: %s", err)
		}
		sdk.Log(fmt.Sprintf("error from run: %s", err))
		os.Exit(3)
	}
}

type streamDeckAction interface {
	Initialize(action, context, deviceID string) error
	UpdateSettings(*fastjson.Object) error

	// Events
	KeyDown(*fastjson.Object) error
	KeyUp(*fastjson.Object) error
	DidReceiveSettings(*fastjson.Object) error
	SendToPlugin(*fastjson.Object) error
	WillAppear(*fastjson.Object) error
	WillDisappear(*fastjson.Object) error
}

var (
	factories = map[string]func() streamDeckAction{
		"com.paultyng.slack.status": func() streamDeckAction {
			return &statusAction{}
		},

		"com.paultyng.slack.presence": func() streamDeckAction {
			return &presenceAction{}
		},

		"com.paultyng.slack.user": func() streamDeckAction {
			return &userAction{}
		},
	}
	actions = map[string]streamDeckAction{}
)

func actionFor(action, context, deviceID string) streamDeckAction {
	if a, ok := actions[context]; ok {
		return a
	}

	if f, ok := factories[action]; ok {
		a := f()
		err := a.Initialize(action, context, deviceID)
		if err != nil {
			handleError(context, err)
			return nil
		}
		actions[context] = a
		return a
	}

	return nil
}

func run(args []string) error {
	// event handlers
	sdk.AddHandler(func(e *sdk.ApplicationLaunchEvent) {
		logHandler("ApplicationLaunchEvent", e.Application, "", "", nil)
	})
	sdk.AddHandler(func(e *sdk.ApplicationTerminateEvent) {
		logHandler("ApplicationTerminateEvent", e.Application, "", "", nil)
	})
	sdk.AddHandler(func(e *sdk.DeviceConnectEvent) {
		logHandler("DeviceConnectEvent", "", "", e.DeviceId, e.Info)
	})
	sdk.AddHandler(func(e *sdk.DeviceDisconnectEvent) {
		logHandler("DeviceDisconnectEvent", "", "", e.DeviceId, nil)
	})
	sdk.AddHandler(func(e *sdk.GlobalSettingsEvent) {
		logHandler("GlobalSettingsEvent", "", "", "", e.Settings)
	})
	sdk.AddHandler(func(e *sdk.KeyDownEvent) {
		logHandler("KeyDownEvent", e.Action, e.Context, e.DeviceId, e.Payload)

		action := actionFor(e.Action, e.Context, e.DeviceId)
		if action == nil {
			return
		}

		data, err := e.Payload.Object()
		if err != nil {
			handleError(e.Context, err)
		}

		settings, err := data.Get("settings").Object()
		if err != nil {
			handleError(e.Context, err)
		}
		err = action.UpdateSettings(settings)
		if err != nil {
			handleError(e.Context, err)
		}

		err = action.KeyDown(data)
		if err != nil {
			handleError(e.Context, err)
		}
	})
	sdk.AddHandler(func(e *sdk.KeyUpEvent) {
		logHandler("KeyUpEvent", e.Action, e.Context, e.DeviceId, e.Payload)

		action := actionFor(e.Action, e.Context, e.DeviceId)
		if action == nil {
			return
		}

		data, err := e.Payload.Object()
		if err != nil {
			handleError(e.Context, err)
		}

		settings, err := data.Get("settings").Object()
		if err != nil {
			handleError(e.Context, err)
		}
		err = action.UpdateSettings(settings)
		if err != nil {
			handleError(e.Context, err)
		}

		err = action.KeyUp(data)
		if err != nil {
			handleError(e.Context, err)
		}
	})
	sdk.AddHandler(func(e *sdk.ReceiveSettingsEvent) {
		logHandler("ReceiveSettingsEvent", e.Action, e.Context, e.DeviceId, e.Payload)

		action := actionFor(e.Action, e.Context, e.DeviceId)
		if action == nil {
			return
		}

		data, err := e.Payload.Object()
		if err != nil {
			handleError(e.Context, err)
		}

		settings, err := data.Get("settings").Object()
		if err != nil {
			handleError(e.Context, err)
		}
		err = action.UpdateSettings(settings)
		if err != nil {
			handleError(e.Context, err)
		}

		err = action.DidReceiveSettings(data)
		if err != nil {
			handleError(e.Context, err)
		}
	})
	sdk.AddHandler(func(e *sdk.SendToPluginEvent) {
		logHandler("SendToPluginEvent", e.Action, e.Context, e.DeviceId, e.Payload)

		action := actionFor(e.Action, e.Context, e.DeviceId)
		if action == nil {
			return
		}

		data, err := e.Payload.Object()
		if err != nil {
			handleError(e.Context, err)
		}

		err = action.SendToPlugin(data)
		if err != nil {
			handleError(e.Context, err)
		}
	})
	sdk.AddHandler(func(e *sdk.WillAppearEvent) {
		logHandler("WillAppearEvent", e.Action, e.Context, e.DeviceId, e.Payload)

		action := actionFor(e.Action, e.Context, e.DeviceId)
		if action == nil {
			return
		}

		data, err := e.Payload.Object()
		if err != nil {
			handleError(e.Context, err)
		}

		settings, err := data.Get("settings").Object()
		if err != nil {
			handleError(e.Context, err)
		}
		err = action.UpdateSettings(settings)
		if err != nil {
			handleError(e.Context, err)
		}

		err = action.WillAppear(data)
		if err != nil {
			handleError(e.Context, err)
		}
	})
	sdk.AddHandler(func(e *sdk.WillDisappearEvent) {
		logHandler("WillDisappearEvent", e.Action, e.Context, e.DeviceId, e.Payload)

		action := actionFor(e.Action, e.Context, e.DeviceId)
		if action == nil {
			return
		}

		data, err := e.Payload.Object()
		if err != nil {
			handleError(e.Context, err)
		}

		settings, err := data.Get("settings").Object()
		if err != nil {
			handleError(e.Context, err)
		}
		err = action.UpdateSettings(settings)
		if err != nil {
			handleError(e.Context, err)
		}

		err = action.WillDisappear(data)
		if err != nil {
			handleError(e.Context, err)
		}
	})

	sdk.Log("Opening Plugin socket")
	err := sdk.Open()
	if err != nil {
		return fmt.Errorf("unable to open SDK: %w", err)
	}

	sdk.Log("Waiting for program to exit")
	sdk.Wait()

	return nil
}

func logHandler(name, action, context, deviceID string, payload *fastjson.Value) {
	sdk.Log(fmt.Sprintf("%s: %q %q %q", name, action, context, deviceID))
	if payload != nil {
		data := payload.MarshalTo(nil)
		sdk.Log(fmt.Sprintf("%s Data: %s", name, string(data)))
	}
}

func handleError(context string, err error) {
	if context != "" {
		// ignore any error here
		_ = sdk.ShowAlert(context)
	}
	if debugLog {
		// This may double log, but so be it
		log.Printf("Error: %s", err)
	}
	sdk.Log(fmt.Sprintf("Error: %s", err))
}
