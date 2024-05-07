package MessageQueue

import (
	"encoding/json"
	application "goOtp/internal/ports/Application"
)

type message struct {
	Operation   string          `json:"operation"`
	FromService string          `json:"fromService"`
	Data        json.RawMessage `json:"data"`
}

type messageHandler func(message, application.ApplicationInterface) bool
