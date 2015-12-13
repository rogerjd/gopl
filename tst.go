package main

import (
	"fmt"
	"github.com/rogerjd/gopl/chap1"
	"os"
)

func main() {
	fmt.Println("hi gopl")
	fmt.Println(os.Args)
	if len(os.Args) < 3 {
		fmt.Println("please enter: chap excer and params if nec")
		return
	}
	chap := os.Args[1]
	switch chap {
	case "1":
		chap1.Main(os.Args)
	}

	/*
	   //	chap9.LinksFromURL("gobook.html") //panic unsupported protocol scheme ""
	   //	chap9.LinksFromURL("http://gobook.html") //w3 searchassist
	   	urls, err := chap9.LinksFromURL("http://www.qtrac.eu/") //ok; goes to index.html by default (home page)
	   	fmt.Println(urls, err)

	   	//todo: for each returned url/link, call http.Get on it, chk error/resp code
	*/
}
