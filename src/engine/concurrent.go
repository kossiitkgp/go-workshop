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
	c := make(chan Result)
	go func() { c <- Web(query) }()
	go func() { c <- Image(query) }()
	go func() { c <- Video(query) }()

	// Can we use for result := range c here instead?
	for i := 0; i < 3; i++ {
		result := <-c
		results = append(results, result)
	}
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
