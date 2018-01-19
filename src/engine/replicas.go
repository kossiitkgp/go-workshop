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

func First(query string, replicas ...Search) Result {
	c := make(chan Result)
	searchReplica := func(i int) { c <- replicas[i](query) }
	for i := range replicas {
		go searchReplica(i)
	}
	return <-c
}

func SearchEngine(query string) (results []Result) {
	c := make(chan Result)
	Web1, Web2 := fakeSearch("web1"), fakeSearch("web2")
	Image1, Image2 := fakeSearch("image1"), fakeSearch("image2")
	Video1, Video2 := fakeSearch("video1"), fakeSearch("video2")

	go func() { c <- First(query, Web1, Web2) }()
	go func() { c <- First(query, Image1, Image2) }()
	go func() { c <- First(query, Video1, Video2) }()
	timeout := time.After(80 * time.Millisecond)
	for i := 0; i < 3; i++ {
		select {
		case result := <-c:
			results = append(results, result)
		case <-timeout:
			fmt.Println("timed out")
			return
		}
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
