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
	return anaconda.NewTwitterApi(os.Getenv("ACCESS_TOKEN"), os.Getenv("ACCESS_TOKEN_SECRET"))
}

func main() {

	loadEnv()

	api := getTwitterApi()

	result_posts := request.GetTrendPosts()
	fmt.Printf("%v", result_posts[0].Title)

	log.Printf("Hello")
}
