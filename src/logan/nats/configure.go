package nats

import (
	"os"
	"strconv"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/nats-io/go-nats"
)

func reconfigLoop() {
	configureOnce()
	go func() {
		for {
			configureOnce()

			time.Sleep(ReconnectionTimeout)
		}
	}()
}

func configureOnce() {
	url := "nats://" + GetStringFromEnvDef("LOGAN_NATS_HOST",
		"127.0.0.1:4242")

	opts := nats.GetDefaultOptions()
	opts.MaxReconnect = -1
	opts.ReconnectWait = ReconnectionTimeout
	opts.Url = url

	if natsInst != nil && natsInst.ConnectedUrl() == url {
		return
	}

	opts.DisconnectedCB = func(nc *nats.Conn) {
		// do nothing
	}

	opts.ReconnectedCB = func(nc *nats.Conn) {
		// do nothing
	}

	nc, err := opts.Connect()
	if err != nil {
		log.Errorf("error on connect to %s: %s", url, err.Error())
		return
	}

	if natsInst != nil {
		log.Errorf("change message bus connection from %s to %s",
			natsInst.ConnectedUrl(),
			nc.ConnectedUrl())
		natsInst.Close()
	}

	natsInst = nc
}

func GetStringFromEnvDef(param string, def string) string {
	if retVal := os.Getenv(param); retVal != "" {
		return retVal
	}
	return def
}

func GetIntFromEnvDef(key string, def int) int {
	if val := os.Getenv(key); val != "" {
		r, err := strconv.Atoi(val)
		if err != nil {
			return def
		}
		return r
	}
	return def
}
