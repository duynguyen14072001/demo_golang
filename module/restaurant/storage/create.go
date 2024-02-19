package restaurantstorage

import (
	"context"
	"learn_golang/common"
	restaurantmodel "learn_golang/module/restaurant/model"
)

func (s *sqlStore) Create(context context.Context, data *restaurantmodel.RestaurantCreate) error {
	if err := s.db.Create(&data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
