package controller

import (
	"container/list"
	//"fmt"
	"strconv"
	authmodel "auth/model"
	authStorage "auth/storage"
	"sync"
	"testing"
	"golang.org/x/crypto/bcrypt"
	profileStorage "profile/storage"
	//profilemodel "profile/model"
)

func MockupUserData() error{

	for i:= 0 ; i < 1000 ; i++{

		user := authmodel.User{
			Username:"user"+strconv.Itoa(i),
			Password: "1234",
			FirstName: "us"+strconv.Itoa(i),
			LastName: "er"+strconv.Itoa(i),
			Followers: list.New(),
			Token: "abcd"+strconv.Itoa(i),
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

func MockUpTweet() {

	for i:= 0 ; i < 1000 ; i++{
		tweetUser := "user"+strconv.Itoa(i)
		profileStorage.Tweets[tweetUser] = list.New()
	}
}

func MockUpTweetData()  {
	for i:= 0 ; i < 1000 ; i++{
		tweetUser := "user"+strconv.Itoa(i)
		tweetList := list.New()
		tweetList.PushBack("user"+strconv.Itoa(i)+" tweet 1")
		tweetList.PushBack("user"+strconv.Itoa(i)+" tweet 2")
		tweetList.PushBack("user"+strconv.Itoa(i)+" tweet 3")
		profileStorage.Tweets[tweetUser] = tweetList
	}

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

func TestSaveTweet(t *testing.T) {
	MockUpTweet()
	wg := sync.WaitGroup{}
	for i := 0 ; i < 10 ; i++ {
		wg.Add(1)
		go func(v int) {
			defer wg.Done()
			tweetUser := "user" + strconv.Itoa(v)
			for j := 0 ; j < 3 ; j++{
				tweetContent := tweetUser + "tweet" + strconv.Itoa(j)
				SaveTweet(tweetUser,tweetContent)
			}
		}(i)
	}
	wg.Wait()

	for i := 0 ; i < 10 ; i++ {
		tweetUser := "user" + strconv.Itoa(i)
		if profileStorage.Tweets[tweetUser].Len() != 3{
			t.Error("Error while saving tweet of"+tweetUser)
		}
	}
	t.Log("Test SaveTweet successful")
}

func TestFeedGenerate(t *testing.T) {

	MockUpTweetData()
	wg := sync.WaitGroup{}
	for i := 0 ; i < 10 ; i++ {
		wg.Add(1)
		go func(v int) {
			defer wg.Done()
			tweetUser := "user" + strconv.Itoa(v)
			feed := FeedGenerate(tweetUser)
			if feed == ""{
				t.Errorf("Feed generation unsuccessful for user%d",i)
			}
		}(i)
	}
	wg.Wait()
	t.Log("Test FeedGenerate successful")
}

func TestSignoutUser(t *testing.T) {

	MockupUserData()
	wg := sync.WaitGroup{}
	for i := 0 ; i < 10 ; i++ {
		wg.Add(1)
		go func(v int) {
			defer wg.Done()
			signoutUser := authStorage.Users["user" + strconv.Itoa(v)]
			SignoutUser(signoutUser)
			if authStorage.Users["user" + strconv.Itoa(v)].Token != ""{
				t.Errorf("Signout unsuccessful for %s",signoutUser.Username)
			}
		}(i)
	}
	wg.Wait()
	t.Log("Test SignoutUser successful")
}