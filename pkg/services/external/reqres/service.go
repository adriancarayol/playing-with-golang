package reqres

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

const apiURL = "https://reqres.in/api/users/"

type Result struct {
	Data User `json:"data"`
}

type User struct {
	Id          int `json:"id"`
	FirstName	string `json:"first_name"`
	LastName	string `json:"last_name"`
	Avatar		string `json:"avatar"`
}

func (user *User) RetrieveUserGivenId(id string) User {
    var resultUser User

	url := fmt.Sprintf("%s%s", apiURL, id)

	client := http.Client{
		Timeout: time.Second * 15,
	}

	request, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		log.Println(err)
		return resultUser
	}

	response, getErr := client.Do(request)

	if getErr != nil {
		log.Println(getErr)
		return resultUser
	}

	body, readErr := ioutil.ReadAll(response.Body)

	if readErr != nil {
		log.Println(readErr)
		return resultUser
	}


	var result Result
	jsonErr := json.Unmarshal(body, &result)

	if jsonErr != nil {
		log.Println(jsonErr)
		return resultUser
	}

	return result.Data
}