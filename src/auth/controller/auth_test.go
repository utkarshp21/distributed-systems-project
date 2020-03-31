package controller

import (
	"net/http"
	"net/url"
	"strings"
	"sync"
	"testing"
	authStorage "auth/storage"
)

func TestSaveUser(t *testing.T) {

	wg := sync.WaitGroup{}
	for i:=0 ; i < 100 ; i++ {
		go func(v int) {
			wg.Add(1)
			defer wg.Done()
			data := url.Values{}
			data.Set("username", "user"+string(v)+"@gmail.com")
			data.Set("firstname", "us"+string(v))
			data.Set("lastname", "er"+string(v))
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
	if len(authStorage.Users) == 100 {
		t.Log("Test SaveUser succesful")
	}else{
		//t.Error("Number of users missing %d",100-len(authStorage.Users))
		t.Error("Test failed")
	}
}
