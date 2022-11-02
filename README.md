# Stream Deck Plugin for Slack

Provided Actions:

![image](https://user-images.githubusercontent.com/422474/199592000-9e7b7cf2-a1b8-4faa-95db-afffa679fa85.png)

- **Slack User** - shows a user's profile picture, presence indicator and custom status emoji
- **Slack Presence** - shows whether your presence is configured as away or "auto" and lets you toggle it
- **Slack Status** - allows you to toggle a custom status emoji

## Installation

There is currently an installation script for Windows (**install.ps1**), eventually I hope to add packaging automation and release this in the marketplace.

To use actions in this plugin, you need to create a Slack app with a user token with the following scopes:

- `users.profile:read`
- `users.profile:write` - for writing custom status
- `users:read`
- `users:write` - for updating presence information
