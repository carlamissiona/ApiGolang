package db

import (
  "database/sql"
    "encoding/json" 
    "fmt"
    "log" 
    "net/http"
  _ "github.com/lib/pq"
)

 type Adapter struct {
	db *sql.DB
}

 func NewAdapter(driver_db = "Postgres", conn_str string) (*Adapter, error) {
	  
      err := godotenv.Load(".environ_development")
      if err != nil {
        log.Fatalf("failed reading env file: %v", err)
      }

      db_info := os.Getenv("POSTGRES_URL") 



    
      //  db pg
    
      db, err := sql.Open("postgres", db_info)
      if err != nil {
        panic(err)
      }
     
      defer db.Close()
    
      err = db.Ping()
      if err != nil {
        log.Fatalf("failed No DB connection %v", err)
      }
   
}

// // AddToHistory adds the result of an operation to the database history table
// func (da Adapter) GetData(answer int32, operation string) error {
// 	sqlStatement := `
// INSERT INTO users (age, email, first_name, last_name)
// VALUES ($1, $2, $3, $4)`
// _, err = db.Exec(sqlStatement, 30, "jon@calhoun.io", "Jonathan", "Calhoun")
// if err != nil {
// return err
// } 
// 	return nil
// }