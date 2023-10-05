package handler

import "net/http"

func Bar(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello from Bar!"))
}
