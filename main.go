package main

import (
	"html/template"
	"log"
	"net/http"
	"router/controller"

	"github.com/julienschmidt/httprouter"
)

var indexTmpl = template.Must(template.ParseFiles("./static/index.html"))
var loginTmpl = template.Must(template.ParseFiles("./static/login.html"))

//var mgoSession *mgo.Session

func main() {
	defer controller.MgoSession.Close()
	router := httprouter.New()
	router.HandleMethodNotAllowed = false //prevent router from sending 405 to request to same rout
	router.GET("/", indexHandler)
	router.GET("/login", loginGetHandler)
	router.POST("/login", loginPostHandler)
	router.POST("/newsoup", newSoupHandler)
	router.POST("/delete", deleteSoupHandler)
	router.GET("/static/:fileName", staticHandler)
	log.Fatal(http.ListenAndServe(":8082", router))
}

func indexHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	results, err := controller.GetAllSoups()
	if err != nil {
		panic(err)
	}
	indexTmpl.Execute(w, results)
}

func loginGetHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	loginTmpl.Execute(w, "") //take away the empty string
}

func loginPostHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	r.ParseForm()
	controller.UserLogin(r.PostForm["login"], r.PostForm["password"])
	loginTmpl.Execute(w, "Logged In")
}

func newSoupHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	//Parse body
	r.ParseForm()

	//create success channel
	successChan := make(chan bool)

	//call function that adds soup to db
	go controller.AddNewSoup(r.PostForm["name"], r.PostForm["origin"], r.PostForm["ingredients"], r.PostForm["spicy"], successChan)

	//receive success bool message
	success := <-successChan

	//write to template
	if success == true {
		http.Redirect(w, r, "/", 302) //fix this redirect
	} else {
		indexTmpl.Execute(w, "There was a problem in inserting your soup") //change this
	}

}

func deleteSoupHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	r.ParseForm()
	successChan := make(chan bool)
	go controller.DeleteSoup(r.PostForm["button"], successChan)

	success := <-successChan

	if success == true {
		http.Redirect(w, r, "/", 302)
	} else {
		http.Redirect(w, r, "/", 304)
	}

}

func staticHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	staticFilePath := "./static/" + ps.ByName("fileName")
	http.ServeFile(w, r, staticFilePath)
}
