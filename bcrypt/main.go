package main

import (
	"context"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/go-redis/redis"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"golang.org/x/crypto/bcrypt"
)

var ctx = context.Background()
var templates *template.Template
var client *redis.Client
var store = sessions.NewCookieStore([]byte("t0p-s3cr3t"))

func main() {
	client = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	templates = template.Must(template.ParseGlob("templates/*.html"))
	r := mux.NewRouter()
	r.HandleFunc("/", indexHandler)
	r.HandleFunc("/register", registerGetHandler).Methods("GET")
	r.HandleFunc("/register", registerPostHandler).Methods("POST")
	r.HandleFunc("/login", loginGetHandler).Methods("GET")
	r.HandleFunc("/login", loginPostHandler).Methods("POST")

	fs := http.FileServer(http.Dir("./static/"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static", fs))
	fmt.Println("Hi!")
	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)

}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")
	_, ok := session.Values["username"]
	if !ok {
		http.Redirect(w, r, "/login", 302)
		return
	}
	templates.ExecuteTemplate(w, "index.html", nil)
}

func loginGetHandler(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "c500.html", nil)
}

func loginPostHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	username := r.PostForm.Get("username")
	password := r.PostForm.Get("password")
	hash, err := client.Get(ctx, "user:"+username).Bytes()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal server error!"))
		return
	}
	err = bcrypt.CompareHashAndPassword(hash, []byte(password))
	if err != nil {
		log.Fatal(err)

	}
	session, _ := store.Get(r, "session")
	session.Values["username"] = username
	fmt.Println(username)
	session.Save(r, w)
	http.Redirect(w, r, "/", 302)
}

func registerGetHandler(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "register.html", nil)
}

func registerPostHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	username := r.PostForm.Get("username")
	password := r.PostForm.Get("password")
	cost := bcrypt.DefaultCost
	hash, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	if err != nil {
		log.Fatal(err)

	}
	client.Set(ctx, "user:"+username, hash, 0)
	http.Redirect(w, r, "/login", 302)

}
