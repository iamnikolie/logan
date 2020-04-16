package mongo

import (
	"sync"
	"time"

	"gopkg.in/mgo.v2"
)

type MongoStore struct {
	dial    *mgo.DialInfo
	session *mgo.Session
	lock    sync.Locker
	watched bool
}

func (s *MongoStore) Ses() *mgo.Session {
	return s.session.Clone()
}

func (s *MongoStore) Db() (*mgo.Database, *mgo.Session) {
	ses := s.Ses()
	return ses.DB(s.dial.Database), ses
}

func (s *MongoStore) Col(name string) (*mgo.Collection, *mgo.Session) {
	ses := s.Ses()
	return ses.DB(s.dial.Database).C(name), ses
}

func (s *MongoStore) Configure(uri string) error {
	dial, err := mgo.ParseURL(uri)
	if err != nil {
		return err
	}
	s.lock.Lock()
	defer s.lock.Unlock()

	s.dial = dial
	if s.session != nil {
		s.session.Close()
		s.session = nil
	}

	s.connect()
	if !s.watched {
		s.watched = true
		go s.watch()
	}
	return nil
}

func (s *MongoStore) watch() error {
	for {
		s.lock.Lock()
		if s.session != nil {
			if err := s.session.Ping(); err != nil {
				// TODO write to log
				s.session.Close()
				s.session = nil
			}
		}
		err := s.connect()
		s.lock.Unlock()
		if err != nil {
			// TODO write to log
			time.Sleep(time.Second * 50)
		}
		time.Sleep(time.Second * 10)
	}
}

func (s *MongoStore) connect() error {
	if s.session != nil {
		return nil
	}
	ses, err := mgo.DialWithInfo(s.dial)
	if err != nil {
		return err
	}
	ses.SetSafe(&mgo.Safe{
		W:        1,
		WMode:    "majority",
		WTimeout: 3000,
		FSync:    true,
	})
	s.session = ses
	return nil
}

func NewMongoStore(url string) (*MongoStore, error) {
	s := MongoStore{lock: &sync.Mutex{}}
	err := s.Configure(url)
	return &s, err
}
