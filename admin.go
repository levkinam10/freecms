package main

import (
	"html/template"
	"net/http"
)

func adminPanelHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("./template/admin.html")
	tmpl.Execute(w, nil)
}
