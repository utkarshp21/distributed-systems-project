package storage

import (
	"container/list"
)

var Tweets = make(map[string]*list.List)