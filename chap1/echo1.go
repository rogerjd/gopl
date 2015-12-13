package chap1

import (
	"fmt"
	"strings"
	"time"
)

func echo1(args []string) {
	fmt.Println("echo1")

	s, sep := "", " "
	for i := 0; i < len(args); i++ {
		s += args[i] + sep
	}
	s = s[0 : len(s)-1]
	fmt.Println(s)
}

func echo2(args []string) {
	fmt.Println("echo1")

	s, sep := "", " "
	for _, arg := range args {
		s += arg + sep
	}
	s = s[0 : len(s)-1]
	fmt.Println(s)
}

func ex1_1(args []string) {
	fmt.Println("ex1_1")

	for i, arg := range args {
		fmt.Println(i, arg)
	}
}

func ex1_2(args []string) {
	fmt.Println("ex1_2")

	fmt.Println(strings.Join(args, " "))
}

func ex1_3(args []string) {
	fmt.Println("ex1_3")
	
	args = append(args, "abcdef", "ghijkl", "mmopqr", "stuvwx", "yz")
	
	t1 := time.Now()
	echo1(args)
	n1 := time.Since(t1).Seconds()
	
	t2 := time.Now()
	ex1_2(args)
	n2 := time.Since(t2).Seconds()
	
	fmt.Println(n1, n2)
}
