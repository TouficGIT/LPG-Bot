package bot

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"Work.go/LPG-Bot/LPGBot/config"
)

type JokeApiResponse struct {
	Value JokeBody `json:"value"`
	Type  string   `json:"type"`
}

type JokeBody struct {
	Id         int      `json:"id"`
	Joke       string   `json:"joke"`
	Categories []string `json:"categories"`
}

//Send joke
func sendJoke() (err error) {
	// Get a Chuck Norris joke
	joke, err := getJoke()
	if err != nil {
		return err
	}
	// Send
	resp, err := http.PostForm(config.Webhook, url.Values{"content": {joke}, "tts": {"false"}})
	fmt.Println(resp)
	if err != nil {
		fmt.Println("Couldn't send message")
		fmt.Println(err)
		return err
	} else {
		fmt.Println(resp)
		return err
	}

	return nil
}

//Fetch Chuck Norris Joke
func getJoke() (string, error) {
	resp, err := http.Get("http://api.icndb.com/jokes/random")
	if err != nil {
		fmt.Println("Could not fetch joke")
		return "nil", err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Unknown response body")
		return "nil", err
	}

	var jokeResp JokeApiResponse
	json.Unmarshal(body, &jokeResp)
	fmt.Println(jokeResp)
	return jokeResp.Value.Joke, nil
}
