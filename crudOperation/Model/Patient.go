package Model

type Patient struct {
	PRN         int    `json:"PRN"`
	FirstName   string `json:"firstname" validate:"required"`
	LastName    string `json:"lastname" validate:"required"`
	Address     string `json:"address" validate:"required"`
	Email       string `json:"email" validate:"required,email"`
	PhoneNumber string `json:"phone_number" validate:"required"`
	Pincode     *int   `json:"pincode,omitempty" validate:"omitempty,len=6"`
}

type SqliteDb struct {
	Id    int
	Key   string
	Value string
}
