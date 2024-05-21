package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"time"
)

type Pair[F, S any] struct {
	First  F
	Second S
}
type IndexContent struct {
	MenuItems []Pair[any, any]
	Img       string
	PostList  []PostPreview
}
type Post struct {
	ID       string
	Title    string
	PostDate time.Time
	PostText any
	Nav      []Pair[any, any]
}
type PostPreview struct {
	ID          string
	Title       string
	PostDate    time.Time
	Description string
	ImageLink   string
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./template/index.html")

	if err != nil {
		fmt.Println(err.Error())
	}
	menuList := queryDB("select name, link from navmenu")
	posts := ListPosts()
	err1 := tmpl.Execute(w, IndexContent{
		MenuItems: menuList,
		Img:       "",
		PostList:  posts,
	})
	if err1 != nil {
		fmt.Println(err1.Error())
	}
}
func PostHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./template/post.html")
	id := r.PathValue("id")
	if err != nil {
		fmt.Println(err.Error())
	}
	err1 := tmpl.Execute(w, GetPost(id))
	if err1 != nil {
		fmt.Println(err1.Error())
	}

}

func main() {
	_, checkFolder := os.Stat("./data/img")
	if checkFolder != nil {
		err3 := os.Mkdir("data/img", os.ModePerm)
		if err3 != nil {
			print(err3.Error())
		}
	}

	http.Handle("/static/", http.FileServer(http.Dir("./")))
	http.Handle("/img/", http.FileServer(http.Dir("./data/")))

	createDB()
	print("hello")
	http.HandleFunc("/{route}/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/"+r.PathValue("route"), 301)
	})
	http.HandleFunc("/admin", adminPanelHandler)
	http.HandleFunc("/post/{id}", PostHandler)
	http.HandleFunc("/", indexHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err.Error())
	}
}
