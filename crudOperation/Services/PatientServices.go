package services

import (
	"CurdOperation/Model"

	"fmt"
)

type IpatientService interface {
	AddPatient(pa Model.Patient) error
	GetPatient(PRN string) []Model.Patient
	UpdatePatient(pa Model.Patient)
	DeletePatient(PRN string) string
}
type PatientserviceStruct struct {
	db Idb
}

func PatientserviceCtor(db Idb) *PatientserviceStruct {
	return &PatientserviceStruct{
		db: db,
	}
}

func (PS *PatientserviceStruct) AddPatient(pa Model.Patient) error {
	db := PS.db.IntializeDB()
	//db := IntializeDB()

	defer db.Close()
	fmt.Println("Addpatient service")
	query := `INSERT INTO patienttable (firstname, PRN, lastname, address, email, phonenumber, pincode) 
	VALUES ($1, $2, $3, $4, $5, $6, $7);`

	_, err := db.Exec(query, pa.FirstName, pa.PRN, pa.LastName, pa.Address, pa.Email, pa.PhoneNumber, pa.Pincode)
	return err
}
func (PS *PatientserviceStruct) UpdatePatient(pa Model.Patient) {
	db := PS.db.IntializeDB()
	//db := IntializeDB()
	defer db.Close()
	selectquery := `update patienttable set firstname=$1, lastname=$2, address=$3, email=$4, phonenumber=$5, pincode=$6 where PRN=$7;`
	_, err := db.Exec(selectquery, pa.FirstName, pa.LastName, pa.Address, pa.Email, pa.PhoneNumber, pa.Pincode, pa.PRN)
	if err != nil {
		panic(err)
	}
	fmt.Println("Update patient successful")

}
func (PS *PatientserviceStruct) DeletePatient(PRN string) string {
	db := PS.db.IntializeDB()
	//db := IntializeDB()
	defer db.Close()
	query := `delete from patienttable where PRN = $1;`
	_, err := db.Exec(query, PRN)
	if err != nil {
		panic(err)
	}
	return PRN + "Deleted"
}

func (PS *PatientserviceStruct) GetPatient(PRN string) []Model.Patient {
	db := PS.db.IntializeDB()
	// db := IntializeDB()
	defer db.Close()
	fmt.Println("GetPatient")
	var Patients []Model.Patient
	query := `SELECT * FROM patienttable`
	Data, err := db.Query(query)
	if err != nil {
		panic(err)
	}
	for Data.Next() {
		var patient Model.Patient

		err := Data.Scan(&patient.FirstName, &patient.PRN, &patient.LastName, &patient.Address, &patient.Email, &patient.PhoneNumber, &patient.Pincode)
		if err != nil {
			panic(err)
		}
		Patients = append(Patients, patient)
	}
	fmt.Println("PRN  from GET", PRN)
	fmt.Println("data", Patients)
	if err != nil {
		panic(err)
	}
	return Patients

}
