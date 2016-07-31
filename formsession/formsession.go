package formsession

import(
  "time"
  "github.com/codebuff95/uafm/session"
)

type FormValidationError int

const(
  MISC = 0
  DOESNOTEXIST = 1
  EXPIRED = 2
)

func (err FormValidationError) Error() string{
  if err == DOESNOTEXIST{
    return "Form Session does not exist"
  }
  if err == EXPIRED{
    return "Form Session has expired"
  }
  return "Miscellaneous Form Session Error"
}

func CreateSession(rid string, life time.Duration) (*string,error){
  mySession,err := session.FormSM.SetSession(rid,life)
  if err != nil{
    return nil,err
  }
  return &mySession.Sid,nil
}

func ValidateSession(sid string) (*string,error){
  mySession,err := session.FormSM.GetSession(sid)
  if err != nil{
    return nil,FormValidationError(MISC)
  }
  if mySession == nil{
    return nil,FormValidationError(DOESNOTEXIST)
  }
  if mySession.Status == session.EXPIRED{
    return nil,FormValidationError(EXPIRED)
  }
  return &mySession.Rid,nil
}

//DeleteSession returns the number of form sessions deleted, and a non-nil error if an error occured (in which case,
// number of sessions deleted = 0).
func DeleteSession(sid string) (int,error){
  return session.FormSM.DeleteSession(sid)
}
