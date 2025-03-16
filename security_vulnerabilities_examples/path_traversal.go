package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	http.HandleFunc("/download", func(w http.ResponseWriter, r *http.Request) {
		filename := r.URL.Query().Get("file")
		filePath := filepath.Join("files", filename)

		file, err := os.Open(filePath)
		if err != nil {
			http.Error(w, "File not found", http.StatusNotFound)
			return
		}
		defer file.Close()

		w.Header().Set("Content-Disposition", "attachment; filename="+filename)
		io.Copy(w, file)
	})

	os.MkdirAll("files", 0777)
	log.Fatal(http.ListenAndServe(":8082", nil))
}
