package chap1

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func fetchall(urls []string) {
	for _, url := range urls {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			return
		}
		b, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			return
		}
	    fmt.Printf("%s", b)
	}
}
