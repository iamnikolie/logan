package logan

import (
	"encoding/json"
	"fmt"
	"time"

	"logan/nats"
	"logan/types"

	log "github.com/sirupsen/logrus"
)

func Printf(format string, args ...interface{}) {
	switch getStoreMethod() {
	case Nats:
		logMessage := types.LogMessage{
			Level:     "info",
			Timestamp: time.Now(),
			Service:   getServiceName(),
			Message:   fmt.Sprintf(format, args...),
		}
		msg, err := json.Marshal(logMessage)
		if err != nil {
			log.WithFields(log.Fields{
				"service": getServiceName(),
			}).Printf("cant unmarshal message: %s", err.Error())
		}
		nats.Pub("logan.logs.printf", msg)
	default:
		log.WithFields(log.Fields{
			"service": getServiceName(),
		}).Printf(format, args...)
	}
}

func Trace(args ...interface{}) {
	switch getStoreMethod() {
	case Nats:
		logMessage := types.LogMessage{
			Level:     "trace",
			Timestamp: time.Now(),
			Service:   getServiceName(),
			Message:   fmt.Sprint(args...),
		}
		msg, err := json.Marshal(logMessage)
		if err != nil {
			log.WithFields(log.Fields{
				"service": getServiceName(),
			}).Printf("cant unmarshal message: %s", err.Error())
		}
		nats.Pub("logan.logs.trace", msg)
	default:
		log.WithFields(log.Fields{
			"service": getServiceName(),
		}).Trace(args...)
	}
}

func Debug(args ...interface{}) {
	switch getStoreMethod() {
	case Nats:
		logMessage := types.LogMessage{
			Level:     "debug",
			Timestamp: time.Now(),
			Service:   getServiceName(),
			Message:   fmt.Sprint(args...),
		}
		msg, err := json.Marshal(logMessage)
		if err != nil {
			log.WithFields(log.Fields{
				"service": getServiceName(),
			}).Printf("cant unmarshal message: %s", err.Error())
		}
		nats.Pub("logan.logs.debug", msg)
	default:
		log.WithFields(log.Fields{
			"service": getServiceName(),
		}).Debug(args...)
	}
}

func Info(args ...interface{}) {
	switch getStoreMethod() {
	case Nats:
		logMessage := types.LogMessage{
			Level:     "info",
			Timestamp: time.Now(),
			Service:   getServiceName(),
			Message:   fmt.Sprint(args...),
		}
		msg, err := json.Marshal(logMessage)
		if err != nil {
			log.WithFields(log.Fields{
				"service": getServiceName(),
			}).Printf("cant unmarshal message: %s", err.Error())
		}
		nats.Pub("logan.logs.info", msg)
	default:
		log.WithFields(log.Fields{
			"service": getServiceName(),
		}).Info(args...)
	}
}

func Error(args ...interface{}) {
	switch getStoreMethod() {
	case Nats:
		logMessage := types.LogMessage{
			Level:     "error",
			Timestamp: time.Now(),
			Service:   getServiceName(),
			Message:   fmt.Sprint(args...),
		}
		msg, err := json.Marshal(logMessage)
		if err != nil {
			log.WithFields(log.Fields{
				"service": getServiceName(),
			}).Printf("cant unmarshal message: %s", err.Error())
		}
		nats.Pub("logan.logs.error", msg)
	default:
		log.WithFields(log.Fields{
			"service": getServiceName(),
		}).Error(args...)
	}
}

func Errorf(layout string, args ...interface{}) {
	switch getStoreMethod() {
	case Nats:
		logMessage := types.LogMessage{
			Level:     "error",
			Timestamp: time.Now(),
			Service:   getServiceName(),
			Message:   fmt.Sprintf(layout, args...),
		}
		msg, err := json.Marshal(logMessage)
		if err != nil {
			log.WithFields(log.Fields{
				"service": getServiceName(),
			}).Printf("cant unmarshal message: %s", err.Error())
		}
		nats.Pub("logan.logs.error", msg)
	default:
		log.WithFields(log.Fields{
			"service": getServiceName(),
		}).Errorf(layout, args...)
	}
}

func Fatal(args ...interface{}) {
	switch getStoreMethod() {
	case Nats:
		logMessage := types.LogMessage{
			Level:     "fatal",
			Timestamp: time.Now(),
			Service:   getServiceName(),
			Message:   fmt.Sprint(args...),
		}
		msg, err := json.Marshal(logMessage)
		if err != nil {
			log.WithFields(log.Fields{
				"service": getServiceName(),
			}).Printf("cant unmarshal message: %s", err.Error())
		}
		nats.Pub("logan.logs.fatal", msg)
	default:
		log.WithFields(log.Fields{
			"service": getServiceName(),
		}).Fatal(args...)
	}
}
