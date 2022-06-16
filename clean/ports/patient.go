package ports

import "gomorganexam/domain" 

type ApplicationPort interface {
  Type string
  Port interface{}
}
  

type ServicePatientsAPI interface { 
	SvcTopPatientsAPI() domain.Patient 
}



type Patient interface {
	GetTopConfirmed(observation_date string, max_result int32) domain.Observation
	GetAllConfirmed(pagelimit int32) domain.Observation
}

type Service struct {
	Database   ApplicationAdapter
  App        ApplicationPort
}


func (svc *Service) GetTopConfirmed(author *Author)   {
 log.Println("No Implementation")
}


func (svc *Service) GetAllConfirmed(pagelimit int32) {
	 log.Println("No Implementation")
} 