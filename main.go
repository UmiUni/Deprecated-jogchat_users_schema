package main

import (
	"database/sql"
	"net/http"
	"log"
	_ "github.com/lib/pq"
)

const hashCost = 8
var (
  db *sql.DB
  var connBuffer bytes.Buffer
  user := "postgres"
  password := "umiuni_jogchat_tiantian"
  ip := "206.189.212.172:5432"
  dbname := "jogchat"
  sslmode := "verify-full"

)

func main() {
	// "Signin" and "Signup" are handler that we will implement
	http.HandleFunc("/signin", Signin)
	http.HandleFunc("/signup", Signup)
	// initialize our database connection
	initDB()
	// start the server on port 8000
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func initDB(){
	var err error
	// Connect to the postgres db
        // why we use buffer to concatenate: http://herman.asia/efficient-string-concatenation-in-go
        connStr := "postgres://pqgotest:password@localhost/pqgotest?sslmode=verify-full"
        connBuffer.writeString("postgres://") 
        connBuffer.writeString(user) 
        connBuffer.writeString(password) 
        connBuffer.writeString("@") 
        connBuffer.writeString(ip) 
        connBuffer.writeString("/") 
        connBuffer.writeString(dbname) 
        connBuffer.writeString("?sslmode=") 
        connBuffer.writeString(sslmode) 
       
        db, err := sql.Open("postgres", connBuffer.String())

	if err != nil {
		panic(err)
	}
}
