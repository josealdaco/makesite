package main

import (
	"io/ioutil";
	"fmt"
)
func readFile() string {
	fileContents, err := ioutil.ReadFile("first-post.txt")
	if err != nil {
		panic(err)
	}

	return string(fileContents)
}
func main() {
	fmt.Println("Hello, world!")
	read:=readFile()
	fmt.Println(read)

}



func renderTemplate(){

}
