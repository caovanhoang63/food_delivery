package restaurantlikemodel

import (
	"food-delivery/common"
	"time"
)

const EntiryName = "UserLikeRestaurant"

type Like struct {
	RestaurantId int                `json:"restaurant_id" gorm:"column:restaurant_id;"`
	UserId       int                `json:"user_id" gorm:"column:user_id;"`
	CreatedAt    *time.Time         `json:"created_at" gorm:"column:created_at;"`
	User         *common.SimpleUser `json:"user" gorm:"preload:false"`
}

func (Like) TableName() string {
	return "restaurant_likes"
}

func (l *Like) GetRestaurantId() int {
	return l.RestaurantId
}

func ErrCannotLikeRestaurant(err error) *common.AppError {
	return common.NewCustomError(
		err,
		"cannot like restaurant",
		"Err_Cannot_Like_Restaurant",
	)
}

func ErrCannotDislikeRestaurant(err error) *common.AppError {
	return common.NewCustomError(
		err,
		"cannot dislike restaurant",
		"Err_Cannot_Dislike_Restaurant",
	)
}
