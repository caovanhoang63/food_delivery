package foodmodel

import (
	"food-delivery/common"
	restaurantmodel "food-delivery/module/restaurant/model"
)

const EntityName = "Food"

type Food struct {
	common.SqlModel `json:",inline"`
	Name            string         `json:"name" gorm:"column:name"`
	Description     string         `json:"description" gorm:"column:description"`
	Price           float32        `json:"price" gorm:"column:price"`
	Images          *common.Images `json:"Images" gorm:"column:images"`
	RestaurantId    int
	Restaurant      *restaurantmodel.Restaurant `json:"restaurant" gorm:"Preload:false;"`
}

func (f Food) TableName() string {
	return "foods"
}

func (f *Food) Mask(isAdmin bool) {
	f.GenUID(common.DbTypeFood)
}

type FoodCreate struct {
	common.SqlModel `json:",inline"`
	Name            string         `json:"name" gorm:"column:name"`
	Description     string         `json:"description" gorm:"column:description"`
	Price           float32        `json:"price" gorm:"column:price"`
	Images          *common.Images `json:"Images" gorm:"column:images"`
	RestaurantId    int
}
