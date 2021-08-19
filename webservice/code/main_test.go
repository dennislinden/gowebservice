package main

import (
	"log"
	"testing"
)

// INSERT TEST
// given a database connection
// when inserting personal data
// and retreiving personal data
// then retreived personal data should be the same as inputed data

func TestInsertPersonalData(t *testing.T) {
	connectToDatabase()
	insertToDatabase("Dennis", 63, 8721, "laan")
	name, age, number, address := getPersonalData()
	if name != "Dennis" {
		t.Error("Name did not match")
	}
	if age != 63 {
		t.Error("Age did not match")
	}
	if number != 8721 {
		t.Error("Number did not match")
	}
	if address != "laan" {
		t.Error("Address did not match")
	}

}

// DATABASE DATA TEST
// given a database connection
// when asking for all data
// count the data in the db
// show all data and check if the total amount of data is matching

func TestRetreivingAllDatabaseData(t *testing.T) {
	connectToDatabase()
	countTotalDataTable()
	totalData, err := db.Exec("SELECT COUNT(*) FROM TEST")

	if err != nil {
		log.Fatal(err)
	}
	if totalData != 5 {
		t.Error("Numbers don't match")
	}

}
