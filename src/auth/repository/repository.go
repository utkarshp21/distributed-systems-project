package repository

import (
	authStorage "auth/storage"
	authmodel "auth/model"
	"golang.org/x/crypto/bcrypt"
	"container/list"
	profilemodel "profile/model"
	profileStorage "profile/storage"
	jwt "github.com/dgrijalva/jwt-go"
)

func ReturnUser(username string)(authmodel.User, bool){
	authmodel.UsersMux.Lock()
	user, exists := authStorage.Users[username]
	authmodel.UsersMux.Unlock()
	return user, exists
}

func SaveUser(user authmodel.User)(error){

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

func CheckLoginPassword(hashedPassword string, inputPassword string)(error){
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(inputPassword))
	return err
}

func SetCurrentUser(username string, user authmodel.User){
	authmodel.UsersMux.Lock()
	authStorage.Users[username] = user
	authmodel.UsersMux.Unlock()
}

func GenerateToken(user authmodel.User) (string,error){

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username":  user.Username,
		"firstname": user.FirstName,
		"lastname":  user.LastName,
	})

	tokenString, err := token.SignedString([]byte("secret"))

	if err != nil {
		return "",err
	}

	user.Token = tokenString

	SetCurrentUser(user.Username, user)

	return tokenString,nil
}

func SignoutUser(signoutUser authmodel.User)  {
	signoutUser.Token = ""
	SetCurrentUser(signoutUser.Username, signoutUser)
}