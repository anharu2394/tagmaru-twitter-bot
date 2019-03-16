package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
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

func tweetEachOfAPost(api *anaconda.TwitterApi, posts []request.Post, first_i int) {
	for _, v := range posts[first_i:] {
		var tweet_text string
		switch v.Url {
		case "qiita":
			tweet_text = "今日のトレンド記事にゃん\n" + v.Title + "\n" + "qiita.com" + v.Url
		case "devto":
			tweet_text = "今日のトレンド記事にゃん\n" + v.Title + "\n" + "dev.to" + v.Url
		default:
			tweet_text = "今日のトレンド記事にゃん\n" + v.Title + "\n" + v.Url
		}
		_, err := api.PostTweet(tweet_text, nil)
		if err != nil {
			log.Fatal(err)
		}
		time.Sleep(30 * time.Minute)
	}
}

func main() {
	flag.Parse()
	loadEnv()

	api := getTwitterApi()

	result_posts := request.GetTrendPosts()

	arg_num, err := strconv.Atoi(flag.Arg(0))

	if flag.Arg(0) == "trend" {
		fmt.Println("I will tweet trend.")
		tweetTrendPosts(api, result_posts)
	} else if err == nil {
		fmt.Println("I will tweet one post.")
		fmt.Printf("Get number arg: %v \n", arg_num)
		tweetEachOfAPost(api, result_posts, arg_num)
	}

	log.Printf("have Finished!")
}
