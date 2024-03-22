package services

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Idb interface {
	IntializeDB() *sql.DB
}
type DBController struct{}

// IntializeDB implements Idb.

func DbCtor() *DBController {
	return &DBController{}
}

// const (
// 	host     = "localhost"
// 	port     = 5432
// 	user     = "postgres"
// 	password = ""
// 	dbname   = "Crud_Operation"
// )

func (Db *DBController) IntializeDB() *sql.DB {
	// db, err := sql.Open("postgres", "postgres://user:password@localhost:5432/dbname?sslmode=disable")
	// if err != nil {
	//     return err
	// }
	// defer db.Close()
	db, err := sql.Open("postgres", "user=postgres dbname=Crud_Operation password=123456789 sslmode=disable")

	//connectionstring := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	//db, err := sql.Open("postgres", connectionstring)
	Checkerr(err)
	fmt.Println("DB intialized successfully")

	return db

}

func Checkerr(err error) {
	if err != nil {
		panic(err)

	}
}
