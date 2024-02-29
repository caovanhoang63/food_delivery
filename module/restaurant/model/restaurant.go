package restaurantmodel

import (
	"errors"
	"food-delivery/common"
	"strings"
)

const (
	EntityName = "restaurant"
)

// Restaurant is a model that represents a restaurant
type Restaurant struct {
	common.SqlModel `json:",inline"`
	Name            string         `json:"name" gorm:"column:name"`
	Addr            string         `json:"addr" gorm:"column:addr"`
	Logo            *common.Image  `json:"logo" gorm:"column:logo"`
	Cover           *common.Images `json:"cover" gorm:"column:cover"`
}

// TableName is a function to change the table name
func (Restaurant) TableName() string { return "restaurants" }

func (data *Restaurant) Mask(isAdminOrOwner bool) {
	data.GenUID(common.DbTypeRestaurant)

}

// RestaurantCreate is a model that client use to create a new restaurant
type RestaurantCreate struct {
	common.SqlModel `json:",inline"`
	Name            string         `json:"name" gorm:"column:name"`
	Addr            string         `json:"addr" gorm:"column:addr"`
	Logo            *common.Image  `json:"logo" gorm:"column:logo"`
	Cover           *common.Images `json:"cover" gorm:"column:cover"`
	OwnerId         int            `json:"-" gorm:"column:owner_id"`
}

func (data *RestaurantCreate) Mask(isAdminOrOwner bool) {
	data.GenUID(common.DbTypeRestaurant)

}

func (data *RestaurantCreate) Validate() error {
	data.Name = strings.TrimSpace(data.Name)

	if data.Name == "" {
		return ErrNameIsEmpty
	}
	return nil
}

// TableName is a function to change the table name
func (RestaurantCreate) TableName() string { return Restaurant{}.TableName() }

// RestaurantUpdate is a model that client use to update a restaurant
type RestaurantUpdate struct {
	Name  *string        `json:"name" gorm:"column:name"`
	Addr  *string        `json:"addr" gorm:"column:addr"`
	Logo  *common.Image  `json:"logo" gorm:"column:logo"`
	Cover *common.Images `json:"cover" gorm:"column:cover"`
}

// TableName is a function to change the table name
func (RestaurantUpdate) TableName() string { return Restaurant{}.TableName() }

var (
	ErrNameIsEmpty = errors.New("name can not be empty")
)
