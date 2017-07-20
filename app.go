package main

import (
	"encoding/json"
	"fmt"
)

type FOO struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	foo := FOO{"Tim", 100}

	b, err := json.Marshal(foo)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s", b)
}
