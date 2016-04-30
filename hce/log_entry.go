package hce

import (
    "time"
)

// A log entry returned in an ErrorResponse.\n
type LogEntry struct {
    // E.g. `error`, `warn` etc. These should become an enum.\n
    Level  string  `json:"level,omitempty"`
    // The log message.\n
    Message  string  `json:"message,omitempty"`
    // The timestamp, in `2016-04-17T11:40:43.448Z` format.\n
    Timestamp  time.Time  `json:"timestamp,omitempty"`
    // Individual error items.\n
    Items  []ErrorItem  `json:"items,omitempty"`
}
