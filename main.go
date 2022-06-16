package main

import ( 
	"log"
	"os"
	_"github.com/joho/godotenv"
  _ "github.com/lib/pq"
  _"database/sql"
	// application
  
  "gomorganexam/clean/adapters/in" 
	// adapters
   // go-exam-morgan/adapters/postgres
)

 

 

func main() { 
  
    log.Println("Empty param ",len( os.Args)) 
    if (len(os.Args)  == 2 ){
            if (os.Args[1] == "minimal"){
              
                // minimal using gofiber + postgres + openAPI
         
                // get os env postgres 
              web_adapter := in.NewServer() 

              
                     
            }
      }
      log.Println("Hi Postg API")
}





 