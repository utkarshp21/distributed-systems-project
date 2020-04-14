package repository

import (
	"container/list"
	"context"
	"time"
	"strconv"
	"sync"
	"testing"
	profileStorage "profile/storage"
)


func MockUpTweet() {
	for i:= 0 ; i < 10 ; i++{
		tweetUser := "user"+strconv.Itoa(i)
		profileStorage.Tweets[tweetUser] = list.New()
	}
}

func MockUpTweetData() {
	for i:= 0 ; i < 10 ; i++{
		tweetUser := "user"+strconv.Itoa(i)
		tweetList := list.New()
		tweetList.PushBack("user"+strconv.Itoa(i)+" tweet 1")
		tweetList.PushBack("user"+strconv.Itoa(i)+" tweet 2")
		tweetList.PushBack("user"+strconv.Itoa(i)+" tweet 3")
		profileStorage.Tweets[tweetUser] = tweetList
	}

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
				err := SaveTweet(tweetUser,tweetContent,context.Background())
				if err != nil {
					t.Log("Problem in saving tweet of user", v)
				}
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

func TestSaveTweetContext(t *testing.T) {
	MockUpTweet()
	wg := sync.WaitGroup{}
	for i := 0 ; i < 10 ; i++ {
		wg.Add(1)
		go func(v int) {
			defer wg.Done()
			tweetUser := "user" + strconv.Itoa(v)
			if v % 2 == 0{
				ctx , cancel := context.WithTimeout(context.Background(),time.Duration(100)*time.Second)
				defer cancel()
				for j := 0 ; j < 3 ; j++ {
					tweetContent := tweetUser + "tweet" + strconv.Itoa(j)
					err := SaveTweet(tweetUser, tweetContent, ctx)
					if err != nil {
						t.Log("Problem in saving tweet of user", v)
					}
				}

			}else{
				ctx , cancel := context.WithTimeout(context.Background(),time.Duration(1)*time.Millisecond)
				defer cancel()
				for j := 0 ; j < 3 ; j++ {
					tweetContent := tweetUser + "tweet" + strconv.Itoa(j)
					err := SaveTweet(tweetUser, tweetContent, ctx)
					if err != nil {
						t.Log("Problem in saving tweet of user", v)
					}
				}
			}
		}(i)
	}
	wg.Wait()

	t.Log("Tweets of user0")
	for e:=profileStorage.Tweets["user0"].Front(); e !=nil; e = e.Next(){
		t.Log(e.Value)
	}

	t.Log("Tweets of user1")
	for e:=profileStorage.Tweets["user1"].Front(); e !=nil; e = e.Next(){
		t.Log(e.Value)
	}

	for i := 0 ; i < 10 ; i=i+2 {
		tweetUser := "user" + strconv.Itoa(i)
		if profileStorage.Tweets[tweetUser].Len() != 3{
			t.Error("Error while saving tweet of"+tweetUser)
		}
		tweetUser2 := "user" + strconv.Itoa(i+1)
		if profileStorage.Tweets[tweetUser2].Len() != 0{
			t.Error("Error while saving tweet of"+tweetUser)
		}
	}
	t.Log("Test SaveTweetContext successful")
}

//Test case for GetTweetList with context cancel is similar to without cancellation(error is returned)
func TestGetTweetList(t *testing.T) {

	MockUpTweetData()
	wg := sync.WaitGroup{}
	for i := 0 ; i < 10 ; i++ {
		wg.Add(1)
		go func(v int) {
			defer wg.Done()
			tweetUser := "user" + strconv.Itoa(v)
			tweetList, err := GetTweetList(tweetUser,context.Background())
			if err != nil || tweetList.Len() != 3{
				t.Errorf("Getting list of tweets unsuccessful for user%d",i)
			}
		}(i)
	}
	wg.Wait()
	t.Log("Test GetTweetList successful")
}

//Test case for InitialiseTweetsDB does nothing but initialises Tweets repo
//Test case  for InitialiseTweetsDB with context cancel is similar to SaveUserRegister test case with context cancel