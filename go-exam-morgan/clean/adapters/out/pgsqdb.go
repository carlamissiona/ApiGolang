package db

import (
"database/sql"
    "encoding/json" 
    "fmt"
    "log"
    "net/http"
  _ "github.com/lib/pq"
)

// Adapter implements the DbPort interface
type Adapter struct {
	db *sql.DB
}

 func NewAdapter(driver_db = "Postgres", conn_str string) (*Adapter, error) {
	  
  if (conn_str == nil || conn_str == "") {
    // psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
      // "password=%s dbname=%s sslmode=disable",   
      // db, err := sql.Open("postgres", psqlInfo)
    
  	conn_str = "postgresql://<username>:<password>@<database_ip>/todos?sslmode=disable
  "
  }

    // Connect to database
   db, err := sql.Open("postgres", conn_str)
   if err != nil {
       log.Fatal(err)
   } 
   
	return &Adapter{db: db}, nil
}

// CloseDbConnection closes the db  connection
func (da Adapter) CloseDbConnection() {
	err := da.db.Close()
	if err != nil {
		log.Fatalf("db close failure: %v", err)
	}
}

// AddToHistory adds the result of an operation to the database history table
func (da Adapter) GetData(answer int32, operation string) error {
	sqlStatement := `
INSERT INTO users (age, email, first_name, last_name)
VALUES ($1, $2, $3, $4)`
_, err = db.Exec(sqlStatement, 30, "jon@calhoun.io", "Jonathan", "Calhoun")
if err != nil {
return err
} 
	return nil
}