package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"os"
)

func createDB() {
	_, checkFolder := os.Stat("./db")
	if checkFolder != nil {
		err3 := os.Mkdir("db", os.ModePerm)
		if err3 != nil {
			print(err3.Error())
		}

	}
	_, checkFile := os.Stat("./db/sql.db")
	if checkFile != nil {
		file, err := os.Create("./db/sql.db") // Create SQLite file
		if err != nil {
			print(err.Error())
		}
		err1 := file.Close()
		if err1 != nil {
			print(err1.Error())
		}
	}
	db, err5 := sql.Open("sqlite3", "./db/sql.db")
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

	defer db.Close()

}
func queryDB(q string, args ...any) []Pair[string, string] {
	db, err := sql.Open("sqlite3", "./db/sql.db")
	if err != nil {
		print(err.Error())
	}
	defer db.Close()
	rows, err := db.Query(q)
	var res []Pair[string, string]
	var name string
	var link string
	for rows.Next() {

		err := rows.Scan(&name, &link)
		if err != nil {
			fmt.Println(err)

		}
		res = append(res, Pair[string, string]{name, link})
	}
	return res
}
func execDB(q string, args ...any) sql.Result {
	db, err := sql.Open("sqlite3", "./db/sql.db")
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
