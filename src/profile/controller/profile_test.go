package controller

import (
	"container/list"
	"strconv"
	authmodel "auth/model"
	authStorage "auth/storage"
	"sync"
	"testing"
	"golang.org/x/crypto/bcrypt"
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

func TestFollowUser(t *testing.T) {

	err := MockupUserData()

	if err != nil {
		t.Error("Error in mocking up data")
	}

	wg := sync.WaitGroup{}
	for i := 0 ; i < 10 ; i+=3 {
		wg.Add(1)
		go func(v int) {
			defer wg.Done()
			authmodel.UsersMux.Lock()
			followUser := authStorage.Users["user"+strconv.Itoa(v)]
			authmodel.UsersMux.Unlock()
			for j := v+1 ; j <= v+2 ; j++ {
				authmodel.UsersMux.Lock()
				user := authStorage.Users["user"+strconv.Itoa(j)]
				authmodel.UsersMux.Unlock()
				FollowUser(user,followUser)
			}
		}(i)
	}
	wg.Wait()
	for i := 0 ; i < 10 ; i+=3{
		authmodel.UsersMux.Lock()
		user := authStorage.Users["user"+strconv.Itoa(i)]
		authmodel.UsersMux.Unlock()
		if user.Followers.Len() != 2{
			t.Errorf("Followers for user %d unsuccesful",i)
		}
	}
	t.Log("Test FolloweUser successful")

}