package middleware

import (
	services "CurdOperation/Services"
	"fmt"
	"log"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type AuthmiddlewareController struct {
	db services.Idb
	pl services.IPayload
}

func AuthCtor(db services.Idb, pload services.IPayload) *AuthmiddlewareController {
	return &AuthmiddlewareController{db: db,
		pl: pload}
}

func (am *AuthmiddlewareController) AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		path := ctx.FullPath()
		fmt.Println("Path", path)
		if path == "/login" {
			ctx.Next()
			return
		}

		var SessionId string
		DB := am.db.IntializeDB()
		query := `SELECT SessionId FROM UserAccount WHERE UserName=$1`
		Authorization := ctx.GetHeader("Authorization")
		Authtoken := string(Authorization)
		log.Println("Auth Header", Authorization)
		if len(Authorization) == 0 {
			ctx.JSON(http.StatusBadRequest, gin.H{"Authorization": Authorization, "message": "Check Auth token"})
			ctx.Abort()
			return
		}
		payload, err := am.pl.VerifyToken(Authtoken)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"Error ": err, "payload": payload})
			ctx.Abort()
			return
		}
		Data := DB.QueryRow(query, payload.UserName)
		Data.Scan(&SessionId)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"Error ": err})
			ctx.Abort()
			return
		}
		id, _ := uuid.Parse(SessionId)
		fmt.Println("SessionId uuid from db", id)

		if payload.ID != id {
			ctx.Abort()
			return
		}
		log.Println("Authmiddleware passed")
		ctx.Next()

	}
}
