package main

import (
	"fmt"

	mydict "github.com/jyp90/practice/dict"
)

func main() {
	dictionary := mydict.Dictionary{"first": "FirstWord"}
	definition, err := dictionary.Search("second")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(definition)
	}

}
