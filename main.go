package main

//Necessary packages are imported here.
import (
	"database/sql"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

// This block of code declares constant variables for the database connection parameters
const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "0919"
	dbname   = "mydb"
)

// Main function
func main() {
	//To format a string connection.
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		panic(err)
	}
	//To create the driver instance with the necessary arguments.
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		panic(err)
	}
	//To create the source instance with the necessary arguments.
	source, err := (&file.File{}).Open("db/migration")
	if err != nil {
		panic(err)
	}
	//It will create the new migration instance.
	m, err := migrate.NewWithInstance("file", source, "postgres", driver)
	if err != nil {
		panic(err)
	}

	// Apply schema migrations for the database.
	if err := m.Up(); err != nil {
		panic(err)
	} else {
		fmt.Println("Schema migrations applied successfully.")
	}

	// Apply data migrations for inserting the row into the migrated database.
	if err := m.Up(); err != nil {
		panic(err)
	} else {
		fmt.Println("Data migrations for insertion applied successfully.")
	}

	// Apply data migrations for update into the migrated database.
	if err := m.Up(); err != nil {
		panic(err)
	} else {
		fmt.Println("Data migrations for update applied successfully.")
	}

	// Execute the update query
	if _, err := db.Exec(`UPDATE "Student1" SET "Name" = 'Magesh' WHERE "Rollno" =1`); err != nil {
		panic(err)
	}

	fmt.Println("Update query succeeded.")
}
