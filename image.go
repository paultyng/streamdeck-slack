package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	"image/png"
	"net/http"
	"os"
	"path/filepath"

	_ "image/gif"  // gif support for custom emojis
	_ "image/jpeg" // jpeg support for gravatar proxy

	"streamdeck-slack/internal/sdk"
)

var (
	activeImage = mustDecodeImage("presenceActiveIcon.png")
	awayImage   = mustDecodeImage("presenceAwayIcon.png")
)

func mustDecodeImage(path string) image.Image {
	f, err := os.Open(path)
	if err != nil {
		// if this failed, check subdir just in case running in tests or something
		path = filepath.Join("./com.paultyng.slack.sdPlugin/", path)
		f, err = os.Open(path)
		if err != nil {
			panic(err)
		}
	}
	defer f.Close()
	img, _, err := image.Decode(f)
	if err != nil {
		panic(err)
	}
	return img
}

// TODO: cache images based on URL in memory
func imageFromURL(url string) (image.Image, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("unable to create request: %w", err)
	}

	req.Header.Set("User-Agent", "streamdeck-slack (https://github.com/paultyng/streamdeck-slack)")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("unable to do request: %w", err)
	}
	defer resp.Body.Close()

	img, _, err := image.Decode(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("unable to decode image: %w", err)
	}

	return img, nil
}

func setImage(context string, img image.Image, state *int) error {
	imageDataBuf := bytes.Buffer{}
	err := png.Encode(&imageDataBuf, img)
	if err != nil {
		return fmt.Errorf("unable to encode PNG image: %w", err)
	}

	imageData := base64.StdEncoding.EncodeToString(imageDataBuf.Bytes())

	err = sdk.SetImage(context, fmt.Sprintf("data:image/png;base64,%s", imageData), 0, state)
	if err != nil {
		return fmt.Errorf("unable to set image: %w", err)
	}
	return nil
}
