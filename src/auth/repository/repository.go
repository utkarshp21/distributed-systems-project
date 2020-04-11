package repository

import (
	authStorage "auth/storage"
	authmodel "auth/model"
	"golang.org/x/crypto/bcrypt"
	"container/list"
	profilemodel "profile/model"
	profileStorage "profile/storage"
)

func CheckUserExists(username string)bool{
	authmodel.UsersMux.Lock()
	_, exists := authStorage.Users[username]
	authmodel.UsersMux.Unlock()
	return exists
}

func SaveUser(user authmodel.User)error{

	//create hash for password
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 5)
	if err != nil {
		return err
	}
	user.Password = string(hash)

	//save user to storage
	authmodel.UsersMux.Lock()
	authStorage.Users[user.Username] = user
	authmodel.UsersMux.Unlock()

	//create empty tweets list for newly registered user
	profilemodel.TweetsMux.Lock()
	profileStorage.Tweets[user.Username] = list.New()
	profilemodel.TweetsMux.Unlock()

	return nil
}