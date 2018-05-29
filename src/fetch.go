package main

import (
	"fmt"
	"fun/net"
	"time"
)

type result struct {
	url,
	title string
	timeit time.Duration
}

func fetch(urls ...string) {
	for _, u := range urls {
		start := time.Now()
		data, err := net.GetRandBytes(u)
		timeit := time.Since(start)
		if err != nil {
			fmt.Println("[Error] Fetching", u, err)
			break
		}

		doc := string(data)
		title := getTitle(doc)
		r := result{
			u,
			title,
			timeit,
		}

		fmt.Printf("[Fetched] %s:\nTitle: %s\nTimeit: %s\n\n",
			r.url, r.title, r.timeit,
		)
	}

}
