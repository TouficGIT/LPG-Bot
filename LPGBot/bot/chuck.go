package bot

import (
	"encoding/json"
	"fmt"
	"html"
	"io/ioutil"
	"net/http"
)

type chuckJoke struct {
	ID   int    `json:"id"`
	Fact string `json:"fact"`
}

// ChuckFact : Fetch Chuck Norris Joke
func ChuckFact() (string, error) {
	fmt.Println("START : ChuckFact function - from chuck command")
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
	fmt.Println("Unmarshal the chuck norris quote")
	json.Unmarshal(body, &joke)
	fmt.Println(html.UnescapeString(joke[0].Fact))
	//UnescapeString used for accented characters (like "é")
	return html.UnescapeString(joke[0].Fact), nil
}
