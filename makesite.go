package main

import (
	"io/ioutil";
	"fmt";
	"html/template";
        "os";
)
type Entry struct {
        Content string

}


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
	paths := []string{
		 "template.tmpl",
	   }
	 td := Entry{read}

	 t := template.Must(template.New("template.tmpl").ParseFiles(paths...))
         var err = t.Execute(os.Stdout, td)
         if err != nil {
           panic(err)
         }

}



func renderTemplate(){

}
