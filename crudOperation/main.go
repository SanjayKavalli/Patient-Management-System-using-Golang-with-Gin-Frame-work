package main

import (
	controllers "CurdOperation/Controllers"
	middleware "CurdOperation/Middleware"
	services "CurdOperation/Services"

	"github.com/gin-gonic/gin"
)

func main() {
	// Create Gin router
	router := gin.Default()
	//Intialize services
	db := services.DbCtor()
	ReJson := services.ReadJsonCtor()
	patientService := services.PatientserviceCtor(db)
	payload := services.Payloadctor(db)
	useraccount := services.UserAccountCtor(db, payload, ReJson)
	//Intialise contollers
	patientController := controllers.PatientCtor(patientService)
	signupcontroller := controllers.SignUpCtor(useraccount, payload)

	//Middleware
	AuthMiddleware := middleware.AuthCtor(db, payload)
	router.Use(middleware.Apikeyvalidator())

	router.Use(AuthMiddleware.AuthMiddleware())
	// Register Routes
	router.POST("/login", signupcontroller.Login)
	router.POST("/SignUp", signupcontroller.SignUp)

	router.GET("/GetPatient/:PRN", patientController.GetPatient)
	router.POST("/DeletePatient/:PRN", patientController.DeletePatient)
	router.POST("/AddPatient", patientController.AddPatient)
	router.PUT("/UpdatePatient", patientController.UpdatePatient)

	//router.POST("/login", signupcontroller.GenerateToken)
	router.POST("/VerifyLoginToken", signupcontroller.VerifyLoginToken)

	// Start the server
	router.Run()
}
