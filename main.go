package main

import (
	"fmt"
	"github.com/hootuu/gelato/io/serializer"
)

type User struct {
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Height int    `json:"height"`
}

func main() {
	dict := map[interface{}]interface{}{
		"D": "XX",
		"A": "B",
		"B": "C",
		"0": "D",
		"1": []string{"A", "B", "E", "0"},
		1:   "0000",
		2:   0,
		3:   true,
		4:   4.0,
		true: &User{
			Name:   "jack",
			Age:    20,
			Height: 180,
		},
	}
	fmt.Println(dict)
	bytes, err := serializer.Serialize(&User{
		Name:   "donnie",
		Age:    20,
		Height: 180,
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("the result::::")
	fmt.Println(bytes)
	fmt.Println(string(bytes))
}
