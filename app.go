package main

import "net/http"

func fn(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello User!"))
}

func main() {
	http.HandleFunc("/", fn)
	http.ListenAndServe(":8080", nil)
}
