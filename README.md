[![CI](https://github.com/paultyng/streamdeck-slack/actions/workflows/ci.yml/badge.svg?branch=main&event=push)](https://github.com/paultyng/streamdeck-slack/actions/workflows/ci.yml)

# Stream Deck Plugin for Slack

Provided Actions:

![image](https://user-images.githubusercontent.com/422474/199592000-9e7b7cf2-a1b8-4faa-95db-afffa679fa85.png)

- **Slack User** - shows a user's profile picture, presence indicator and custom status emoji
- **Slack Presence** - shows whether your presence is configured as away or "auto" and lets you toggle it
- **Slack Status** - allows you to toggle a custom status emoji

## Installation

Create a Slack app using the included [manifest](./slack-app/manifest.yml). You can also use the included [icon](./slack-app/icon.png). Once it is created, copy your user token for use is the action button properties.

## Developing / Running from Source

There is currently an installation script for Windows (**install.ps1**).
