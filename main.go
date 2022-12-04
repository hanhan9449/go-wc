package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type tuple struct {
	lines, words, bytes uint
	name                string
}

func main() {
	ch := make(chan tuple)
	done := 0
	params := os.Args[1:]
	fmt.Println(params)
	for _, path := range params {
		go wc(ch, path)
	}
	for done != len(params) {
		value := <-ch
		done++
		fmt.Println(value)
	}
	close(ch)

}

/**
 * -c only Bytes
 * -l only Lines
 * -w only Words
 */
func wc(ch chan tuple, path string) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal("参数错误")
	}
	var lines, words, bytes uint
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines++
		next := scanner.Bytes()
		bytes += uint(len(next))
		words += uint(len(strings.Fields(fmt.Sprintf("%s", next))))
	}
	result := tuple{lines: lines, words: words, bytes: bytes, name: path}
	ch <- result
}
