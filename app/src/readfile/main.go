package main

import (
  "fmt"
	"io/ioutil"
	"log"
  "reflect"
)

func main() {
  content, err := ioutil.ReadFile("pod.yml")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("File contents: %s", content)
  fmt.Println(reflect.TypeOf(content)) // []uint8
  string(content) // conversion to string
}
