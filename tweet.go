package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"unicode/utf8"

	"./request"
	"github.com/ChimeraCoder/anaconda"
	"github.com/joho/godotenv"
)

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func getTwitterApi() *anaconda.TwitterApi {
	anaconda.SetConsumerKey(os.Getenv("TWITTER_KEY"))
	anaconda.SetConsumerSecret(os.Getenv("TWITTER_SECRET"))
	return anaconda.NewTwitterApi(os.Getenv("ACCESS_TOKEN"), os.Getenv("ACCESS_SECRET"))
}

func tweetTrendPosts(api *anaconda.TwitterApi, posts []request.Post) {
	var post_titles []string
	var text_count int
	for _, v := range posts {
		text_count += utf8.RuneCountInString(v.Title)
		if text_count > 120 {
			break
		}
		post_titles = append(post_titles, v.Title)
	}
	posts_text := strings.Join(post_titles, "\n")
	fmt.Println(posts_text)
	fmt.Printf("%v \n", utf8.RuneCountInString(posts_text))
	tweet_text := posts_text + "\n#tagmaru\ntagmaru.me"
	_, err := api.PostTweet(tweet_text, nil)
	if err != nil {
		log.Fatal(err)
	}
}
func main() {

	loadEnv()

	api := getTwitterApi()

	result_posts := request.GetTrendPosts()
	fmt.Printf("%v", result_posts[0].Title)

	tweetTrendPosts(api, result_posts)

	log.Printf("Hello")
}
