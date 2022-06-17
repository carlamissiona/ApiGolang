package ports

import (
    "gomorganexam/clean/core/domain"

    "gomorganexam/clean/adapters/in"
    "gomorganexam/clean/adapters/out/db"

    // _"database/sql"
    "log"
    "os"
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

func NewAPIService(data * db.Adapter, records[] byte, engine * in .WebAdapter)( * Service) {

    return &Service {
        Data: data,
        Records: nil,
        Server: engine,
    }

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
        log.Println("=>ERROR!! NO CSV FILE \n %v \n=>END OF ERROR!! NO CSV FILE " , err ); 
        handl_err_nocsv()
    }

    // remember to close the file at the end of the program
    defer csvfile.Close()
 
    // read csv values using csv.Reader
    csv_reader := csv.NewReader(csvfile)

    data, err := csv_reader.ReadAll()
    if err != nil {
        log.Println("=>ERROR!! CANT READ CSV FILE \n %v \n=>END OF ERROR!! NO CSV FILE " , err ); 
        
        handl_err_readcsv()
    }
    log.Println("API data")
    log.Printf("=>DEBUG!! API DATA")
    log.Println(data)
    log.Printf("=>END OF DEBUG!! API DATA")

}
 
func handl_err_nocsv() {


  
}
func handl_err_readcsv(){
  
}