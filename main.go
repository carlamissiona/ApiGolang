package main

import (
    "log"
    "os"
    _ "net/http"
    _ "github.com/joho/godotenv"
    _ "github.com/lib/pq"
    _ "database/sql"

    _ "github.com/gin-gonic/gin"

    "gomorganexam/clean/adapters/in"
    "gomorganexam/clean/adapters/out/db"
    "github.com/joho/godotenv"
)




func main() {


    log.Println("Empty param ", len(os.Args))
    log.Println("Empty param ", len(os.Args))
    if (len(os.Args) == 2) {
        log.Println(os.Args)
        log.Println(os.Args[1])
        if (os.Args[1] == "minimal") {
            log.Println(os.Args[1])
              
            err := godotenv.Load(".environ_development")
            if err != nil {
                log.Fatalf("failed reading env file: %v", err)
            }

            db_info:= os.Getenv("POSTGRES_URL")
            DBAD := * db.NewAdapter("postgres", db_info)
                // get os env postgres 
            server := in.NewServer(DBAD)
            log.Println(server) 
            in.Start()


        }
    }
    log.Println("Hi Postg API")
}