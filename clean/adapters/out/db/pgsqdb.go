package db

import (
  "database/sql"
    "encoding/json" 
    "fmt"
  "os"
    "log" 
   _ "net/http"
  "github.com/joho/godotenv"
  _ "github.com/lib/pq"
)

 type Adapter struct {
	db *sql.DB
}
var Database_Instance *sql.DB
 func NewAdapter(driver_db, conn_str string) (*Adapter, error) {
	  
      err := godotenv.Load(".environ_development")
      if err != nil {
        log.Fatalf("failed reading env file: %v", err)
      }

      db_info := os.Getenv("POSTGRES_URL") 



    
      //  db pg
      Database_Instance
      , err := sql.Open("postgres", db_info)
      if err != nil {
        panic(err)
      }
      
      defer .Close()
    
      err = db.Ping()
      if err != nil {
        log.Fatalf("failed No DB connection %v", err)
      }
      // cCheck Initial Data
     	 
    
   
} 
func Pg_DB() *sql.DB{
 return Database_Instance
  
}

func Pg_Execute(sqlStatement string) *sql.Rows {
     
  // sqlStatement := `SELECT * FROM TBL_PATIENTS)`
      rows, err := db.Query(sqlStatement)
      if err != nil {
          return err
      } 
  return rows
}
func Pg_ExecuteNoResult(sqlStatement string) *sql.Rows {
     
  // sqlStatement := `SELECT * FROM TBL_PATIENTS)`
      rows, err := db.Exec(sqlStatement)
      if err != nil {
          return err
      } 
  return rows
}