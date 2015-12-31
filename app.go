package main

import (
	"net/http"
	"os"
)

func fn(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello User!"))
}

func main() {
	http.HandleFunc("/", fn)
	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}
