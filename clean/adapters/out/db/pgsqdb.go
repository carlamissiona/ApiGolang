package db

import (
  "database/sql"
    _"encoding/json" 
    _"fmt"
  _ "os" 
    "log" 
   _ "net/http"
  _"github.com/joho/godotenv"
  _ "github.com/lib/pq"
)

type DBAdapter interface{
 NewAdapter(driver_db, conn_str string) (*Adapter)
 Desc_DB() *sql.DB 
 Query(sqlStatement string) *sql.Rows
 Execute(sqlStatement string) sql.Result
}

type Adapter struct {
	db      *sql.DB
  driver  string
}
var Database_Instance *sql.DB
func NewAdapter(driver_db, conn_str string) (*Adapter) {
	  
    
      //  db pg
      Database_Instance, err := sql.Open(driver_db, conn_str)
      if err != nil {
        panic(err)
      }
      
      defer Database_Instance.Close()
    
      err = Database_Instance.Ping()
      if err != nil {
        log.Fatalf("failed No DB connection %v", err)
      }
      // cCheck Initial Data
     	 
    return &Adapter{
      db: Database_Instance,
    }
    
}  

func Desc_DB() *sql.DB{
  return Database_Instance
  
}
 
func (a *Adapter)Query(sqlStatement string) *sql.Rows {
    var rows *sql.Rows = nil
  
    log.Println("Database_Instance")
    log.Println(Database_Instance)
  // sqlStatement := `SELECT * FROM TBL_PATIENTS)`
      rows, err := Database_Instance.Query(sqlStatement)
      if err != nil {
            log.Println("error ->", err)
      } 
  return rows
}
func (a *Adapter)Execute(sqlStatement string) sql.Result {
       var rows sql.Result = nil
  // sqlStatement := `SELECT * FROM TBL_PATIENTS)`
      rows, err := Database_Instance.Exec(sqlStatement)
      if err != nil {
           log.Println("error ->", err)
        
      } 
   return rows
}