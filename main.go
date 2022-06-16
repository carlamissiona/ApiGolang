package main

import ( 
	"log"
	"os"
	"github.com/joho/godotenv"
  _ "github.com/lib/pq"
  "database/sql"
	// application
 

	// adapters
   // go-exam-morgan/adapters/postgres
)

func main() { 
 
 
    program_attr := os.Args[0]
    if (program_attr == "minimal"){
        // minimal using gofiber + postgres + openAPI

        // get os env postgres 
    

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
      		panic(err)
      	}

             
    }
     
}





 