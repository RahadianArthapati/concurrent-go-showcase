package main

import (
	"fmt"
	"time"
	"sync"
	"net/http"
)

func returnType(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("error $s\n", err)
		return
	}
	defer resp.Body.Close()
	ctype := resp.Header.Get("content-type")
	fmt.Println("%s -> %s\n", url, ctype)
}
func siteSerial(urls []string) {
	for _, url := range urls {
		returnType(url)
	}
}
func siteConcurrent(urls []string) {
	var wg sync.WaitGroup
	for _,url := range urls {
		wg.Add(1)
		go func(url string) {
			returnType(url)
			wg.Done()
		}(url)
		wg.Wait()
	}
}
func main() {
	urls := []string {
		"https://golang.org",
		"https://api.github.com",
	}
	start := time.Now()
	siteSerial(urls)
	fmt.Println(time.Since(start))

	startConcurrent := time.Now()
	siteConcurrent(urls)
	fmt.Println(time.Since(startConcurrent))

}