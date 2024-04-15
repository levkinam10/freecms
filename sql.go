package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"github.com/russross/blackfriday"
	"html/template"
	"os"
	"time"
)

func createDB() {
	_, checkFolder := os.Stat("./data")
	if checkFolder != nil {
		err3 := os.Mkdir("data", os.ModePerm)
		if err3 != nil {
			print(err3.Error())
		}

	}
	_, checkFile := os.Stat("./data/sql.db")
	if checkFile != nil {
		file, err := os.Create("./data/sql.db") // Create SQLite file
		if err != nil {
			print(err.Error())
		}
		err1 := file.Close()
		if err1 != nil {
			print(err1.Error())
		}
	}
	db, err5 := sql.Open("sqlite3", "./data/sql.db")
	if err5 != nil {
		print(err5.Error)
	}
	_, err10 := db.Exec("CREATE TABLE IF NOT EXISTS config ( key TEXT, value TEXT)")
	if err10 != nil {
		print(err10.Error())
	}
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS navmenu (name TEXT, link TEXT)")
	if err != nil {
		print(err.Error())
	}
	_, err11 := db.Exec("CREATE TABLE IF NOT EXISTS posts (id TEXT, title TEXT, postdate DATETIME, posttext TEXT, description TEXT, description_imagelink TEXT)")
	if err11 != nil {
		print(err11.Error())
	}

	defer db.Close()

}
func queryDB(q string, args ...any) []Pair[any, any] {
	db, err := sql.Open("sqlite3", "./data/sql.db")
	if err != nil {
		print(err.Error())
	}
	defer db.Close()
	rows, err := db.Query(q)
	var res []Pair[any, any]
	var first any
	var second any
	for rows.Next() {

		err := rows.Scan(&first, &second)
		if err != nil {
			fmt.Println(err)

		}
		res = append(res, Pair[any, any]{first, second})
	}
	return res
}
func execDB(q string, args ...any) sql.Result {
	db, err := sql.Open("sqlite3", "./data/sql.db")
	if err != nil {
		print(err.Error())
	}
	defer db.Close()
	res, err := db.Exec(q, args)
	if err != nil {
		print(err.Error())
	}
	return res
}

func ListPosts() []PostPreview {
	db, err := sql.Open("sqlite3", "./data/sql.db")
	if err != nil {
		print(err.Error())
	}
	defer db.Close()
	rows, err := db.Query("SELECT id, title, postdate, description, description_imagelink FROM posts ORDER BY postdate DESC ")
	var res []PostPreview
	var id string
	var title string
	var postdate time.Time
	var description string
	var description_imagelink string
	for rows.Next() {

		err := rows.Scan(&id, &title, &postdate, &description, &description_imagelink)
		if err != nil {
			fmt.Println(err)

		}
		res = append(res, PostPreview{id, title, postdate, description, description_imagelink})
	}
	return res
}
func GetPost(id string) Post {
	db, err := sql.Open("sqlite3", "./data/sql.db")
	if err != nil {
		print(err.Error())
	}
	defer db.Close()
	rows, err := db.Query("SELECT id, title, postdate, posttext FROM posts WHERE id =?", id)
	menuList := queryDB("select name, link from navmenu")
	var res Post
	var id1 string
	var title string
	var postdate time.Time
	var posttext string
	for rows.Next() {
		err := rows.Scan(&id, &title, &postdate, &posttext)
		if err != nil {
			fmt.Println(err)

		}
		res = Post{id1, title, postdate, template.HTML(blackfriday.MarkdownCommon([]byte(posttext))), menuList}
	}
	return res
}
