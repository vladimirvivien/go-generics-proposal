package main

import (
	"fmt"
	"log"
	"strings"
)

func StreamIntSlice(stream []int, fn func(int)) error {
	if stream == nil || fn == nil {
		return fmt.Errorf("stream and func are required")
	}
	for _, i := range stream {
		fn(i)
	}
	return nil
}

func StreamStringSlice(stream []string, fn func(string)) error {
	if stream == nil || fn == nil {
		return fmt.Errorf("stream and func are required")
	}
	for _, i := range stream {
		fn(i)
	}
	return nil
}

func StreamIntChan(stream <-chan int, fn func(int)) error {
	if stream == nil || fn == nil {
		return fmt.Errorf("stream and func are required")
	}
	for i := range stream {
		fn(i)
	}
	return nil
}

func StreamStringChan(stream <-chan string, fn func(string)) error {
	if stream == nil || fn == nil {
		return fmt.Errorf("stream and func are required")
	}
	for i := range stream {
		fn(i)
	}
	return nil
}

func main() {

	if err := StreamIntSlice(
		[]int{2, 4, 6, 8},
		func(val int) {
			fmt.Println(val * val)
		},
	); err != nil {
		log.Fatal(err)
	}

	if err := StreamStringSlice(
		[]string{"2", "4", "6", "8"},
		func(val string) {
			fmt.Printf("%s little piggies\n", val)
		},
	); err != nil {
		log.Fatal(err)
	}

	stream := make(chan string)
	go func(){
		defer close(stream)
		stream <- "Hello"; stream <- "World!"
	}()

	if err := StreamStringChan(
		stream,
		func(val string) {
			fmt.Println(strings.ToUpper(val))
		},
	); err != nil {
		log.Fatal(err)
	}
}
