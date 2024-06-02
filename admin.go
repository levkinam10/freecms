package main

import (
	"crypto/sha256"
	"crypto/subtle"
	"fmt"
	"html/template"
	"net/http"
)

var (
	username1 string = "1234"
	password1 string = "1234"
)

type editPost struct {
	Title    string
	Id       string
	Desc     string
	PostText string
	Img      string
}

func adminPanelHandler(w http.ResponseWriter, r *http.Request) {
	username, password, ok := r.BasicAuth()
	if ok {
		usernameHash := sha256.Sum256([]byte(username))
		passwordHash := sha256.Sum256([]byte(password))
		expectedUsernameHash := sha256.Sum256([]byte(username1))
		expectedPasswordHash := sha256.Sum256([]byte(password1))

		// Use the subtle.ConstantTimeCompare() function to check if
		// the provided username and password hashes equal the
		// expected username and password hashes. ConstantTimeCompare
		// will return 1 if the values are equal, or 0 otherwise.
		// Importantly, we should to do the work to evaluate both the
		// username and password before checking the return values to
		// avoid leaking information.
		usernameMatch := (subtle.ConstantTimeCompare(usernameHash[:], expectedUsernameHash[:]) == 1)
		passwordMatch := (subtle.ConstantTimeCompare(passwordHash[:], expectedPasswordHash[:]) == 1)

		// If the username and password are correct, then call
		// the next handler in the chain. Make sure to return
		// afterwards, so that none of the code below is run.
		if usernameMatch && passwordMatch {
			if r.Method == "POST" {
				if err := r.ParseForm(); err != nil {
					fmt.Fprintf(w, "ParseForm() err: %v", err)
				} else {
					id := r.FormValue("postId")
					title := r.FormValue("title")
					desc := r.FormValue("desc")
					postText := r.FormValue("post")
					img := r.FormValue("img")
					UpdatePost(id, title, desc, img, postText)
				}
			}
			posts := ListPosts()
			tmpl, _ := template.ParseFiles("./template/admin.html")
			tmpl.Execute(w, IndexContent{
				Title:    "Admin Panel",
				PostList: posts,
			})

			return
		}
	}

	// If the Authentication header is not present, is invalid, or the
	// username or password is wrong, then set a WWW-Authenticate
	// header to inform the client that we expect them to use basic
	// authentication and send a 401 Unauthorized response.
	w.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
	http.Error(w, "Unauthorized", http.StatusUnauthorized)
}
func editHandler(w http.ResponseWriter, r *http.Request) {
	username, password, ok := r.BasicAuth()
	if ok {
		usernameHash := sha256.Sum256([]byte(username))
		passwordHash := sha256.Sum256([]byte(password))
		expectedUsernameHash := sha256.Sum256([]byte(username1))
		expectedPasswordHash := sha256.Sum256([]byte(password1))
		usernameMatch := (subtle.ConstantTimeCompare(usernameHash[:], expectedUsernameHash[:]) == 1)
		passwordMatch := (subtle.ConstantTimeCompare(passwordHash[:], expectedPasswordHash[:]) == 1)

		if usernameMatch && passwordMatch {
			//posts := ListPosts()
			id := r.PathValue("id")
			post := GetPost1(id)
			tmpl, err := template.ParseFiles("./template/editPost.html")
			if err != nil {
				println(err)
			}
			tmpl.Execute(w, post)
			return
		}
	}

	w.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
	http.Error(w, "Unauthorized", http.StatusUnauthorized)
}
func CreateHandler(w http.ResponseWriter, r *http.Request) {
	username, password, ok := r.BasicAuth()
	if ok {
		usernameHash := sha256.Sum256([]byte(username))
		passwordHash := sha256.Sum256([]byte(password))
		expectedUsernameHash := sha256.Sum256([]byte(username1))
		expectedPasswordHash := sha256.Sum256([]byte(password1))
		usernameMatch := (subtle.ConstantTimeCompare(usernameHash[:], expectedUsernameHash[:]) == 1)
		passwordMatch := (subtle.ConstantTimeCompare(passwordHash[:], expectedPasswordHash[:]) == 1)

		if usernameMatch && passwordMatch {
			//posts := ListPosts()
			id := CreatePost()
			http.Redirect(w, r, "/admin/edit/"+id, http.StatusFound)
			return
		}
	}

	w.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
	http.Error(w, "Unauthorized", http.StatusUnauthorized)
}
func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	username, password, ok := r.BasicAuth()
	if ok {
		usernameHash := sha256.Sum256([]byte(username))
		passwordHash := sha256.Sum256([]byte(password))
		expectedUsernameHash := sha256.Sum256([]byte(username1))
		expectedPasswordHash := sha256.Sum256([]byte(password1))
		usernameMatch := (subtle.ConstantTimeCompare(usernameHash[:], expectedUsernameHash[:]) == 1)
		passwordMatch := (subtle.ConstantTimeCompare(passwordHash[:], expectedPasswordHash[:]) == 1)

		if usernameMatch && passwordMatch {
			id := r.PathValue("id")
			DeletePost(id)
			http.Redirect(w, r, "/admin/", http.StatusFound)
			return
		}
	}

	w.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
	http.Error(w, "Unauthorized", http.StatusUnauthorized)
}
