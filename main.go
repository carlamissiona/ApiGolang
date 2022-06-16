package main

import ( 
	"log"
	"os"
	_"github.com/joho/godotenv"
  _ "github.com/lib/pq"
  _"database/sql"
	// application
 
 "gomorganexam/ports" 
  "gomorganexam/adapters" 
	// adapters
   // go-exam-morgan/adapters/postgres
)

 

 type ApplicationAdapter interface {
  Type string
  Adapter interface{}
}


func main() { 
  
    log.Println("Empty param ",len( os.Args)) 
    if (len(os.Args)  == 2 ){
            if (os.Args[1] == "minimal"){
              
                // minimal using gofiber + postgres + openAPI
         
                // get os env postgres 
              
             
                     
            }
      }
      log.Println("Hi Postg API")
}





 