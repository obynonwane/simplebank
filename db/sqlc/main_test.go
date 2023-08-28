package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:gobank@localhost:5432/simple_bank?sslmode=disable"
)

// global variable - hold reference to a set of prepared sql queries
var testQueries *Queries

// this is a special function in go testing framework that is called be	fore any test is ran

func TestMain(m *testing.M) {

	//open  a new database connection
	conn, err := sql.Open(dbDriver, dbSource)

	//check if err exist while creating connection
	if err != nil {
		log.Fatal("can not connect to db:", err)
	}

	//extablish connection
	testQueries = New(conn)

	//exit the test once finished running returning a status code 0 for succesful or otherwise
	os.Exit(m.Run())

}
