package ports

import "gomorganexam/domain"

type ServicePatientsAPI interface {
	ServeReadAPI() domain.Patient
	 
}
