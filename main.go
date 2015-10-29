package main

import (
	"github.com/julienschmidt/httprouter"
	"io"
	"log"
	"net/http"
	"os"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var (
		f   *os.File
		err error
	)

	if f, err = os.Open("Inigo-Dragon-Slayer.jpg"); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "image/jpeg")
	io.Copy(w, f)
}

func main() {
	router := httprouter.New()
	router.GET("/", Index)

	log.Fatal(http.ListenAndServe(":1337", router))
}
