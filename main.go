package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
)

type tuple struct {
	lines, words, bytes uint
	name                string
}

func main() {
	params := os.Args[1:]
	wg := sync.WaitGroup{}
	fmt.Println(params)
	for _, path := range params {
		wg.Add(1)
		path := path
		go func() {
			tuple := wc(path)
			fmt.Println(tuple)
			wg.Done()
		}()
	}
	wg.Wait()

}

/**
 * -c only Bytes
 * -l only Lines
 * -w only Words
 */
func wc(path string) tuple {
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
	return result
}
