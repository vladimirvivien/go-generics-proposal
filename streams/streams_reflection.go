package main

import (
	"fmt"
	"log"
	"reflect"
	"strings"
)

// Stream reads data from stream and applies fn on each
// item read from stream. This function uses good old empty
// interface{} and reflection to implement its functionality.
// The code below will handle the following types
// - [] T, [size] T
// - chan T
// The code can extend to support many more types.  Note that
// this will work fine. However, the compiler cannot provide
// any type safety and relies on runtime code to catch illegal
// values.
func Stream(stream interface{}, fn func(interface{}) error) error {
	if stream == nil || fn == nil {
		return fmt.Errorf("stream and func required")
	}

	sType := reflect.TypeOf(stream)
	sVal := reflect.ValueOf(stream)

	switch sType.Kind() {
	// [] T, [size] T
	case reflect.Slice, reflect.Array:
		for i := 0; i < sVal.Len(); i++ {
			elem := sVal.Index(i)
			if err := fn(elem.Interface()); err != nil {
				log.Println(err)
			}
		}
	case reflect.Chan:
		for{
			elem, ok := sVal.Recv()
			if !ok {
				break
			}
			if err := fn(elem.Interface()); err != nil{
				log.Println(err)
			}
		}
	default:
		return fmt.Errorf("unsupported type: %s", sType.Name())
	}

	return nil
}

func main() {
	if err := Stream(
		[]int{2, 4, 6, 8},
		func(val interface{}) error {
			i,ok := val.(int)
			if !ok {
				return fmt.Errorf("unsupported type %s", val)
			}
			fmt.Println(i * i)
			return nil
		},
	); err != nil {
		log.Fatal(err)
	}

	if err := Stream(
		[]string{"2", "4", "6", "8"},
		func(val interface{}) error {
			i,ok := val.(string)
			if !ok {
				return fmt.Errorf("unsupported type %s", val)
			}
			fmt.Printf("%s little piggies\n", i)
			return nil
		},
	); err != nil {
		log.Fatal(err)
	}

	stream := make(chan string)
	go func(){
		defer close(stream)
		stream <- "Hello"; stream <- "World!"
	}()

	if err := Stream(
		stream,
		func(val interface{}) error{
			i,ok := val.(string)
			if !ok {
				return fmt.Errorf("unsupported type %s", val)
			}
			fmt.Println(strings.ToUpper(i))
			return nil
		},
	); err != nil {
		log.Fatal(err)
	}
}
