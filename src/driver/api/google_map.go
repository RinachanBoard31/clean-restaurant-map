package api

type ApiGoogleMapDriver struct{}

func NewGoogleMapDriver() *ApiGoogleMapDriver {
	return &ApiGoogleMapDriver{}
}

type Store struct {
	Id                  string   `json:"places.id"`
	Name                string   `json:"places.displayName"`
	RegularOpeningHours []string `json:"places.regularOpeningHours.weekdayDescriptions"`
	PriceLevel          string   `json:"places.priceLevel"`
}

func (ap *ApiGoogleMapDriver) GetStores() ([]*Store, error) {
	dummyStores := make([]*Store, 0)
	dummyStores = append(dummyStores, &Store{
		Id:                  "Id001",
		Name:                "UEC cafe",
		RegularOpeningHours: []string{"Sat: 06:00 - 22:00", "Sun: 06:00 - 22:00"},
		PriceLevel:          "PRICE_LEVEL_MODERATE",
	})
	dummyStores = append(dummyStores, &Store{
		Id:                  "Id002",
		Name:                "UEC restaurant",
		RegularOpeningHours: []string{"Sat: 11:00 - 20:00", "Sun: 11:00 - 20:00"},
		PriceLevel:          "PRICE_LEVEL_INEXPENSIVE",
	})
	return dummyStores, nil
}
