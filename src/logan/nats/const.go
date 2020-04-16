package nats

import "time"

const (
	ReconnectionTimeout = 1 * time.Minute

	EncodingBinary Encoding = 0
	EncodingJSON   Encoding = 1
	EncodingBSON   Encoding = 3

	EncodingKey = "encv1"
)
