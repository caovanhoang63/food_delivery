package categorymodel

import "food-delivery/common"

const EntityName = "Category"

type Category struct {
	common.SqlModel `json:",inline"`
	Name            string        `json:"name" gorm:"column:name"`
	Description     string        `json:"description" gorm:"column:description"`
	Icon            *common.Image `json:"icon" gorm:"column:icon"`
}

func (c Category) TableName() string {
	return "categories"
}

func (c *Category) Mask(isAdmin bool) {
	c.GenUID(common.DbTypeCategory)
}

type CategoryCreate struct {
	common.SqlModel `json:",inline"`
	Name            string        `json:"name" gorm:"column:name"`
	Description     string        `json:"description" gorm:"column:description"`
	Icon            *common.Image `json:"icon" gorm:"column:icon"`
}

func (c CategoryCreate) TableName() string {
	return Category{}.TableName()
}

func (c *CategoryCreate) Mask() {
	c.GenUID(common.DbTypeCategory)
}

type CategoryUpdate struct {
	Name        string        `json:"name" gorm:"column:name"`
	Description string        `json:"description" gorm:"column:description"`
	Icon        *common.Image `json:"icon" gorm:"column:icon"`
}

func (c CategoryUpdate) TableName() string {
	return Category{}.TableName()
}
