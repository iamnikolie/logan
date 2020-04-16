package rpc

import (
	"encoding/json"
	"logan"
	"logan/nats"
	"logan/types"

	"../mongo"
	"../shared"
	"gopkg.in/mgo.v2/bson"

	l "github.com/sirupsen/logrus"
)

func log(msg *nats.Msg) {
	e := types.LogMessage{}
	err := json.Unmarshal(msg.Data, &e)
	if err != nil {
		logan.Error(err.Error())
		return
	}
	storeType := shared.GetStringFromEnvDef("LOGAN_SERVER_STORE_METHOD", "syslog")
	switch storeType {
	default:
		l.SetLevel(l.TraceLevel)
		entry := l.WithFields(l.Fields{
			"service": e.Service,
		})
		switch e.Level {
		case "debug":
			entry.Debug(e.Message)
			break
		case "info":
			entry.Info(e.Message)
			break
		case "fatal":
			entry.Fatal(e.Message)
			break
		case "trace":
			entry.Trace(e.Message)
			break
		case "error":
			entry.Error(e.Message)
			break
		case "warn":
			entry.Warn(e.Message)
			break
		}
	}

	if mongoUse := shared.GetStringFromEnvDef("LOGAN_STORE_IN_DB", "false"); mongoUse == "true" {
		e.ID = bson.NewObjectId()
		col, ses := mongo.MongoCol(e.Level)
		defer ses.Close()
		if err := col.Insert(&e); err != nil {
			l.Errorf("error while inserting, %s", err.Error())
			return
		}
	}
}
