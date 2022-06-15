package domain

type Country struct {
    confirmed int32
    deaths  int32
  recovered int32
}
  

type Observation struct {
    date string
    countries  Country
}


type Patient interface{
     GetTopConfirmed(observation_date string , max_result int32 ) ( Observation )

  
}