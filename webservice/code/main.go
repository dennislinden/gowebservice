package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var tpl *template.Template

//struct for database
type Tag struct {
	Name    sql.NullString
	Age     sql.NullInt64
	Number  sql.NullInt64
	Address sql.NullString
}

//struct for website data
type Sub struct {
	Name    string
	Age     string
	Number  string
	Address string
}

func connectToDatabase() {
	// Get a database handle.
	var err error
	db, err = sql.Open("mysql", "root:my-password@tcp(127.0.0.1:3306)/users")
	if err != nil {
		log.Fatal(err)
	}

	pingerr := db.Ping()
	if pingerr != nil {
		log.Fatal(pingerr)
	}
}

func generateTableDatabase(table string) {
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS test (users TEXT NOT NULL, age INT NOT NULL, number INT NOT NULL, address TEXT NOT NULL)")
	if err != nil {
		panic(err)
	}
	fmt.Println("Tables added")
}

func insertToDatabase(name string, age int, number int, address string) {
	// perform a db.Query insert
	insert, err := db.Query("INSERT INTO test VALUES(?, ?, ?, ?)", name, age, number, address)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Data added!")
	// be careful deferring Queries if you are using transactions
	defer insert.Close()
}

func countTotalDataTable() {
	rowData := db.QueryRow("SELECT COUNT(*) FROM test") //Query row returns one row of DB data. DB always returns data in a row also the count of rows.
	totalRowSize := 0                                   // make a variable to store the result of Scan as an int

	err := rowData.Scan(&totalRowSize) //scan and checks for errors. If there are no errors Print totalRowSize
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(totalRowSize)
}

func getPersonalData() (string, int, int, string) {
	name := "Dennis"
	age := 63
	number := 8721
	address := "laan"

	return name, age, number, address
}

func executeQuery() {
	// Execute the query
	results, err := db.Query("SELECT users, age, number, address  FROM test")
	if err != nil {
		panic(err.Error()) // Check if there are errors
	}

	for results.Next() {
		var tag Tag
		// for each row, scan the result into our tag composite object
		err = results.Scan(&tag.Name, &tag.Age, &tag.Number, &tag.Address)
		if err != nil {
			panic(err.Error())
		}
		// and then print out the tag's Name attribute
		log.Println(tag.Name, tag.Age, tag.Number, tag.Address)
	}
}

func processGetHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("processGetHandler running")
	if r.Method == http.MethodGet {
		readfile, err := ioutil.ReadFile("templates/getform.html")
		if err != nil {
			panic(err)
		}
		w.Write(readfile)
	}

	if r.Method == http.MethodPost {
		var s Sub                        // Variable to struct Sub
		s.Name = r.FormValue("userName") // FormValue only outputs strings
		s.Age = r.FormValue("ageValue")
		s.Number = r.FormValue("numberValue")
		s.Address = r.FormValue("addressValue")
		//fmt.Fprintf(w, "<h1>%s %s %s %s</h1>", s.Name, s.Age, s.Number, s.Address)
		tpl.ExecuteTemplate(w, "added_data.html", s)

		//convert string to int
		ageConvert, err := strconv.Atoi(s.Age)
		if err != nil {
			fmt.Println(err)
		}

		//convert string to int
		numberConvert, err := strconv.Atoi(s.Number)
		if err != nil {
			fmt.Println(err)
		}
		insertToDatabase(string(s.Name), ageConvert, numberConvert, s.Address)
	}

}

func main() {
	connectToDatabase()
	//generateTableDatabase("test")
	//executeQuery()
	countTotalDataTable()
	//tpl, _ = tpl.ParseGlob("templates/*.html")

	//http.HandleFunc("/processget", processGetHandler)
	//http.ListenAndServe(":8081", nil) //startup html server
}
