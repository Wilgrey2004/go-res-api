package routes

import "net/http"

func HomeHandle(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hellow world!"))
}
