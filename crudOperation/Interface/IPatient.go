package Interface

import "CurdOperation/Model"

type Ipatient interface {
	AddPatient(patient Model.Patient) error
	GetPatient(name string) (Model.Patient, error)
	UpdatePatient(name string) (string, error)
	DeletePatient(name string)
}
