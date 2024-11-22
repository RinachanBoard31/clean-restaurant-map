package db

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

type DbStoreDriver struct{}

func NewStoreDriver() *DbStoreDriver {
	return &DbStoreDriver{}
}

/* interfaceと型は同義。仮にgatewayがDBの型を知ったとしても、どんなDBから来たかわかるわけではないのでおk */
type FavoriteStore struct {
	Id                  string `gorm:"primaryKey"`
	UserId              int    `gorm:"not null"`
	User                User   `gorm:"foreignKey:UserId;references:Id"`
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

func (dbs *DbStoreDriver) FindFavorite(storeId string, userId int) (*FavoriteStore, error) {
	var stores []FavoriteStore
	err := DB.Where("store_id = ? AND user_id = ?", storeId, userId).Find(&stores).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	if len(stores) == 0 {
		return nil, nil
	}

	return &stores[0], nil
}

func (dbs *DbStoreDriver) SaveStore(dbStore *FavoriteStore) error {
	err := DB.Create(&dbStore).Error
	if err != nil {
		return err
	}
	return nil
}

func (dbs *DbStoreDriver) GetTopStores() ([]*FavoriteStore, error) {
	var stores []*FavoriteStore
	oneWeekAgo := time.Now().AddDate(0, 0, -7)

	// 1週間以内のデータを集計し、store_idごとにカウント、上位10件を取得
	subQuery := DB.Model(&FavoriteStore{}).
		Select("store_id, COUNT(*) as count").
		Where("created_at >= ?", oneWeekAgo).
		Group("store_id").
		Order("count(*) desc").
		Limit(10)

	// subQueryをJOINして取得
	err := DB.Model(&FavoriteStore{}).
		Joins("JOIN (?) as top_stores ON favorite_stores.store_id = top_stores.store_id", subQuery).
		Find(&stores).Error
	if err != nil {
		return nil, err
	}

	return stores, nil
}
