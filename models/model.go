package models


type Service struct {
	Id string
	Category  string
	types []ServiceType
}

type ServiceType struct {
	Id string
	Name  string
	Time string
	Gender string
	Price string
}