package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
)

var tmpl = template.Must(template.ParseFiles("./static/index.html"))

var mgoSession *mgo.Session

func main() {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	mgoSession = session
	defer mgoSession.Close() // take a closer look at this

	router := httprouter.New()
	router.GET("/", IndexHandler)
	router.POST("/newsoup", NewSoup)
	router.GET("/static/:fileName", StaticHandler)
	log.Fatal(http.ListenAndServe(":8082", router))
}

//IndexHandler exported function
func IndexHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	results, err := getAllSoups()
	if err != nil {
		panic(err)
	}
	tmpl.Execute(w, results)
}

//NewSoup Handler exported function
func NewSoup(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	//Parse body
	r.ParseForm()

	//create success channel
	successChan := make(chan bool)

	//call function that adds soup to db
	go addNewSoup(r.PostForm["name"], r.PostForm["origin"], r.PostForm["ingredients"], r.PostForm["spicy"], successChan)

	//receive success bool message
	success := <-successChan

	//write to template
	if success == true {
		tmpl.Execute(w, "Your new soup was inserted")
	} else {
		tmpl.Execute(w, "There was a problem in inserting your soup")
	}

}

//StaticHandler exported function
func StaticHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	staticFilePath := "./static/" + ps.ByName("fileName")
	http.ServeFile(w, r, staticFilePath)
}
