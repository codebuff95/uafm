package database

import(
  "gopkg.in/mgo.v2"
  //"gopkg.in/mgo.v2/bson"
  "encoding/json"
  "os"
)

var DatabaseName string
var DatabaseSession *mgo.Session

type DatabaseDetails struct{
  url string `json:"url"`
  dbname string `json:"dbname"`
}

type InitDatabaseError int

func (err InitDatabaseError) Error() string{
  if err == PROBLEMOPENINGFILE{
    return "Could not open file uafmconfigs/dbconfig.json"
  }
  return "Miscellaneous Issues in database"
}

const(
  PROBLEMOPENINGFILE = 0
  PROBLEMDECODING = 1
)

func InitDatabaseSession() error{
  var err error

  myFile, err := os.Open("uafmconfigs/dbconfig.json")
  defer myFile.Close()

  if err != nil{
    return InitDatabaseError(PROBLEMOPENINGFILE)
  }
  var myDBDetails DatabaseDetails
  err = json.NewDecoder(myFile).Decode(&myDBDetails)
  if err != nil{
    return InitDatabaseError(PROBLEMDECODING)
  }
  DatabaseSession,err = mgo.Dial(myDBDetails.url)
  //defer DatabaseSession.Close().
  DatabaseName = myDBDetails.dbname
  return nil
}
