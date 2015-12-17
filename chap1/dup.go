package chap1

import (
	"bufio"
	"fmt"
	"os"
)

//dup1 prints the text of each line that appears more than
//in the standard input, preceded by its count.
func dup1() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		counts[input.Text()]++
	}
	//NOTE: ignoring potential errors from input.Err()
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func dup2() {
	counts := make(map[string]int)
	files := os.Args[3:]
	d, _ := os.Getwd()
	println(d)

	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

type tots struct {
	cnt   int
	files map[string]int
}

func dup2_1() {

	counts := make(map[string]tots)
	files := os.Args[3:]
	d, _ := os.Getwd()
	println(d)

	if len(files) == 0 {
		countLines_1(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines_1(f, counts)
			f.Close()
		}
	}

	for line, n := range counts {
		//	    println(line)
		if n.cnt > 1 {
			fmt.Printf("%d\t%s\n", n.cnt, line)
		}
	}
}

func countLnsCtrl(f *os.File, counts interface{}) {
	switch counts := counts.(type) {
	case map[string]int:
		fmt.Println("can call func", counts)
	case map[string]tots:
		fmt.Println("can call another func, based on param type")
	}
}

//io.Reader will work too
func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
	//NOTE: ignoring potential errors from input.Err()
}

//io.Reader will work too
func countLines_1(f *os.File, counts map[string]tots) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		t, ok := counts[input.Text()]
		if !ok {
			t.files = make(map[string]int)
		}
		t.cnt++
		t.files[f.Name()]++
		counts[input.Text()] = t
	}
	fmt.Printf("%v", counts)
	//NOTE: ignoring potential errors from input.Err()
}
