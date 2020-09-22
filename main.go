package main

import "time"

func main() {
}


type TweetCounts struct {
 tweetName string
 date time.Time
}


func Constructor() TweetCounts {

}


func (this *TweetCounts) RecordTweet(tweetName string, time int)  {

}


func (this *TweetCounts) GetTweetCountsPerFrequency(freq string, tweetName string, startTime int, endTime int) []int {

}
