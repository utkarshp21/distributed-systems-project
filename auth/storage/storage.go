package storage

import (
	"auth/model"
	"container/list"
)

var Users = make(map[string]model.User)