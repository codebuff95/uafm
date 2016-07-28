package main

import(
  "uafm/database"
  "uafm/session"
  "log"
  "time"
)

func main(){
  err := database.InitDatabaseSession("127.0.0.1","sessionmanagementtry")
  if err != nil{
    log.Fatal(err)
  }
  session.InitSMs("usersession","formsession")
  var mySession *session.Session
  mySession, err = session.UserSM.SetSession("mymyrid1",time.Hour)
  if err != nil{
    log.Fatal("Could not SetSession:",err)
  }
  log.Println("Successfully SetSession with Sid:",mySession.Sid)
  mySession,err = session.UserSM.GetSession(mySession.Sid)
  if err != nil{
    log.Fatal("Could not GetSession:",err)
  }
  if mySession == nil{
    log.Fatal("mySession is nil")
  }
  if mySession.Status == session.EXPIRED{
    log.Println("mySession is expired:",mySession)
  }
  log.Println("mySession is active:",mySession)
  err = session.UserSM.DeleteSession(mySession.Sid)
  if err != nil{
    log.Fatal("Could not delete session",err)
  }
  log.Println("Deleted session!")
}
