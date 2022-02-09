package main

import (
    "database/sql"
    "fmt"
    "log"
    "os"       // used to read the environment variable
    //"github.com/joho/godotenv" // package used to read the .env file
    _ "github.com/lib/pq"      // postgres golang driver
)

// response format
type response struct {
    ID      int64  `json:"id,omitempty"`
    Message string `json:"message,omitempty"`
}

// create connection with postgres db
func createConnection() *sql.DB {
    // load .env file
    //err := godotenv.Load(".env")

    //if err != nil {
    //    log.Fatalf("Error loading .env file")
    //}

    // Open the connection
    db, err := sql.Open("postgres", os.Getenv("POSTGRES_URL"))

    if err != nil {
        panic(err)
    }

    // check the connection
    err = db.Ping()

    if err != nil {
        panic(err)
    }

    fmt.Println("Successfully connected!")
    // return the connection
    return db
}

//------------------------- handler functions ----------------
func main() {

    // create the postgres db connection
    db := createConnection()

    // close the db connection
    defer db.Close()

    // create the insert sql query
    // returning userid will return the id of the inserted user
    // execute the sql statement
    res, err := db.Exec(fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (userid SERIAL PRIMARY KEY, name TEXT, age INT, location TEXT)", "users"))
    if err != nil {
      log.Fatal(err)
    }
    fmt.Println("Table Created!:",res)
}
