package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

type UserToken struct {
	Token string
}
type UserGetByUsername struct {
	Username string `json:"username"`
}
type JsonData struct {
	Message     string `json:"message"`
	Status_code int    `json:"status_code"`
}
type Data struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	RoleID string `json:"role_id"`
	Status string `json:"status"`
}
type UserDetail struct {
	Data       Data `json:"data"`
	StatusCode int  `json:"status_code"`
}
type Token struct {
	Token   string
	App_url string
	Api_url string
	ID      string
	User_id int
}

var (
	app_url                        = "http://localhost:3000"
	app_path_dashboard             = "/admin/dashboard"
	api_url                        = "http://localhost:3001"
	api_path_users_get_by_username = "/api/v1/backend/users/get-by-username"
	// key must be 16, 24 or 32 bytes long (AES-128, AES-192 or AES-256)
	key   = []byte("xxxxxxx")
	store = sessions.NewCookieStore(key)
)

const (
	YYYYMMDDhhmmss = "2006-01-02 15:04:05"
)

//const app_url = "http://localhost:3000"

func main() {
	log.SetFlags(log.Lmicroseconds)
	r := mux.NewRouter()
	r.HandleFunc("/", loginHandler).Methods("GET")
	admin := r.PathPrefix("/admin").Subrouter()
	admin.HandleFunc("/dashboard", authMiddleware(dashboardHandler)).Methods("GET")
	admin.HandleFunc("/groups", authMiddleware(groupsHandler)).Methods("GET")
	admin.HandleFunc("/groups/add", authMiddleware(groupsAddHandler)).Methods("GET")
	admin.HandleFunc("/groups/edit/{id}", authMiddleware(groupsEditHandler)).Methods("GET")
	admin.HandleFunc("/users", authMiddleware(usersHandler)).Methods("GET")
	admin.HandleFunc("/users/add", authMiddleware(usersAddHandler)).Methods("GET")
	admin.HandleFunc("/users/edit/{id}", authMiddleware(usersEditHandler)).Methods("GET")
	admin.HandleFunc("/articles/category", authMiddleware(categoryHandler)).Methods("GET")
	admin.HandleFunc("/articles/category/add", authMiddleware(categoryAddHandler)).Methods("GET")
	admin.HandleFunc("/articles/category/edit/{id}", authMiddleware(categoryEditHandler)).Methods("GET")
	admin.HandleFunc("/articles", authMiddleware(articlesHandler)).Methods("GET")
	admin.HandleFunc("/articles/add", authMiddleware(articlesAddHandler)).Methods("GET")
	admin.HandleFunc("/articles/edit/{id}", authMiddleware(articlesEditHandler)).Methods("GET")

	r.HandleFunc("/user-session", userSessionHandler).Methods("POST")
	r.PathPrefix("/login-assets/").Handler(http.StripPrefix("/login-assets/", http.FileServer(http.Dir("assets/login_template"))))
	r.PathPrefix("/backend-assets/").Handler(http.StripPrefix("/backend-assets/", http.FileServer(http.Dir("assets/backend_template"))))
	http.ListenAndServe(":3000", r)
}

func authMiddleware(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		log.Println(r.URL.Path)
		session, _ := store.Get(r, "golang-session")
		if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
			http.Redirect(w, r, app_url, http.StatusSeeOther)
		} else {
			time_now := time.Now().Local()
			dateString := session.Values["time_check"].(string)
			date_parse, error := time.Parse(YYYYMMDDhhmmss, dateString)
			midday := time.Date(date_parse.Year(), date_parse.Month(), date_parse.Day(), date_parse.Hour(), date_parse.Minute(), date_parse.Second(), 0, time.Local)

			if error != nil {
				fmt.Println(error)
				return
			}

			// fmt.Printf("Type of dateString: %T\n", dateString)
			// fmt.Printf("Type of date: %T\n", date_parse)
			// fmt.Println()
			// fmt.Printf("Value of dateString: %v\n", dateString)
			// fmt.Printf("Value of date: %v", date_parse)
			// fmt.Println()
			// fmt.Printf("Value of dateNow: %v", time_now)
			// fmt.Println()
			// fmt.Printf("Value of dateMidday: %v", midday)
			// fmt.Println()
			//time_now := time.Now().Local()
			if time_now.After(midday) {
				//log.Println("Case ---> true")
				session.Values["authenticated"] = false
				session.Save(r, w)
				http.Redirect(w, r, app_url, http.StatusSeeOther)
			}
			fmt.Println()
			log.Println("Session in authMiddleware---->", session)
			f(w, r)
		}
	}
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "golang-session")
	if auth, _ := session.Values["authenticated"].(bool); auth {
		log.Println("Session authenticated--->", session.Values["authenticated"].(bool), auth)
		http.Redirect(w, r, app_url+app_path_dashboard, http.StatusSeeOther)
	}
	t := template.Must(template.ParseFiles("views/layouts/login.html"))
	t.Execute(w, nil)
}

func userSessionHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		log.Panicln("Fatal Level:Error Method") //app exits here
		return
	}
	var userToken UserToken
	json.NewDecoder(r.Body).Decode(&userToken)
	//log.Println("User Token ------>", userToken.Token)

	session, _ := store.Get(r, "golang-session")
	jwt_value, err := jwt.Parse(userToken.Token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte("Echo31"), nil
	})
	if err != nil {
		log.Fatalf("Error jwt.Parse--->: %s", err)
	}
	if claims, ok := jwt_value.Claims.(jwt.MapClaims); ok && jwt_value.Valid {
		log.Println("Jwt Value-->", claims, claims["jti"])

		username := UserGetByUsername{
			Username: claims["jti"].(string),
		}
		//log.Println(username)
		marshalled, err := json.Marshal(username)
		if err != nil {
			log.Fatalf("impossible to marshall teacher: %s", err)
		}
		req, err := http.NewRequest("POST", api_url+api_path_users_get_by_username, bytes.NewReader(marshalled))
		if err != nil {
			log.Fatalf("impossible to build request: %s", err)
		}
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "Bearer "+userToken.Token)
		// create http client
		// do not forget to set timeout; otherwise, no timeout!
		client := http.Client{Timeout: 10 * time.Second}
		// send the request
		res, err := client.Do(req)
		if err != nil {
			log.Fatalf("impossible to send request: %s", err)
		}
		log.Printf("status Code: %d", res.StatusCode)
		defer res.Body.Close()
		// read body
		resBody, err := io.ReadAll(res.Body)
		if err != nil {
			log.Fatalf("impossible to read all body of response: %s", err)
		}
		/*Case Unmarshal*/
		dataStruct := UserDetail{}
		v := &dataStruct
		json.Unmarshal(resBody, v)
		log.Printf("res body: %s", dataStruct.Data.Name)

		// session.Values["authenticated"] = true
		// session.Values["id"] = 0
		// session.Values["name"] = nil
		// session.Values["avatar"] = nil
		// session.Values["role"] = nil
		// session.Values["status"] = nil
		//

		// Authentication goes here
		// ...
		// Set user as authenticated
		session.Values["authenticated"] = true
		session.Values["id"] = dataStruct.Data.ID
		session.Values["name"] = dataStruct.Data.Name
		session.Values["avatar"] = nil
		session.Values["role"] = dataStruct.Data.RoleID
		session.Values["status"] = dataStruct.Data.Status
		session.Values["token"] = userToken.Token
		session.Values["time_check"] = time.Now().Local().Add(time.Minute * time.Duration(14)).Format(YYYYMMDDhhmmss)

		session.Save(r, w)
		log.Println("Session Update --> ", session)
	} else {
		log.Printf("Invalid JWT Token")
		http.Redirect(w, r, app_url, http.StatusSeeOther)
	}

	//session.Values["time"] = time.Now().Local().Add(time.Minute * time.Duration(14))
	session.Save(r, w)
	w.WriteHeader(http.StatusOK)
	response := JsonData{
		Message:     "200 Ok!",
		Status_code: http.StatusOK,
	}
	json.NewEncoder(w).Encode(response)
	log.Println("User Session->", session, session.Values["authenticated"], session.Values["token"])

}
