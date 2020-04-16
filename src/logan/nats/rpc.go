package nats

import (
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/nats-io/go-nats"
)

func Ready() bool {
	return !(natsInst == nil || !natsInst.IsConnected())
}

func Pub(top string, msg []byte) error {
	if natsInst == nil || !natsInst.IsConnected() {
		return ConnectionError
	}
	err := natsInst.Publish(top, msg)
	if err != nil {
		log.Errorf("error while publish message: %s", err.Error())
	}
	natsInst.Flush()
	return nil
}

func Sub(top string, handler func(msg *Msg)) error {
	var err error
	if Ready() {
		if _, err = natsInst.Subscribe(top, func(nm *nats.Msg) {
			handler(&Msg{*nm})
		}); err != nil {
			log.Printf("error while subscribe on topic %s: %s",
				top,
				err.Error())
		}
	} else {
		log.Errorf("%s", err.Error())
	}
	return err
}

func Request(subj string, msg []byte) (*Msg, error) {
	if natsInst == nil || !natsInst.IsConnected() {
		return nil, ConnectionError
	}
	t := time.Duration(GetIntFromEnvDef("LOGAN_MBUS_RPC_TIMEOUT",
		60))
	res, err := natsInst.Request(subj, msg, t*time.Second)
	if err != nil {
		return nil, err
	}
	return &Msg{*res}, nil
}
