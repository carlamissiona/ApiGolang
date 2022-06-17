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
    // obs, err := db_adapter.Db.QueryRow(`SELECT Count(*) FROM tbl_countrypatient`).Scan(&count)
    var count int32
    err := db_adapter.Db.QueryRow(`SELECT Count(*) FROM tbl_countrypatient`).Scan( & count)

    switch {
        case err != nil:
            log.Printf("Error CheckInitData")
            log.Fatal(err)
        default:
            log.Printf("Number of rows are %s\n", count)
    }
    // log.Println("=====obs====");

    //   log.Println(obs); log.Println(err)
    //   log.Println("=====obs====");
    //   if ( obs == nil ){
    //      log.Println("=====empty obs====");log.Println("=====empty obs====");
    //   }else{
    //      log.Println("=====obs has values====");
    //   }
    return true
}


// err := db.QueryRow("SELECT COUNT(*) FROM main_table").Scan(&count)
// switch {    
// case err != nil:
//     log.Fatal(err)
// default:
//     fmt.Printf("Number of rows are %s\n", count)
// }