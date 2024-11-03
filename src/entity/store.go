package model

type Location struct {
	Lat string
	Lng string
}

type Store struct {
	Id                  string
	Name                string
	RegularOpeningHours string
	PriceLevel          string
	Location            Location
}
