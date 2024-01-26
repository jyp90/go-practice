package main

import (
	"fmt"

	mydict "github.com/jyp90/practice/dict"
)

func main() {
	dictionary := mydict.Dictionary{}
	word := "hello"
	baseWord := "First"
	dictionary.Add(word, baseWord)
	fmt.Println("Base Word: ", baseWord)
	err2 := dictionary.Update(word, "Second")
	if err2 != nil {
		fmt.Println(err2)
	}
	dictionary.Delete(word)
	word2, _ := dictionary.Search(word)
	fmt.Println("After Update: ", word2)
}
