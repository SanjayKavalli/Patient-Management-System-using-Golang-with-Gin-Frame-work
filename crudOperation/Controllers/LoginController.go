package controllers

import (
	"CurdOperation/Model"
	services "CurdOperation/Services"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SignUpController struct {
	useracc services.IUserAccount
	payload services.IPayload
}

func SignUpCtor(useracc services.IUserAccount, payload services.IPayload) *SignUpController {
	return &SignUpController{useracc: useracc, payload: payload}
}
func (su *SignUpController) SignUp(ctx *gin.Context) {
	var SignUp *Model.Signup
	if err := ctx.ShouldBindJSON(&SignUp); err != nil {

		ctx.JSON(http.StatusInternalServerError, gin.H{"error ": "Internal server error please try again later"})
		panic(err)

	}
	res, _ := su.useracc.CreateUserAccount(*SignUp)
	ctx.JSON(http.StatusOK, gin.H{"message": res})

}

func (su *SignUpController) Login(ctx *gin.Context) {
	fmt.Println("signup>>Login")
	var logreq *Model.LoginReq
	if err := ctx.ShouldBindJSON(&logreq); err != nil {

		ctx.JSON(http.StatusInternalServerError, gin.H{"error ": "Internal server error please try again later"})
		panic(err)

	}
	res, err := su.useracc.Login(*logreq)
	ctx.JSON(http.StatusOK, gin.H{"message": res, "Error": err})
}

func (su *SignUpController) GenerateToken(ctx *gin.Context) {
	var Req services.LoginReq
	//fmt.Println(Req.CLientID, Req.UserName, Req.Password)
	if err := ctx.ShouldBindJSON(&Req); err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error ": "Unauthorized LOgin"})
		return
	}

	token, err := su.payload.GenerateToken(Req.CLientID, Req.UserName)
	ctx.JSON(http.StatusOK, gin.H{"token": token, "err": err})

}

func (su *SignUpController) VerifyLoginToken(ctx *gin.Context) {
	var Req struct {
		Token string `json:"token"`
	}
	if err := ctx.ShouldBindJSON(&Req); err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error ": "invalid token"})
		return
	}
	payload, err := su.payload.VerifyToken(Req.Token)
	ctx.JSON(http.StatusOK, gin.H{"payload": payload, "err": err})

}
