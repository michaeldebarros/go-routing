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
	router.GET("/", indexHandler)
	router.POST("/newsoup", newSoupHandler)
	router.POST("/delete", deleteSoupHandler)
	router.GET("/static/:fileName", staticHandler)
	log.Fatal(http.ListenAndServe(":8082", router))
}

func indexHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	results, err := getAllSoups()
	if err != nil {
		panic(err)
	}
	tmpl.Execute(w, results)
}

func newSoupHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
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
		http.Redirect(w, r, "/", 200)
		//tmpl.Execute(w, "Your new soup was inserted")
	} else {
		tmpl.Execute(w, "There was a problem in inserting your soup")
	}

}

func deleteSoupHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	r.ParseForm()
	successChan := make(chan bool)
	go deleteSoup(r.PostForm["button"], successChan)

	success := <-successChan

	if success == true {
		http.Redirect(w, r, "/", 200)
	} else {
		w.Write([]byte("There was a problem deleting your soup"))
	}

}

func staticHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	staticFilePath := "./static/" + ps.ByName("fileName")
	http.ServeFile(w, r, staticFilePath)
}
