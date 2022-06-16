package main

import (
	"log"
	"os"
	"github.com/joho/godotenv"
	// application
 

	// adapters
   go-exam-morgan/adapters/postgres
)

func main() { 
 
 
    program_attr := os.Args[0]
    if (program_attr == "minimal"){
        // Postgres + Fiber 

        // get os postgres 
        // env config

        err := godotenv.Load(".environ_development")
	      if err != nil {
      		log.Fatalf("failed reading env file: %v", err)
      	}

    	 
        db_info := os.Getenv("POSTGRES_URL") 
       
  	    defer db.Close()

    }
    else{
       // Postgres + Fiber + Apigateway
       
    }
     
}




  
	var err error

	dbaseDriver := os.Getenv("DB_DRIVER")
	dsourceName := os.Getenv("DS_NAME")

	dbAdapter, err := db.NewAdapter(dbaseDriver, dsourceName)
	if err != nil {
		log.Fatalf("failed to initiate dbase connection: %v", err)
	}
	defer dbAdapter.CloseDbConnection()

	// core
	core := arithmetic.New()

	// NOTE: The application's right side port for driven
	// adapters, in this case, a db adapter.
	// Therefore the type for the dbAdapter parameter
	// that is to be injected into the NewApplication will
	// be of type DbPort
	applicationAPI := api.NewApplication(dbAdapter, core)

	
}
