package basics

import "net/http"

type helloHandler struct{}

func (h *helloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world"))
}

func main4() {
	http.Handle("/", &helloHandler{})
	http.ListenAndServe(":9090", nil)
}
