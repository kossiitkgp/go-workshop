package main

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	Web   = fakeSearch("web")
	Image = fakeSearch("image")
	Video = fakeSearch("video")
)

type Result string

type Search func(query string) Result

func fakeSearch(kind string) Search {
	return func(query string) Result {
		t := time.Duration(rand.Intn(100)) * time.Millisecond
		time.Sleep(t)
		return Result(fmt.Sprintf("%5s result for %q in %v\n", kind, query, t))
	}
}

func SearchEngine(query string) (results []Result) {
	results = append(results, Web(query))
	results = append(results, Image(query))
	results = append(results, Video(query))
	return
}

func main() {
	rand.Seed(time.Now().UnixNano())
	start := time.Now()
	results := SearchEngine("golang")
	elapsed := time.Since(start)
	fmt.Println(results)
	fmt.Println(elapsed)
}
