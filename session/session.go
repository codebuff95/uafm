package session

import(
  "uafm/database"
  "gopkg.in/mgo.v2"
  "gopkg.in/mgo.v2/bson"
  "time"
  "strings"
)

var UserSM *SessionManager
var FormSM *SessionManager

type SessionManager struct{
  Collection *mgo.Collection
}

type Session struct{
  Sid string `bson:"sid"`
  Rid string `bson:"rid"`
  Expires time.Time `bson:"expires"`
  Status int `bson:"-"`
}

const (
	EXPIRED int = 1
	ACTIVE  int = 2
	DELETED int = 3
	//SIDLEN Length of each SessionID.
	SIDLEN int = 16
)

func InitSMs(usercollectionname, formcollectionname string){
  UserSM = &SessionManager{Collection: database.DatabaseSession.DB(database.DatabaseName).C(usercollectionname)}
  FormSM = &SessionManager{Collection: database.DatabaseSession.DB(database.DatabaseName).C(formcollectionname)}
}

func GenerateUniqueSid() string{
  return "aeiouasd"
}

func (sm *SessionManager) GetSession(sid string) (*Session,error){
  var mySession *Session = &Session{}
  err := sm.Collection.Find(bson.M{"sid":sid}).Limit(1).One(mySession)
  if err != nil{
    return nil,err
  }
  if mySession == nil{
    return nil,nil
  }
  mySession.Status = ACTIVE
  if mySession.Expires.IsZero(){
    //Non-expirable session.
    return mySession,nil
  }
  if i := strings.Compare(mySession.Expires.Format("2006-01-02 15:04:05"), time.Now().Format("2006-01-02 15:04:05")); i <= 0 {
    mySession.Status = EXPIRED
	}
  return mySession,nil
}


func (sm *SessionManager) SetSession(rid string, life time.Duration) (*Session,error){
  mySession := &Session{}
  mySession.Sid = GenerateUniqueSid()
  mySession.Expires = time.Now().Add(life)
  mySession.Rid = rid
  err := sm.Collection.Insert(mySession)
  if err != nil{
    return nil,err
  }
  mySession.Status = ACTIVE
  return mySession,nil
}


func (sm *SessionManager) DeleteSession(sid string) error {
	err := sm.Collection.Remove(bson.M{"sid":sid})
  return err
}
