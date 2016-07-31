package uafm

import(
  "github.com/codebuff95/uafm/database"
  "github.com/codebuff95/uafm/session"
)

//Driver package for uafm.

func Init(dirtoconfig string) error{
  err := database.InitDatabaseSession(dirtoconfig)
  if err != nil{
    return err
  }
  session.InitSMs("usersession","formsession")
  //go session.UserSM.Clean()
  //go session.FormSM.Clean()
  return nil
  //Initialising collection cleaners.

  /*var mySession *session.Session
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
  }else{
    log.Println("mySession is active:",mySession)
  }
  deleted,err := session.UserSM.DeleteSession(mySession.Sid)
  if err != nil{
    log.Fatal("Could not delete session",err)
  }
  log.Println("Deleted",deleted,"sessions!")*/
}
