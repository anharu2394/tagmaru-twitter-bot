package request

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type Post struct {
	Id       int    `json:"id"`
	Title    string `json:"title"`
	Url      string `json:"url"`
	FabCount int    `json:"fab_count"`
}

func GetTrendPosts() []Post {
	resp, err := http.Get("https://api.tagmaru.me/api/posts/trend")
	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		log.Fatal(err)
	}

	var posts []Post
	err = json.Unmarshal(body, &posts)
	if err != nil {
		log.Fatal(err)
	}
	return posts
}
