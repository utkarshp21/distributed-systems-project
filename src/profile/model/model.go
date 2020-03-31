package model

import "sync"

var TweetsMux = &sync.Mutex{}