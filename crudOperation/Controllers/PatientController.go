package controllers

import (
	"CurdOperation/Model"
	services "CurdOperation/Services"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PatientController struct {
	patientservice services.IpatientService
}

func PatientCtor(PS services.IpatientService) *PatientController {
	return &PatientController{patientservice: PS}
}

func (Patient *PatientController) GetPatient(ctx *gin.Context) {
	Prn := ctx.Param("PRN")
	fmt.Println("PRN ", Prn)
	//data := services.GetKeyValue(Prn)
	data := Patient.patientservice.GetPatient(Prn)
	//data := services.GetPatient(Prn)
	ctx.JSON(http.StatusOK, gin.H{"data": data})
}
func (Patient *PatientController) AddPatient(ctx *gin.Context) {
	var req *Model.Patient
	if err := ctx.ShouldBindJSON(&req); err != nil {
		panic(err)
	}
	fmt.Println(req)
	//services.AddPatient(*req)
	Patient.patientservice.AddPatient(*req)
	ctx.JSON(http.StatusOK, gin.H{"data": req})
}
func (Patient *PatientController) UpdatePatient(ctx *gin.Context) {
	var req *Model.Patient
	if err := ctx.ShouldBindJSON(&req); err != nil {
		panic(err)
	}
	fmt.Println(req)
	Patient.patientservice.UpdatePatient(*req)
	ctx.JSON(http.StatusOK, gin.H{"data": req})
}
func (Patient *PatientController) DeletePatient(ctx *gin.Context) {
	PRN := ctx.Param("PRN")
	fmt.Print(PRN)
	//Patient.patientService.DeletePatient(PRN)
	ctx.JSON(http.StatusOK, gin.H{"data": PRN})
}
