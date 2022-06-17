package in

    import (
        // "net/http"
        // "log"
        "github.com/gin-gonic/gin"

        // "gomorganexam/clean/adapters/out/db"
    )

var Server * gin.Engine

type WebServerAdapter interface { 
	CheckServer()
  Start()
  
} 


type WebAdapter struct {
	HttpEngine     *gin.Engine
}

func NewServer() (*WebAdapter){
    Server = gin.Default()
  

   

   return &WebAdapter {
        HttpEngine: Server,
    }

}
func (h *WebAdapter)Start() {
     h.HttpEngine.Run()
}

// err := db.QueryRow("SELECT COUNT(*) FROM main_table").Scan(&count)
// switch {    
// case err != nil:
//     log.Fatal(err)
// default:
//     fmt.Printf("Number of rows are %s\n", count)
// }