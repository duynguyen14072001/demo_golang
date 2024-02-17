package restaurantmodel

import (
	"errors"
	"learn_golang/common"
	"strings"
)

type RestaurantType string

const TypeNormal RestaurantType = "normal"
const TypePremium RestaurantType = "permium"

type Restaurant struct {
	common.SQLModel `json:",inline"`
	Id              int            `json:"id" gorm:"column:id;"`
	Name            string         `json:"name" gorm:"column:name;"`
	Addr            string         `json:"addr" gorm:"column:addr;"`
	Type            RestaurantType `json:"type" gorm:"column:type;"`
}

func (Restaurant) TableName() string { return "restaurants" }

type RestaurantCreate struct {
	common.SQLModel `json:",inline"`
	Id              int    `json:"id" gorm:"column:id;"`
	Name            string `json:"name" gorm:"column:name"`
	Addr            string `json:"addr" gorm:"column:addr"`
}

func (RestaurantCreate) TableName() string { return Restaurant{}.TableName() }

func (data *RestaurantCreate) Validate() error {
	data.Name = strings.TrimSpace(data.Name)

	if data.Name == "" {
		return ErrNameIsEmpty
	}

	return nil
}

type RestaurantUpdate struct {
	Name *string `json:"name" gorm:"column:name"`
	Addr *string `json:"addr" gorm:"column:addr"`
}

func (RestaurantUpdate) TableName() string { return Restaurant{}.TableName() }

var (
	ErrNameIsEmpty = errors.New("Name can not be empty")
)
