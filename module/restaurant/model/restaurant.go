package restaurantmodel

import "food-delivery/common"

// Restaurant is a model that represents a restaurant
type Restaurant struct {
	common.SqlModel `json:",inline"`
	Name            string `json:"name" gorm:"column:name"`
	Addr            string `json:"addr" gorm:"column:addr"`
}

// TableName is a function to change the table name
func (Restaurant) TableName() string { return "restaurants" }

// RestaurantCreate is a model that client use to create a new restaurant
type RestaurantCreate struct {
	common.SqlModel `json:",inline"`
	Name            string `json:"name" gorm:"column:name"`
	Addr            string `json:"addr" gorm:"column:addr"`
}

// TableName is a function to change the table name
func (RestaurantCreate) TableName() string { return Restaurant{}.TableName() }

// RestaurantUpdate is a model that client use to update a restaurant
type RestaurantUpdate struct {
	Name *string `json:"name" gorm:"column:name"`
	Addr *string `json:"addr" gorm:"column:addr"`
}

// TableName is a function to change the table name
func (RestaurantUpdate) TableName() string { return Restaurant{}.TableName() }
