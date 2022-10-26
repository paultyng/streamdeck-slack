# Stream Deck Plugin for Slack

Provided Actions:

- **Slack User** - shows a user's profile picture, presence indicator and custom status emoji
- **Slack Presence** - shows whether your presence is configured as away or "auto" and lets you toggle it
- **Slack Status** - allows you to toggle a custom status emoji

The actions all require a user token with the following scopes:

- `users.profile:read`
- `users.profile:write` - for writing custom status
- `users:read`
- `users:write` - for updating presence information

## Some things to investigate

- [ ] Socket Mode / RTM for presence? Is this possible?
