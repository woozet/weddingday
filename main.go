package weddingday

import (
	"net/http"
)

func init() {
	http.HandleFunc("/", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
	http.FileServer(http.Dir("static"))
}

func main() {

}
