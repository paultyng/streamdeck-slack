module streamdeck-slack

go 1.19

require (
	github.com/gorilla/websocket v1.5.0
	github.com/slack-go/slack v0.11.3
	github.com/valyala/fastjson v1.6.3
	golang.org/x/image v0.1.0
)

replace github.com/slack-go/slack => github.com/paultyng/slack v0.0.0-20221031165840-048adff3fa23
