package handler

import "net/http"

func Foo(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello from foo!"))
}
