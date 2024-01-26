package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

type Restaurant struct {
	Id   int    `json:"id" gorm:"column:id"`
	Name string `json:"name" gorm:"column:name"`
	Addr string `json:"addr" gorm:"column:addr"`
}

func (Restaurant) TableName() string { return "restaurants" }

type RestaurantUpdate struct {
	Name *string `json:"name" gorm:"column:name"`
	Addr *string `json:"addr" gorm:"column:addr"`
}

func (RestaurantUpdate) TableName() string { return Restaurant{}.TableName() }

func main() {
	//dsn aka connection string
	dsn := os.Getenv("MYSQL_CONN_STRING")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(db, err)
	}

	//newRestaurant := Restaurant{Name: "aaa", Addr: "aa1"}
	//
	//db.Create(&newRestaurant)

	//var myRestaurant Restaurant

	//myRestaurant.Name = ""
	//if err = db.Where("id = ?", 2).First(&myRestaurant).Error; err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(myRestaurant)

	//if err = db.Where("id = ?", 2).Updates(&myRestaurant).Error; err != nil {
	//	log.Fatal(err)
	//}
	//
	//fmt.Println(myRestaurant)

	//if err = db.Table("restaurants").Where("id = ?", 1).Delete(nil).Error; err != nil {
	//	log.Fatal(err)
	//}

	//newName := "aaa"
	//dataUpdate := RestaurantUpdate{Name: &newName}
	//if err = db.Where("id = ?", 2).Updates(&dataUpdate).Error; err != nil {
	//	log.Fatal(err)
	//}

}
