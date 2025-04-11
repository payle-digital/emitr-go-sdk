package emitr

import "time"

type IncomingMessage struct {
	Offset    int64             `json:"offset"`
	Timestamp time.Time         `json:"timestamp"`
	Key       string            `json:"key,omitempty"`
	Headers   map[string]string `json:"headers,omitempty"`
	Payload   string            `json:"payload"`
}
