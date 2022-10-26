package main

// TODO: user action, shows current status if they have one, otherwise button to DM (start url?)

// online / offline, current status, etc.

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	"image/draw"
	"image/png"
	"regexp"
	"time"

	"github.com/slack-go/slack"
	"github.com/valyala/fastjson"
	xdraw "golang.org/x/image/draw"

	"streamdeck-slack/sdk"
)

var emojiColonSyntaxRe = regexp.MustCompile(`^:([^:]+):(:([^:]+):)?$`)

func emojiColonSyntaxToName(s string) string {
	m := emojiColonSyntaxRe.FindStringSubmatch(s)
	if m == nil || len(m) < 2 {
		return ""
	}
	return m[1]
}

type userAction struct {
	// action instance info
	action  string
	context string

	// streamdeck properties
	apiKey string
	user   string

	client *slack.Client
	ticker *time.Ticker
	close  chan bool
	team   string
}

func (a *userAction) Initialize(action, context, deviceID string) error {
	a.action = action
	a.context = context

	sdk.GetSettings(a.context)

	return nil
}

func (a *userAction) updateImage() error {
	if a.client == nil || a.user == "" {
		return nil
	}

	sdk.Log(fmt.Sprintf("getting user %q profile", a.user))

	user, err := a.client.GetUserProfile(&slack.GetUserProfileParameters{
		UserID: a.user,
	})
	if err != nil {
		return fmt.Errorf("unable to get user profile for %q: %w", a.user, err)
	}

	emojiURL := ""
	emojiName := emojiColonSyntaxToName(user.StatusEmoji)
	if emojiName != "" {
		for _, ei := range user.StatusEmojiDisplayInfo {
			if ei.EmojiName != emojiName {
				continue
			}
			emojiURL = ei.DisplayURL
		}
	}

	presence, err := a.client.GetUserPresence(a.user)
	if err != nil {
		return fmt.Errorf("unable to get user presence: %w", err)
	}

	err = a.setProfileImageWithPresence(user.Image192, emojiURL, presence.Presence == "active")
	if err != nil {
		return fmt.Errorf("unable to set image: %w", err)
	}

	return nil
}

func (a *userAction) UpdateSettings(settings *fastjson.Object) error {
	changes := false
	settings.Visit(func(key []byte, value *fastjson.Value) {
		switch key := string(key); key {
		case "api-token":
			newAPIKey := sdk.JsonStringValue(value)
			if a.apiKey != newAPIKey {
				a.client = nil
				a.team = ""

				if a.ticker != nil {
					a.ticker.Stop()
					a.ticker = nil
				}
				if a.close != nil {
					a.close <- true
					close(a.close)
					a.close = nil
				}
			}
			a.apiKey = newAPIKey
		case "user":
			newUser := sdk.JsonStringValue(value)
			if a.user != newUser {
				a.user = newUser
				changes = true
			}
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
		a.team = auth.TeamID
		changes = true

		// start ticker for updates
		a.ticker = time.NewTicker(60 * time.Second)
		go func() {
			for {
				select {
				case <-a.close:
					return
				case <-a.ticker.C:
					err := a.updateImage()
					if err != nil {
						handleError(a.context, err)
					}
				}
			}
		}()
	}

	if !changes || a.client == nil || a.user == "" {
		return nil
	}

	return a.updateImage()
}

func (a *userAction) KeyDown(payload *fastjson.Object) error {
	return nil
}
func (a *userAction) KeyUp(payload *fastjson.Object) error {
	if a.client == nil {
		return nil
	}

	return openDM(a.team, a.user)
}
func (a *userAction) DidReceiveSettings(payload *fastjson.Object) error {
	return nil
}
func (a *userAction) SendToPlugin(payload *fastjson.Object) error {
	err := a.UpdateSettings(payload)
	if err != nil {
		return fmt.Errorf("unable to update settings in send to plugin: %w", err)
	}

	sdk.SetSettings(a.context, map[string]string{
		"api-token": a.apiKey,
		"user":      a.user,
	})
	return nil
}
func (a *userAction) WillAppear(*fastjson.Object) error {
	return nil
}
func (a *userAction) WillDisappear(*fastjson.Object) error {
	return nil
}

func (a *userAction) setProfileImageWithPresence(avatarURL, statusEmojiURL string, active bool) error {
	profileImage, err := imageFromURL(avatarURL)
	if err != nil {
		return fmt.Errorf("unable to download avatar image: %w", err)
	}

	compositeImage := image.NewRGBA(profileImage.Bounds())
	draw.Draw(compositeImage, profileImage.Bounds(), profileImage, profileImage.Bounds().Min, draw.Src)

	presenceImage := activeImage
	if !active {
		presenceImage = awayImage
	}

	presenceDstR := compositeImage.Bounds()
	presenceDstR.Min = presenceDstR.Max.Div(2)

	xdraw.BiLinear.Scale(compositeImage, presenceDstR, presenceImage, presenceImage.Bounds(), xdraw.Over, nil)

	if statusEmojiURL != "" {
		statusImage, err := imageFromURL(statusEmojiURL)
		if err != nil {
			return fmt.Errorf("unable to download status emoji image: %w", err)
		}

		statusDstR := compositeImage.Bounds()
		statusDstR.Min.Y = statusDstR.Max.Y / 2
		statusDstR.Max.X = statusDstR.Max.X / 2

		xdraw.BiLinear.Scale(compositeImage, statusDstR, statusImage, statusImage.Bounds(), xdraw.Over, nil)
	}

	imageDataBuf := bytes.Buffer{}
	err = png.Encode(&imageDataBuf, compositeImage)
	if err != nil {
		return fmt.Errorf("unable to encode PNG image: %w", err)
	}

	imageData := base64.StdEncoding.EncodeToString(imageDataBuf.Bytes())

	err = sdk.SetImage(a.context, fmt.Sprintf("data:image/png;base64,%s", imageData), 0)
	if err != nil {
		return fmt.Errorf("unable to set image: %w", err)
	}
	return nil
}
