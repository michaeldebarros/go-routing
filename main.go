package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	router.GET("/", IndexHandler)
	router.GET("/static/:fileName", StaticHandler)
	log.Fatal(http.ListenAndServe(":8082", router))
}

//IndexHandler exported function
func IndexHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	http.ServeFile(w, r, "./static/index.html")
}

//StaticHandler exported function
func StaticHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	staticFilePath := "./static/" + ps.ByName("fileName")
	http.ServeFile(w, r, staticFilePath)
}
