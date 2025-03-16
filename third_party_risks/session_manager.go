package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/sessions"
)

var (
	store = sessions.NewCookieStore([]byte("super-secret-key"))
)

func main() {
	http.HandleFunc("/set", func(w http.ResponseWriter, r *http.Request) {
		session, _ := store.Get(r, "session-name")

		session.Options = &sessions.Options{
			Path:     "/",
			MaxAge:   86400 * 7,
		}

		session.Values["authenticated"] = true
		session.Save(r, w)
		fmt.Fprintf(w, "Session set")
	})

	http.HandleFunc("/get", func(w http.ResponseWriter, r *http.Request) {
		session, _ := store.Get(r, "session-name")

		auth, ok := session.Values["authenticated"].(bool)
		if !ok || !auth {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}

		fmt.Fprintf(w, "Welcome!")
	})

	fmt.Println("Server is running on :8081")
	http.ListenAndServe(":8081", nil)
}
