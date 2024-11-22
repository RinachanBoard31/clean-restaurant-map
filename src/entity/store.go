package model

import (
	"errors"
	"strconv"
)

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

func (l Location) Validate() error {
	// 緯度が-90から90の間にあるかチェック
	lat, err := strconv.ParseFloat(l.Lat, 64)
	if err != nil {
		return errors.New("latitude is invalid")
	}
	if lat < -90 || lat > 90 {
		return errors.New("latitude must be between -90 and 90, got " + l.Lat)
	}

	// 経度が-180から180の間にあるかチェック
	lng, err := strconv.ParseFloat(l.Lng, 64)
	if err != nil {
		return errors.New("longitude is invalid")
	}
	if lng < -180 || lng > 180 {
		return errors.New("longitude must be between -180 and 180, got " + l.Lng)
	}

	return nil
}

func NewStore(id string, name string, regularOpeningHours string, priceLevel string, lat string, lng string) (*Store, error) {
	location := Location{Lat: lat, Lng: lng}

	if err := location.Validate(); err != nil {
		return nil, err
	}

	store := &Store{
		Id:                  id,
		Name:                name,
		RegularOpeningHours: regularOpeningHours,
		PriceLevel:          priceLevel,
		Location:            location,
	}

	return store, nil
}
