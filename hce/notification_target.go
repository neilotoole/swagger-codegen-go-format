package hce

import (

)

// Models a notification, (eg. hipChat Notification).
type NotificationTarget struct {
	// The internal id of this notification instance.
	Id int32 `json:"id,omitempty"`
	// The name of the notification target
	Name string `json:"name,omitempty"`
	// The type the notification target, (eg. hipChat Notification).
	Type_ string `json:"type,omitempty"`
	// The url where to send the notification.
	Location string `json:"location,omitempty"`
	// The token to send notification to the location.
	Token string `json:"token,omitempty"`
}
