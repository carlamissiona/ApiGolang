package in

    import (
        "net/http"
        "log"
        "github.com/gin-gonic/gin"

        _ "gomorganexam/clean/ports"
        "gomorganexam/clean/adapters/out/db"
    )

var Server * gin.Engine


    // db_instance *sql.DB
        // db.Exec()   NOTE *db.Database_Instance  errored

    // log.Println( ports.NewAPIService( db_instance  ) )

func NewServer(db_adapter db.Adapter) * gin.Engine {
    Server = gin.Default()
     

    Server.GET("/", func(c * gin.Context) {
        if (CheckInitData(db_adapter)) {
            log.Println("Ok")
        }
        c.String(http.StatusOK, "hello world  port")


    })

    return Server

}
func Start() {
    Server.Run()
}
func CheckInitData(db_adapter db.Adapter) bool {
    obs, err := db_adapter.Db.Query(`SELECT * FROM tbl_countrypatient`)
    log.Println(obs); log.Println(err)
    return true
}