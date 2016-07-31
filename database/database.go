package database

import(
  "gopkg.in/mgo.v2"
  "log"
  "encoding/json"
  "os"
)

var DatabaseName string
var DatabaseSession *mgo.Session

type DatabaseDetails struct{
  Url string `json:"url"`
  Dbname string `json:"dbname"`
}

type InitDatabaseError int

func (err InitDatabaseError) Error() string{
  if err == PROBLEMOPENINGFILE{
    return "Could not open file uafmconfig.json"
  }
  return "Miscellaneous Issues in database"
}

const(
  PROBLEMOPENINGFILE = 0
  PROBLEMDECODING = 1
)

func InitDatabaseSession(dirtoconfig string) error{
  var err error

  myFile, err := os.Open(dirtoconfig+"/uafmconfig.json")
  defer myFile.Close()

  if err != nil{
    return InitDatabaseError(PROBLEMOPENINGFILE)
  }
  var myDBDetails DatabaseDetails
  err = json.NewDecoder(myFile).Decode(&myDBDetails)
  if err != nil{
    return InitDatabaseError(PROBLEMDECODING)
  }
  DatabaseSession,err = mgo.Dial(myDBDetails.Url)
  DatabaseName = myDBDetails.Dbname
  log.Println("UAFM: Initialised new database session to url:",myDBDetails.Url,"and Dbname:",myDBDetails.Dbname)
  return nil
}
