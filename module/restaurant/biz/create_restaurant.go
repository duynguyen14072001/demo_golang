package restaurantbiz

import (
	"context"
	"errors"
	restaurantmodel "learn_golang/module/restaurant/model"
)

type CreateRestaurantStore interface {
	CreateRestaurant(context context.Context, data *restaurantmodel.RestaurantCreate) error
}
type createRestaurantBiz struct {
	store CreateRestaurantStore
}

func NewCreateRestaurantBiz(store CreateRestaurantStore) *createRestaurantBiz {
	return &createRestaurantBiz{store: store}
}

func (biz *createRestaurantBiz) CreateRestaurant(context context.Context, data *restaurantmodel.RestaurantCreate) error {
	if data.Name == "" {
		return errors.New("Name khong duoc de trong")
	}
	if err := biz.store.CreateRestaurant(context, data); err != nil {
		return err
	}

	return nil
}
