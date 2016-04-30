package hce

import (
    "time"
)

// An individual error item.\n
type ErrorItem struct {
    // The affected property\n
    Property  string  `json:"property,omitempty"`
    // The log message.\n
    Message  string  `json:"message,omitempty"`
    // The timestamp, in `2016-04-17T11:40:43.448Z` format.\n
    Timestamp  time.Time  `json:"timestamp,omitempty"`
}
