package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	db, err := sql.Open(
		"postgres",
		"postgres://tut_user:tut_user@localhost:5432/golang_tut?sslmode=disable",
	)

	if err != nil {
		log.Fatal(err)

	} else {
		// defer db.Close()
		fmt.Println("Database connection successful")

	}

	connectivity := db.Ping()
	if connectivity != nil {
		log.Fatal(connectivity)
	} else {
		fmt.Println("Good to go")
	}

	fmt.Println("Creating table ...")
	createTable(db)
	fmt.Println("Inserting data into table ...")
	InsertIntoTable(db)

	db.Close()

	// GORM
	UsingGORM()
}

func createTable(db *sql.DB) {
	DBCreate := `
	CREATE TABLE IF NOT EXISTS staffs (
		id SERIAL PRIMARY KEY,
		name character varying(100) COLLATE pg_catalog."default",
		age integer,
		department character varying(100) COLLATE pg_catalog."default"
	)
	WITH (
		OIDS = FALSE
	);
	`

	_, err := db.Exec(DBCreate)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Table created successfully")
	}
}

func InsertIntoTable(db *sql.DB) {
	insert, err := db.Prepare("INSERT INTO staffs(name, age, department) VALUES($1, $2, $3)")

	if err != nil {
		log.Fatal(err)
	}

	_, err = insert.Exec("John Doe", 30, "IT")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted data successfully")
}

type User struct {
	gorm.Model
	FirstName string
	LastName  string
	Age       int
	Email     string
}

func UsingGORM() {
	connection_string := "host=localhost user=tut_user password=tut_user dbname=golang_tut port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(connection_string), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&User{})

	new_user_1 := &User{
		FirstName: "John",
		LastName:  "Doe",
		Age:       30,
		Email:     "john.doe@dev.google.com",
	}
	db.Create(new_user_1)

	new_user_2 := &User{
		FirstName: "Jane",
		LastName:  "Smith",
		Age:       25,
		Email:     "jane.smith@dev.waymo.com",
	}
	db.Create(new_user_2)

	var user User
	tx := db.First(&user, &User{FirstName: "John"})
	if tx.Error != nil {
		log.Fatal(tx.Error)
	} else {
		fmt.Println(user)
	}

}
