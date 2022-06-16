package ports

import (
  "gomorganexam/clean/core/domain" 
  "database/sql"
  "log"
  )
 
  

type ServicePatientsAPI interface { 
	GetTopConfirmed(observation_date string, max_result int32) domain.PatientObservation 
  GetAllConfirmed(pagelimit int32) domain.PatientObservation
}

 type Service struct {
	Database   *sql.DB
  
}

 func NewAPIService(db *sql.DB ) *Service {
   
    return &Service{
    		Database: db, 
    	}
   
 }

func (svc *Service) GetTopConfirmed(observation_date string, max_result int32)  {
   // svc.Database 
   
  
}


func (svc *Service) GetAllConfirmed(pagelimit int32) {
	 log.Println("No Implementation")
} 