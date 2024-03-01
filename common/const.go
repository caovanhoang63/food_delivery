package common

import "log"

// DbType is a type to represent the type of database
const (
	DbTypeRestaurant = 1
	DbTypeUser       = 2
	DbTypeFood       = 3
	DbTypeCategory   = 4
)

const CurrentUser = "user"

type Requester interface {
	GetUserId() int
	GetEmail() string
	GetRole() string
}

// AppRecover is an intelligent function to recover from panic
func AppRecover() {
	if err := recover(); err != nil {
		log.Println("Recovery error:", err)
	}
}
