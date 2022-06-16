package in

import (
    "net/http"
"log"
    "github.com/gin-gonic/gin"
  
    "gomorganexam/clean/ports" 
   "gomorganexam/clean/adapters/out/db" 
)


func NewServer() {
    r := gin.Default()
    
     // db.Exec()
   
    log.Println( ports.NewAPIService(db.Pg_DB   ) )

    r.GET("/", func(c *gin.Context) {
         if( CheckInitData() ) {
      log.Println("Ok")
 }  
      c.String(http.StatusOK, "hello world")

      
    })

	 r.Run()
  
}
func CheckInitData() bool {
   _ = db.Pg_Execute(`SELECT * FROM TBL_PATIENTS`)
}