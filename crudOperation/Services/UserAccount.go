package services

import (
	"CurdOperation/Model"
	"fmt"
	"log"

	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

type IUserAccount interface {
	CreateUserAccount(Ac Model.Signup) (string, error)
	Login(logreq Model.LoginReq) (string, error)
	HashPassowrd(Password string) (string, error)
	IsvalidPassword(Hasedpassword string, ReqPass string) bool
}
type UserAccountController struct {
	db       Idb
	Payload  IPayload
	readjson IReadJson
}

func UserAccountCtor(db Idb, Payload IPayload, readjson IReadJson) *UserAccountController {
	return &UserAccountController{db: db, Payload: Payload, readjson: readjson}
}

func (UC *UserAccountController) CreateUserAccount(Ac Model.Signup) (string, error) {
	db := UC.db.IntializeDB()
	query := `INSERT INTO UserAccount (PRNumber, UserName, Password, IsAccountLocked, WrongPasswordCount,  FirstVisit,  IsAccountDeleted)
	VALUES ($1, $2,$3, $4,$5,$6,$7);`
	Hashpass, _ := UC.HashPassowrd(Ac.Password)
	fmt.Println("Hashed Password :", Hashpass)
	_, err := db.Exec(query, Ac.PRN, Ac.UserName, Hashpass, false, 0, true, false)
	if err != nil {
		fmt.Println(err)
		return " ", err
	}
	return "User Account created successfully", nil
}

// login user
func (uc *UserAccountController) Login(logreq Model.LoginReq) (string, error) {
	log.Println("login username " + logreq.UserName + " and password " + logreq.Password + " @  Useraccount >> Login")
	db := uc.db.IntializeDB()
	query := `SELECT password FROM UserAccount WHERE UserName=$1;`
	var password string
	//check valid clientid
	JsonData, err := uc.readjson.ReadJson()
	if err != nil {
		log.Panic(err)
		return " ", errors.New("Error while Reading Json data in useraccount ")
	}
	if logreq.ClientId != JsonData.ClientID {
		log.Println(" JsonData.ClientID", JsonData.ClientID)
		return " ", errors.New("Invalid ClientID  ")
	}

	errr := db.QueryRow(query, logreq.UserName).Scan(&password)

	if errr != nil {
		log.Panic(err)
		return " ", errors.New("Error while scanning useraccount ")
	}
	//fmt.Println("UserAccount.Password ", logreq.Password)
	IsUser := uc.IsvalidPassword(password, logreq.Password)

	//IsUser := uc.IsvalidPassword(UserAccount.Password, logreq.Password)
	if !IsUser {
		return " ", errors.New("Invalid Password")

	}

	//Generate Token

	token, err := uc.Payload.GenerateToken(logreq.ClientId, logreq.UserName)
	return token, err
}

func (UC *UserAccountController) HashPassowrd(Password string) (string, error) {
	Hasedpassword, err := bcrypt.GenerateFromPassword([]byte(Password), bcrypt.DefaultCost)
	if err != nil {
		return " ", err
	}
	return string(Hasedpassword), nil
}
func (UC *UserAccountController) IsvalidPassword(Hasedpassword string, ReqPass string) bool {
	fmt.Println("Login>>IsvalidPassword")
	err := bcrypt.CompareHashAndPassword([]byte(Hasedpassword), []byte(ReqPass))
	if err != nil {
		return err == nil
	}
	return true
}
