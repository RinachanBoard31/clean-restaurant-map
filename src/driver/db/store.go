package db

import (
	"time"
)

type DbStoreDriver struct{}

func NewStoreDriver() *DbStoreDriver {
	return &DbStoreDriver{}
}

/* interfaceと型は同義。仮にgatewayがDBの型を知ったとしても、どんなDBから来たかわかるわけではないのでおk */
type Store struct {
	Id                  string `gorm:"primaryKey"`
	Name                string
	RegularOpeningHours string
	PriceLevel          string
	Latitude            string
	Longitude           string
	CreatedAt           time.Time
	UpdatedAt           time.Time
}

func (dbs *DbStoreDriver) GetStores() ([]*Store, error) {
	var stores []*Store
	err := DB.Find(&stores).Error
	if err != nil {
		return nil, err
	}
	return stores, nil
}
