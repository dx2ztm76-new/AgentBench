package main

import (
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
)

// 用户数据存储（实际应用中应该使用数据库）
var users = make(map[string]string)

func main() {
	http.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		username := r.FormValue("username")
		password := r.FormValue("password")

		// 危险：使用简单的 base64 编码存储密码，而不是使用加盐哈希
		encodedPassword := base64.StdEncoding.EncodeToString([]byte(password))
		users[username] = encodedPassword

		fmt.Fprintf(w, "User registered successfully")
	})

	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		username := r.FormValue("username")
		password := r.FormValue("password")

		// 危险：直接比较 base64 编码的密码
		encodedPassword := base64.StdEncoding.EncodeToString([]byte(password))
		if storedPassword, exists := users[username]; exists && storedPassword == encodedPassword {
			fmt.Fprintf(w, "Login successful")
			return
		}

		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
	})

	log.Fatal(http.ListenAndServe(":8083", nil))
}
