package mongo

import (
	"sync"
	"time"

	log "logan"

	"../shared"

	"gopkg.in/mgo.v2"
)

var (
	mongo      *MongoStore
	mongoReady bool
)

func init() {
	mongoReconfigureLoop()
}

func Start() {

}

func mongoReconfigureLoop() {
	configureOnce()
	go func() {
		lock := &sync.Mutex{}
		for {
			lock.Lock()
			configureOnce()
			lock.Unlock()

			time.Sleep(5 * time.Second)
		}
	}()
}

func configureOnce() {
	if mongo == nil {
		mongoReady = false
	}
	if mongoReady {
		return
	}

	url := shared.GetStringFromEnvDef("LOGAN_MONGO_HOST",
		"127.0.0.1:27017") + "/logan"
	m, err := NewMongoStore(url)
	if err != nil {
		log.Fatal("cant connect to mongodb on: mongodb://" + url)
		return
	}
	mongo = m
	mongoReady = true

	log.Printf("connected to mongodb on: mongodb://%s", url)
}

func Mongo() *mgo.Session {
	return mongo.Ses()
}

func MongoDB() (*mgo.Database, *mgo.Session) {
	return mongo.Db()
}

func MongoCol(name string) (*mgo.Collection, *mgo.Session) {
	return mongo.Col(name)
}
