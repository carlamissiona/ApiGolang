package ports

import (
    "gomorganexam/clean/core/domain"

    "gomorganexam/clean/adapters/in"
    "gomorganexam/clean/adapters/out/db"
 "github.com/gin-gonic/gin"
    // _"database/sql"
    "log"
    "os"
   "net/http"
    "encoding/csv"
)


 
type ServicePatientsAPI interface {
    GetTopConfirmed(observation_date string, max_result int32) domain.PatientObservation
    GetAllConfirmed(pagelimit int32) domain.PatientObservation
    StartSvc(hasRecords bool)
}

type Service struct {
    Data *db.Adapter
    Records[] byte
    Server *in.WebAdapter

}
var InstanceAPI *Service
func NewAPIService(data *db.Adapter, records[] byte, engine *in.WebAdapter)( *Service) {
     InstanceAPI =  &Service {
        Data: data,
        Records: nil,
        Server: engine,
    }
 return InstanceAPI
}
func(svc *Service) Start(hasRecords bool) {

    if !hasRecords {

        csv_ToPostgres()


    }
    svc.Server.Start()
}

func(svc *Service) GetTopConfirmed(observation_date string, max_result int32) {
    // svc.Database 


} 


func(svc *Service) GetAllConfirmed(pagelimit int32) {
    log.Println("No Implementation")
}

func csv_ToPostgres() {


    csvfile, err := os.Open("data.csv")
    if err != nil {
        log.Printf("=>ERROR!! NO CSV FILE \n %v \n=>END OF ERROR!! NO CSV FILE " , err ); 
        handl_err_nocsv()
    }

    // remember to close the file at the end of the program
    defer csvfile.Close()
 
    // read csv values using csv.Reader
    csv_reader := csv.NewReader(csvfile)

    data, err := csv_reader.ReadAll()
    if err != nil {
        log.Printf("=>ERROR!! CANT READ CSV FILE \n %v \n=>END OF ERROR!! NO CSV FILE " , err ); 
        
        handl_err_readcsv()
    }
    log.Println("API data")
    log.Printf("=>DEBUG!! API DATA")
    log.Println(data)
    log.Printf("=>END OF DEBUG!! API DATA")

}
 
func handl_err_nocsv() {
 InstanceAPI.Server.HttpEngine.GET("/", func(c *gin.Context) {
        // if (CheckInitData(db_adapter)) {
            
        // }
        c.String(http.StatusOK, "No Routes Created \nPls. Input CSV File To Continue. ")


    }) 
// partial error
  

  
}
func handl_err_readcsv(){
  // partial error
   InstanceAPI.Server.HttpEngine.GET("/", func(c * gin.Context) {
        // if (CheckInitData(db_adapter)) {
            
        // }
        c.String(http.StatusOK, "No Routes Created \nCan't Recover From Reading CSV. ")

 
    }) 
}