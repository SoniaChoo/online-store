package test

import (
	db2 "github.com/SoniaChoo/online-store/db"
	"io/ioutil"
	"log"
	"os"
	"testing"
)

const DatabaseName = "DB_NAME"

func TestMain(m *testing.M) {
	// change environment variable to use test database
	testDB := "test_store"
	originalDB := os.Getenv(DatabaseName)
	if originalDB == "" { // error
		log.Fatalf("fail to read environment variable %s, will exit all tests!", DatabaseName)
		os.Exit(1)
	}
	if err := os.Setenv(DatabaseName, testDB); err != nil {
		log.Fatalf("set environment variable failed! %s should be set to %s, but failed!", DatabaseName, testDB)
	}

	// create test database
	sqlBytes, err := ioutil.ReadFile("zzstore.sql")
	if err != nil {
		log.Fatalf("fail to read sql file, will exit all tests!")
		os.Exit(1)
	}

	sqlTable := string(sqlBytes)
	//fmt.Println(sqlTable)
	db, err := db2.DBFactory()
	if err != nil {
		log.Fatalf("fail to connect database, will exit all tests!")
		os.Exit(1)
	}
	defer db.Close()

	if _, err = db.Exec(sqlTable); err != nil {
		log.Fatalf("fail to init test database, will exit all tests!")
		os.Exit(1)
	}

	// start test
	exitCode := m.Run()

	// delete test database
	if _, err = db.Exec("DROP TABLE IF EXISTS `" + testDB + "`;"); err != nil {
		log.Fatalf("fail to delete test database")
		os.Exit(1)
	}

	// revert the environment change
	if err := os.Setenv(DatabaseName, originalDB); err != nil {
		log.Fatalf("revert environment variable failed! %s should be revert to %s, but failed!", DatabaseName, originalDB)
	}

	os.Exit(exitCode)
}
