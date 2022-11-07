package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const APIURL string = "https://randomuser.me/api"

var client *http.Client

type RandomUsers struct {
	Results []struct {
		Gender string `json:"gender"`
		Name   struct {
			Title string `json:"title"`
			First string `json:"first"`
			Last  string `json:"last"`
		} `json:"name"`
		Location struct {
			Street struct {
				Number int    `json:"number"`
				Name   string `json:"name"`
			} `json:"street"`
			City        string `json:"city"`
			State       string `json:"state"`
			Country     string `json:"country"`
			Postcode    int    `json:"postcode"`
			Coordinates struct {
				Latitude  string `json:"latitude"`
				Longitude string `json:"longitude"`
			} `json:"coordinates"`
			Timezone struct {
				Offset      string `json:"offset"`
				Description string `json:"description"`
			} `json:"timezone"`
		} `json:"location"`
		Email string `json:"email"`
		Login struct {
			UUID     string `json:"uuid"`
			Username string `json:"username"`
			Password string `json:"password"`
			Salt     string `json:"salt"`
			Md5      string `json:"md5"`
			Sha1     string `json:"sha1"`
			Sha256   string `json:"sha256"`
		} `json:"login"`
		Dob struct {
			Date time.Time `json:"date"`
			Age  int       `json:"age"`
		} `json:"dob"`
		Registered struct {
			Date time.Time `json:"date"`
			Age  int       `json:"age"`
		} `json:"registered"`
		Phone string `json:"phone"`
		Cell  string `json:"cell"`
		ID    struct {
			Name  string `json:"name"`
			Value string `json:"value"`
		} `json:"id"`
		Picture struct {
			Large     string `json:"large"`
			Medium    string `json:"medium"`
			Thumbnail string `json:"thumbnail"`
		} `json:"picture"`
		Nat string `json:"nat"`
	} `json:"results"`
	Info struct {
		Seed    string `json:"seed"`
		Results int    `json:"results"`
		Page    int    `json:"page"`
		Version string `json:"version"`
	} `json:"info"`
}

func GetJson(url string, target interface{}) error {
	resp, err := client.Get(url)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(target)
}

func GetRandomUser() {
	var randomUsers RandomUsers

	err := GetJson(APIURL, &randomUsers)

	if err != nil {
		fmt.Printf("error getting BitCoint rate: %s\n", err.Error())
		return
	}

	b, err := json.MarshalIndent(randomUsers.Results[0], "", "  ")
	fmt.Println(string(b))
}

func main() {
	client = &http.Client{Timeout: 10 * time.Second}
	GetRandomUser()
}
