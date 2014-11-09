# gomon/mongod

[![GoDoc](https://godoc.org/github.com/gomon/mongod?status.svg)](http://godoc.org/github.com/gomon/mongod)

A simple start/stop struct for *mgo* sessions.

## Example
    
    m := mongod.New("databaseName")
    db, err := m.Start()
    if err != nil {
      // handle dial error
    }
    defer m.Stop()

    // do mongo stuff

---

If you need to do something just prior to `Stop`, eg. Test teardowns. `Stop` accepts `func(s)` that will be run before the session is closed

    var cleandb = func(db *mgo.Database) {
      for _, v := range []string{
        "collection1",
        "collection2",
      } {
        c := db.C(v)
        _, err := c.RemoveAll(bson.M{})
        if err != nil {
          panic(err)
        }
      }
    }

    defer m.Stop(cleandb)

---

`mongod` will return the database from a clone of the original `Dial` session. You can call `Session` to grab the original session.

    s := m.Session() // original Dial session

## License

MIT
