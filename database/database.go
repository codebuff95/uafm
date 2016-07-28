package database

import(
  "gopkg.in/mgo.v2"
  //"gopkg.in/mgo.v2/bson"
)

var DatabaseName string
var DatabaseSession *mgo.Session

func InitDatabaseSession(url,dbname string) error{
  var err error
  DatabaseSession,err = mgo.Dial(url)
  //defer DatabaseSession.Close().
  DatabaseName = dbname
  return err
}
