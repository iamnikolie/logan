package nats

import "github.com/nats-io/go-nats"

type Encoding byte

type Msg struct {
	nats.Msg
}
