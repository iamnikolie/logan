package nats

import (
	"errors"

	"github.com/nats-io/go-nats"

	"logan/logrus"
)

var (
	natsRegistryName string
	natsInst         *nats.Conn

	ConnectionError = errors.New("not found")
)

func Start() {
}

func init() {
	logrus.SetLevel()
	reconfigLoop()
}
