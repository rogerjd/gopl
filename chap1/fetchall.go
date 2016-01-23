package chap1

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

func fetchall(urls []string) {
	start := time.Now()
	ch := make(chan string)

	for _, url := range urls {
	    println(url)
		go fetchOne(url, ch) // start a goroutine
	}

	for range urls {
		fmt.Println(<-ch)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetchOne(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body) // dev/null
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, url)
}
