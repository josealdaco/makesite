package main

import (
	"flag"
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
)

type ContentFile struct {
	Title   string
	Content string
}

func save(filename string) {
	fileContents, err := ioutil.ReadFile(filename) // Get text contents

	fmt.Println(string(fileContents)) // Convert and print contents to terminal
	if err != nil {
		panic(err)
	}

	content := ContentFile{
		Title:   "first-post-w" ,  // Place all content in a struct
		Content: string(fileContents),
	}
	t := template.Must(template.ParseFiles("template.tmpl")) // Create template

	file, err := os.Create(filename[:len(filename)-4] + ".html") /// Create html file
	t.Execute(os.Stdout, content) // Print html to terminal
	t.Execute(file, content) /// Exctract data to new file that we created
	//file_flag := flag.String("latest-post.html", "defaultValue", ".txt") /// For the last step we updated the file name to extension html
	//flag.Parse()
	//fmt.Println(*file_flag) // We will have a pointerto the file flag value
	//fmt.Println(*file_flag) // We will have a pointerto the file flag value

}


func FlagEvent() {
	fileFlag := flag.String("file", ".txt", "a string")
	dirFlag := flag.String("dir", "search", "a string")
	flag.Parse()
	fmt.Println(*fileFlag)
	fmt.Println(*dirFlag)
}

func main() {
	save("first-post.txt")
	save("test1.txt")  // Save creates an html file for every filename we give it
	save("test2.txt")
	save("text3.txt")
	FlagEvent() // Prints all txt files in the  same directory.
}
