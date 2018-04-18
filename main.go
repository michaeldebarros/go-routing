package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

var tmpl = template.Must(template.ParseFiles("./static/index.html"))

func main() {
	router := httprouter.New()
	router.GET("/", IndexHandler)
	router.POST("/newsoup", NewSoup)
	router.GET("/static/:fileName", StaticHandler)
	log.Fatal(http.ListenAndServe(":8082", router))
}

//IndexHandler exported function
func IndexHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	tmpl.Execute(w, "Welcom to my soup")
}

//NewSoup Handler exported function
func NewSoup(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	//Parse body
	r.ParseForm()
	fmt.Println(r.Form)
	addNewSoup(r.Form["name"], r.Form["origin"], r.Form["ingredients"], r.Form["spicy"])

	tmpl.Execute(w, "Your new soup was inserted")

}

//StaticHandler exported function
func StaticHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	staticFilePath := "./static/" + ps.ByName("fileName")
	http.ServeFile(w, r, staticFilePath)
}
