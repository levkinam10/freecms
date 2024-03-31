package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./template/index.html")
	if err != nil {
		fmt.Print("Err:")
		fmt.Println(err)
	}
	err1 := tmpl.Execute(w, nil)
	if err1 != nil {
		fmt.Print("Err:")
		fmt.Println(err1)
	}
}

func main() {
	print("hello")
	http.HandleFunc("/", indexHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Print("Err:")
		fmt.Println(err)
	}
}
