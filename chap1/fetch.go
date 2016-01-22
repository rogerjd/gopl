package chap1

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func fetch(args []string) {
	for _, url := range args {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		b, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v", url, err)
		}
		fmt.Printf("%s", b)
	}
}
