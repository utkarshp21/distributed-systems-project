package controller

import (
	"container/list"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"sync"
	"testing"
	authStorage "auth/storage"
	"golang.org/x/crypto/bcrypt"
	authmodel "auth/model"
)

func MockupUserData() error{

	for i:= 0 ; i < 1000 ; i++{

		user := authmodel.User{
			Username:"user"+strconv.Itoa(i),
			Password: "1234",
			FirstName: "us"+strconv.Itoa(i),
			LastName: "er"+strconv.Itoa(i),
			Followers: list.New(),
		}
		hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 5)

		if err != nil {
			return err
		}

		user.Password = string(hash)
		authStorage.Users[user.Username] = user
	}
	return nil
}

func TestSaveUser(t *testing.T) {

	wg := sync.WaitGroup{}
	for i:=0 ; i < 1000 ; i++ {
		wg.Add(1)
		go func(v int) {
			defer wg.Done()
			data := url.Values{}
			data.Set("username", "user"+strconv.Itoa(v)+"@gmail.com")
			data.Set("firstname", "us"+strconv.Itoa(v))
			data.Set("lastname", "er"+strconv.Itoa(v))
			data.Set("password", "1234")
			r, _ := http.NewRequest("POST", "", strings.NewReader(data.Encode()))
			r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
			_ , err := SaveUser(r)
			if err != nil {
				t.Error("Problem in adding user")
			}
		}(i)
	}
	wg.Wait()
	if len(authStorage.Users) == 1000 {
		t.Log("Test SaveUser succesful")
	}else{
		t.Errorf("Number of users missing %d",100-len(authStorage.Users))
	}
}

func TestLoginUser(t *testing.T) {

	err := MockupUserData()

	if err != nil {
		t.Error("Error in mocking up data")
	}
	wg := sync.WaitGroup{}
	for i := 0 ; i < 1000 ; i++{
		wg.Add(1)
		go func(v int) {
			defer wg.Done()
			authmodel.UsersMux.Lock()
			user := authStorage.Users["user"+strconv.Itoa(v)]
			authmodel.UsersMux.Unlock()
			_, err := LoginUser(user)
			if err != nil {
				t.Error("Error in login")
			}
		}(i)
	}
	wg.Wait()
	userCount := 0
	for i := 0 ; i < 1000 ; i++{
		user := authStorage.Users["user"+strconv.Itoa(i)]
		userToken := user.Token
		if userToken != ""{
			userCount = userCount + 1
		}else {
			t.Errorf("User%d is not logged in",i)
		}
	}
	if userCount == 1000{
		t.Log("Test case successful")
	}else {
		t.Errorf("Number of users not logged in: %d",1000-userCount)
	}
}
