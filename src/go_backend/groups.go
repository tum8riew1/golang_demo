package main

import (
	"html/template"
	"log"
	"net/http"
	"reflect"

	"github.com/gorilla/mux"
)

func groupsHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "golang-session")
	//session.Values["token"]
	token := &Token{
		Token:   session.Values["token"].(string),
		App_url: app_url,
		Api_url: api_url,
	}
	t := template.Must(template.ParseFiles("views/layouts/layout.html", "views/partials/sidebar.html", "views/layouts/groups.html"))
	t.Execute(w, token)
}

func groupsAddHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "golang-session")
	//session.Values["token"]
	token := &Token{
		Token:   session.Values["token"].(string),
		App_url: app_url,
		Api_url: api_url,
	}
	t := template.Must(template.ParseFiles("views/layouts/layout.html", "views/partials/sidebar.html", "views/layouts/groups_add.html"))
	t.Execute(w, token)
}

func groupsEditHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	session, _ := store.Get(r, "golang-session")
	//session.Values["token"]
	token := &Token{
		Token:   session.Values["token"].(string),
		App_url: app_url,
		Api_url: api_url,
		ID:      params["id"],
	}
	log.Println(params["id"], reflect.TypeOf(params["id"]))
	t := template.Must(template.ParseFiles("views/layouts/layout.html", "views/partials/sidebar.html", "views/layouts/groups_edit.html"))
	t.Execute(w, token)
}
