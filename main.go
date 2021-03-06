package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"router/controller"
	"router/db"
	"router/usersession"

	"github.com/julienschmidt/httprouter"
)

var indexTmpl = template.Must(template.ParseFiles("./static/index.html"))
var loginTmpl = template.Must(template.ParseFiles("./static/login.html"))

//var mgoSession *mgo.Session

func main() {
	defer db.MgoSession.Close()
	router := httprouter.New()
	router.HandleMethodNotAllowed = false //prevent router from sending 405 to request to same rout
	router.GET("/", indexHandler)
	router.GET("/login", loginGetHandler)
	router.POST("/login", loginPostHandler)
	router.GET("/logout", logOutHandler)
	router.POST("/newsoup", newSoupHandler)
	router.POST("/delete", deleteSoupHandler)
	router.GET("/assets/*filePath", staticHandler)
	//	log.Fatal(http.ListenAndServe(":8082", usersession.LoginWall(router)))
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
	message := make(chan string)
	cookiePointer := make(chan *http.Cookie)
	go controller.UserLogin(r.PostForm["login"], r.PostForm["password"], message, cookiePointer)

	//receive cookie from channel and put in variable
	cookieToSet := <-cookiePointer
	//set the cookie
	if cookieToSet != nil {
		http.SetCookie(w, cookieToSet)
	}

	//receive messsage from channel
	//This message will be used for toast message/notifications
	//for now just print
	messageToPrint := <-message

	fmt.Println(messageToPrint)

	http.Redirect(w, r, "/", 302) //fix this redirect
}

func logOutHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	cookieToBeDeleted, err := r.Cookie("session")
	if err != nil {
		fmt.Println(err)
	}

	//delete session from SessionMAP and database
	success := make(chan bool)

	go usersession.DeleteSession(cookieToBeDeleted.Value, success)

	//wait for OK

	ok := <-success

	var message string

	if ok == true {
		//send new empty cookie
		newCookie := http.Cookie{
			Name:   "session",
			Value:  cookieToBeDeleted.Value,
			MaxAge: -1,
		}
		http.SetCookie(w, &newCookie)
		message = "User Logged Out"
	} else if ok == false {
		message = "Problem Logging Out. Try Again."
	}

	loginTmpl.Execute(w, message)
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
	staticFilePath := "./assets/" + ps.ByName("filePath")
	http.ServeFile(w, r, staticFilePath)
}
