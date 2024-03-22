package Model

import "time"

type UserAccount struct {
	UserAccountId      int        `json:"useraccountid"`
	PRN                int        `json:"prnumber"`
	UserName           string     `json:"username"`
	Password           string     `json:"password"`
	IsAccountLocked    *bool      `json:"isaccountlocked,omitempty"`
	WrongPasswordCount int        `json:"wrongpasswordcount"`
	ExpiredAt          *time.Time `json:"expiredat"`
	AccountLockedAt    *time.Time `json:"accountlockedat"`
	FirstVisit         bool       `json:"firstvisit"`
	SessionId          string     `json:"sessionid"`
	SessionExpiresIn   *time.Time `json:"sessionexpiresin"`
	IsAccountDeleted   bool       `json:"isaccountdeleted"`
}
type Signup struct {
	PRN      int
	UserName string
	Password string
}
