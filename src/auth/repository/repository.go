package repository

import (
	authStorage "auth/storage"
	authmodel "auth/model"
	"container/list"
	profilemodel "profile/model"
	profileStorage "profile/storage"
	//jwt "github.com/dgrijalva/jwt-go"
	//"golang.org/x/crypto/bcrypt"
)

func ReturnUser(username string)(authmodel.User, bool){
	resultChan := make(chan authmodel.User)
	errChan := make(chan bool)
	go func() {
		authmodel.UsersMux.Lock()
		user, exists := authStorage.Users[username]
		authmodel.UsersMux.Unlock()
		resultChan <- user
		errChan <- exists
	}()

	return <-resultChan, <-errChan
}

func SaveUser(user authmodel.User){
	resultChan := make(chan bool)
	go func() {
		//save user to storage
		authmodel.UsersMux.Lock()
		authStorage.Users[user.Username] = user
		authmodel.UsersMux.Unlock()

		//create empty tweets list for newly registered user
		profilemodel.TweetsMux.Lock()
		profileStorage.Tweets[user.Username] = list.New()
		profilemodel.TweetsMux.Unlock()

		resultChan <- true
	}()
	<-resultChan
}

func SetCurrentUser(username string, user authmodel.User) {
	resultChan := make(chan bool)
	go func() {

		authmodel.UsersMux.Lock()
		authStorage.Users[username] = user
		authmodel.UsersMux.Unlock()

		resultChan <- true
	}()
	<-resultChan
}

