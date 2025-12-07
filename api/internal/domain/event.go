package domain

type EventType string

const (
	RawMessageReceived   EventType = "raw_message_received"
	RawMessageNormalized EventType = "raw_message_normalized"
	RawMessageAnalyzed   EventType = "raw_message_analyzed"

	MediaReceived EventType = "media_received"
)
