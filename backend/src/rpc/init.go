package rpc

import (
	"logan/nats"
)

func init() {
	nats.Sub("logan.logs.*", log)
}
