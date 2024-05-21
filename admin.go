package main

import (
	"crypto/sha256"
	"crypto/subtle"
	"html/template"
	"net/http"
)

func adminPanelHandler(w http.ResponseWriter, r *http.Request) {
	username, password, ok := r.BasicAuth()
	if ok {
		usernameHash := sha256.Sum256([]byte(username))
		passwordHash := sha256.Sum256([]byte(password))
		expectedUsernameHash := sha256.Sum256([]byte("your expected username"))
		expectedPasswordHash := sha256.Sum256([]byte("your expected password"))

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
			tmpl, _ := template.ParseFiles("./template/admin.html")
			err := tmpl.Execute(w, nil)
			print(err)
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
