package bot

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type chuckJoke struct {
	ID   int    `json:"id"`
	Fact string `json:"fact"`
}

// ChuckFact : Fetch Chuck Norris Joke
func ChuckFact() (string, error) {
	resp, err := http.Get("https://www.chucknorrisfacts.fr/api/get?data=tri:alea;type:txt;nb:1")
	if err != nil {
		fmt.Println("Could not fetch joke")
		return "nil", err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Unknown response body")
		return "nil", err
	}

	var joke []chuckJoke
	json.Unmarshal(body, &joke)
	return joke[0].Fact, nil
}
