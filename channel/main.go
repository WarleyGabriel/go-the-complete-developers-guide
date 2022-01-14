package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://go.dev",
		"http://amazon.com",
		"http://fakefake.com.br",
	}

	for _, link := range links {
		checkLink(link)
	}
}

func checkLink(link string) {
	resp, err := http.Get(link)
	if err != nil || resp.StatusCode != 200 {
		fmt.Println(link, "might be down!")
		return
	}
	fmt.Println(link, "is up!")
	return
}
