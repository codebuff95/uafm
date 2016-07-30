package user

import(
  "time"
  "uafm/session"
)

type UserValidationError int

const(
  MISC = 0
  DOESNOTEXIST = 1
  EXPIRED = 2
)

func (err UserValidationError) Error() string{
  if err == DOESNOTEXIST{
    return "User Session does not exist"
  }
  if err == EXPIRED{
    return "User Session has expired"
  }
  return "Miscellaneous User Session Error"
}

func CreateSession(rid string, life time.Duration) (*string,error){
  mySession,err := session.UserSM.SetSession(rid,life)
  if err != nil{
    return nil,err
  }
  return &mySession.Sid,nil
}

func ValidateSession(sid string) (*string,error){
  mySession,err := session.UserSM.GetSession(sid)
  if err != nil{
    return nil,UserValidationError(MISC)
  }
  if mySession == nil{
    return nil,UserValidationError(DOESNOTEXIST)
  }
  if mySession.Status == session.EXPIRED{
    return nil,UserValidationError(EXPIRED)
  }
  return &mySession.Rid,nil
}

//DeleteSession returns the number of user sessions deleted, and a non-nil error if an error occured
//(in which case, the number of sessions deleted = 0).
func DeleteSession(sid string) (int,error){
  return session.UserSM.DeleteSession(sid)
}
