package model

import (
	"container/list"
	"sync"
)

type User struct {
	Username  string `json:"username"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Password  string `json:"password"`
	Token     string `json:"token"`
	Followers *list.List `json:"followers"`
}

var UsersMux = &sync.Mutex{}
