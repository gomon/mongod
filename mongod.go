package mongod

import (
	"gopkg.in/mgo.v2"
)

type Mongod interface {
	Session() *mgo.Session
	Start() (*mgo.Database, error)
	Stop(...func(*mgo.Database))
}

// mongod provides a way to start/stop an mgo.Session
type mongod struct {
	session *mgo.Session
	c       *mgo.Session
	db      *mgo.Database

	Addr     string
	Database string
}

var _ Mongod = &mongod{}

func New(name string, opts ...func(m *mongod)) *mongod {
	m := &mongod{
		Addr:     "127.0.0.1:27017",
		Database: name,
	}

	for _, v := range opts {
		v(m)
	}
	return m
}

// Session returns the original mgo.Sesssion
func (m mongod) Session() *mgo.Session {
	return m.session
}

// Start dials the mongo server and returns a mgo.Database through a clone of
// the original dial session
func (m *mongod) Start() (*mgo.Database, error) {
	s, err := mgo.Dial(m.Addr)
	if err != nil {
		return nil, err
	}
	m.session = s

	m.c = s.Clone()
	m.db = m.c.DB(m.Database)
	return m.db, nil
}

// Stop closes both the original session and the initial clone made to grab the
// database.
// It also supports running funcs against the database before the sessions get
// closed.
func (m *mongod) Stop(fn ...func(*mgo.Database)) {
	defer m.session.Close()
	defer m.c.Close()

	for _, v := range fn {
		v(m.db)
	}
}
