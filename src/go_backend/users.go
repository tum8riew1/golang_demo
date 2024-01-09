package main

import (
	"html/template"
	"log"
	"net/http"
	"reflect"

	"github.com/gorilla/mux"
)

func usersHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "golang-session")
	//session.Values["token"]
	token := &Token{
		Token:   session.Values["token"].(string),
		App_url: app_url,
		Api_url: api_url,
	}
	t := template.Must(template.ParseFiles("views/layouts/layout.html", "views/partials/sidebar.html", "views/layouts/users.html"))
	t.Execute(w, token)
}

func usersAddHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "golang-session")
	//session.Values["token"]
	token := &Token{
		Token:   session.Values["token"].(string),
		App_url: app_url,
		Api_url: api_url,
		User_id: session.Values["id"].(int),
	}
	t := template.Must(template.ParseFiles("views/layouts/layout.html", "views/partials/sidebar.html", "views/layouts/users_add.html"))
	t.Execute(w, token)
}

func usersEditHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	session, _ := store.Get(r, "golang-session")
	//session.Values["token"]
	token := &Token{
		Token:   session.Values["token"].(string),
		App_url: app_url,
		Api_url: api_url,
		User_id: session.Values["id"].(int),
		ID:      params["id"],
	}
	log.Println(params["id"], reflect.TypeOf(params["id"]))
	t := template.Must(template.ParseFiles("views/layouts/layout.html", "views/partials/sidebar.html", "views/layouts/users_edit.html"))
	t.Execute(w, token)
}
