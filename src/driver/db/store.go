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
	UserId              string `gorm:"not null"`
	StoreId             string `gorm:"not null"`
	StoreName           string `gorm:"not null"`
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

func (dbs *DbStoreDriver) SaveStore(dbStore *FavoriteStore) error {
	err := DB.Create(&dbStore).Error
	if err != nil {
		return err
	}
	return nil
}
