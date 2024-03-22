package services

import (
	"CurdOperation/Model"
	"database/sql"
	"fmt"
	"log"

	_ "modernc.org/sqlite"
)

func IntializeSqliteDb() *sql.DB {

	SqliteDb, err := sql.Open("sqlite", "KeyvalueSqlite3db.db")

	if err != nil {
		log.Fatalf("error opening SQLite database: %v", err)
	}

	Query := `CREATE TABLE IF NOT EXISTS keyvalue (id INTEGER PRIMARY KEY AUTOINCREMENT,key TEXT NOT NULL,value TEXT)`
	//	CREATE TABLE IF NOT EXISTS keyvalue (id INTEGER PRIMARY KEY AUTOINCREMENT, key TEXT NOT NULL, value TEXT)

	_, err = SqliteDb.Exec(Query)
	CheckErr(err)

	return SqliteDb
}

func InsetKeyValue(key string, value string) {
	db := IntializeSqliteDb()
	query := `INSERT INTO keyvalue (key,value) VALUES($1,$2)`
	_, err := db.Exec(query, key, value)
	CheckErr(err)
	fmt.Println("Key-Value added to keyvalue ")

}

func GetKeyValue(key string) string {
	fmt.Println("key :", key)
	db := IntializeSqliteDb()
	//InsetKeyValue("ApiKey", "123456789")
	query := `SELECT * FROM keyvalue WHERE key=?;`
	row, err := db.Query(query, key)
	CheckErr(err)
	var keyvalue Model.SqliteDb

	for row.Next() {
		err := row.Scan(&keyvalue.Id, &keyvalue.Key, &keyvalue.Value)
		CheckErr(err)
	}
	fmt.Println("value :", keyvalue.Value)
	return keyvalue.Value

}

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}
