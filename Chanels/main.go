package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"https://github.com/",
		"https://www.instagram.com/",
		"https://www2.kickassanime.ro/",
		"https://www.dizibox.vip/",
		"https://www.youtube.com/",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
		//fmt.Println(<-c)
	}
	for l := range c {
		go func(link string) {
			time.Sleep(time.Second * 5)
			checkLink(link, c)
		}(l)

	}

}

func checkLink(s string, c chan string) {
	_, err := http.Get(s)
	if err != nil {
		fmt.Println(s + " might be down")
		c <- s
		return
	}
	fmt.Println(s + " works")
	c <- s
}
