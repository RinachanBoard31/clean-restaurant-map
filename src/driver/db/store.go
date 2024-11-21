package db

import (
	"time"
)

type DbStoreDriver struct{}

func NewStoreDriver() *DbStoreDriver {
	return &DbStoreDriver{}
}

/* interfaceと型は同義。仮にgatewayがDBの型を知ったとしても、どんなDBから来たかわかるわけではないのでおk */
type FavoriteStore struct {
	Id                  string `gorm:"primaryKey"`
	Name                string `gorm:"not null"`
	RegularOpeningHours string
	PriceLevel          string
	Latitude            string `gorm:"not null"`
	Longitude           string `gorm:"not null"`
	CreatedAt           time.Time
	UpdatedAt           time.Time
}

func (dbs *DbStoreDriver) GetStores() ([]*FavoriteStore, error) {
	var stores []*FavoriteStore
	err := DB.Find(&stores).Error
	if err != nil {
		return nil, err
	}
	return stores, nil
}

func (dbs *DbStoreDriver) SaveStore(store *FavoriteStore) error {
	err := DB.Create(&store).Error
	if err != nil {
		return err
	}
	return nil
}
