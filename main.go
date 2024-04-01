package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Pair[F, S any] struct {
	First  F
	Second S
}
type IndexContent struct {
	MenuItems []Pair[string, string]
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./template/index.html")

	if err != nil {
		fmt.Println(err.Error())
	}
	menuList := queryDB("select name, link from navmenu")

	err1 := tmpl.Execute(w, IndexContent{
		MenuItems: menuList,
	})
	if err1 != nil {
		fmt.Println(err1.Error())
	}
}

func main() {
	createDB()
	print("hello")
	http.HandleFunc("/", indexHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err.Error())
	}
}
