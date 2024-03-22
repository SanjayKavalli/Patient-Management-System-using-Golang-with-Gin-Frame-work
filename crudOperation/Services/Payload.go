package services

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/o1egl/paseto"
)

type LoginReq struct {
	CLientID string
	UserName string
	Password string
}
type Payload struct {
	ID        uuid.UUID `json:"id"`
	CLientID  string    `json:"client_id"`
	UserName  string    `json:"user_name"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiresIn time.Time `json:"expires_in"`
}
type IPayload interface {
	NewPayload(CLientID string, UserName string) (*Payload, error)
	GenerateToken(CLientID string, UserName string) (string, error)
	VerifyToken(token string) (*Payload, error)
}
type PayloadController struct {
	db Idb
}

func Payloadctor(db Idb) *PayloadController {
	return &PayloadController{db: db}
}

func (pc *PayloadController) NewPayload(CLientID string, UserName string) (*Payload, error) {
	tokenid, err := uuid.NewRandom()
	fmt.Println("New payload")
	fmt.Println("NewPayload UUId ", tokenid)
	fmt.Println("payload username ", UserName)

	if err != nil {
		return nil, err
	}
	Payload := &Payload{
		ID:        tokenid,
		CLientID:  CLientID,
		UserName:  UserName,
		IssuedAt:  time.Now(),
		ExpiresIn: time.Now().Add(60 * time.Minute),
	}
	fmt.Println("payload.username ", Payload.UserName)
	db := pc.db.IntializeDB()
	query := `UPDATE UserAccount SET SessionId=$1 WHERE UserName=$2`
	_, err = db.Exec(query, tokenid, UserName)
	if err != nil {
		return nil, err
	}
	return Payload, nil

}

func (pc *PayloadController) GenerateToken(CLientID string, UserName string) (string, error) {
	symmetricKey := []byte("0C615CA8BEC0415190D068D66C6838E2")
	fmt.Println("Symmetric ", symmetricKey)
	fmt.Println("username"+UserName, "ClientId"+CLientID)

	Payload, _ := pc.NewPayload(CLientID, UserName)
	//fmt.Println("payload ", Payload.UserName)
	token, err := paseto.NewV2().Encrypt(symmetricKey, Payload, nil)
	if err != nil {
		return " notoken", err
	}
	fmt.Println("TOken ", token)
	return token, err

}

func (pc *PayloadController) VerifyToken(token string) (*Payload, error) {
	fmt.Println("Ver token ", token)
	Payload := &Payload{}
	fmt.Println("verify token")
	symmetricKey := []byte("0C615CA8BEC0415190D068D66C6838E2")
	err := paseto.NewV2().Decrypt(token, symmetricKey, Payload, nil)
	if err != nil {
		return nil, err
	}
	fmt.Println(Payload.UserName)
	fmt.Println("VerifyToken id ", Payload.ID)
	return Payload, nil
}
