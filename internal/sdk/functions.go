package sdk

import "log"

func OpenURL(url string) error {
	return conn.WriteJSON(&sentEvent{
		Event:   EventOpenURL,
		Payload: &openUrlPayload{url},
	})
}

func GetGlobalSettings() error {
	return conn.WriteJSON(&sentEvent{Event: EventGetGlobalSettings, Context: PluginUUID})
}

func GetSettings(context string) error {
	return conn.WriteJSON(&sentEvent{Event: EventGetSettings, Context: context})
}

func SetGlobalSettings(v interface{}) error {
	return conn.WriteJSON(&sentEvent{Event: EventSetGlobalSettings, Context: PluginUUID, Payload: v})
}

func SetSettings(context string, v interface{}) error {
	return conn.WriteJSON(&sentEvent{Event: EventSetSettings, Context: context, Payload: v})
}

func Log(message string) {
	// TODO: handling logging when the websocket is down somehow? queue these up for
	// when we have a connection? write to a file?

	log.Print(message)
	if conn == nil {
		return
	}
	err := conn.WriteJSON(&sentEvent{
		Event:   EventLogMessage,
		Context: PluginUUID,
		Payload: &logMessagePayload{message},
	})
	if err != nil {
		log.Printf("error logging: %s", err)
	}
}

func SetTitle(context, title string, target int) error {
	return conn.WriteJSON(&sentEvent{
		Event:   EventSetTitle,
		Context: context,
		Payload: &setTitlePayload{Title: title, Target: target},
	})
}

func SetState(context string, state int) error {
	return conn.WriteJSON(&sentEvent{
		Event:   EventSetState,
		Context: context,
		Payload: &setStatePayload{state},
	})
}

func SetImage(context, imageData string, target int) error {
	return conn.WriteJSON(&sentEvent{
		Event:   EventSetImage,
		Context: context,
		Payload: &setImagePayload{imageData, target},
	})
}

func ShowAlert(context string) error {
	return conn.WriteJSON(&sentEvent{
		Event:   EventShowAlert,
		Context: context,
	})
}

func ShowOk(context string) error {
	return conn.WriteJSON(&sentEvent{
		Event:   EventShowOK,
		Context: context,
	})
}

func SendToPropertyInspector(context string, object interface{}) error {
	return conn.WriteJSON(&sentEvent{
		Event:   EventSendToPropertyInspector,
		Context: context,
		Payload: object,
	})
}
