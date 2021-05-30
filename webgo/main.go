package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"text/template"
	"time"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

var tmpl_login = template.Must((template.ParseFiles("login.html")))

type Todo struct {
	Title string
	Done  bool
}

type TodoPageData struct {
	PageTitle string
	Todos     []Todo
}

type ContactDetails struct {
	Email   string
	Subject string
	Message string
}

type UserLogin struct {
	Username string
	Password string
}

var (
	key   = "super-secret-key"
	store = sessions.NewCookieStore([]byte(key))
)

func secret(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "cookie-name")

	// Check if user is authenticated
	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	// Print secret message
	fmt.Fprintln(w, "The cake is a lie!")
}

func getlogin(wr http.ResponseWriter, r *http.Request) {
	tmpl_login.Execute(wr, nil)
}

func postlogin(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "cookie-name")

	// parse envs here
	username := os.Getenv("USERNAME")
	log.Printf(" : [DEBUG] : %s", username)
	passhash := os.Getenv("PASSWORD")

	// parse forms values here
	userUserName := r.FormValue("username")
	userPass := r.FormValue("password")

	// authentication goes here
	/// ...

	// passhash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	// if err != nil {
	// 	log.Println(" : [DEBUG] : ", err)
	// }

	log.Printf("PASS HASH :  %s", passhash)

	if bcrypt.CompareHashAndPassword([]byte(passhash), []byte(userPass)) != nil {
		log.Printf("Wrong Username : \"%s\" or Password : statucode : %d ", userUserName, http.StatusForbidden)
		loginreport := struct {
			Failed  bool
			Message string
		}{
			Failed:  true,
			Message: "Wrong Username or Password",
		}
		tmpl_login.Execute(w, loginreport)
		return
	}

	// set user as authenticated
	session.Values["authenticated"] = true
	session.Save(r, w)

	http.Redirect(w, r, "/", http.StatusFound)
}

func logout(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "cookie-name")

	// Revoke user authentication
	session.Values["authenticated"] = false
	session.Save(r, w)
}

type MiddleWare func(http.HandlerFunc) http.HandlerFunc

// logging Middleware: A middleware is a func that take http.Handlerfunc and return http.Handlerfunc
func logging() MiddleWare {
	return func(hf http.HandlerFunc) http.HandlerFunc {
		return func(rw http.ResponseWriter, r *http.Request) {

			// Do middleware things here
			start := time.Now()

			defer func() {
				log.Println(" : [DEBUG] : ", fmt.Sprintf("%s:%s/%s", r.URL.Hostname(),
					r.URL.Port(), r.URL.Path), " : ", time.Since(start), " : [ ", r.Method, " ]")
			}()

			// call the next middleware | handler in chain
			hf(rw, r)
		}
	}
}

func isAuthenticated() MiddleWare {
	return func(hf http.HandlerFunc) http.HandlerFunc {
		return func(rw http.ResponseWriter, r *http.Request) {

			//DO the middleware thing here
			session, _ := store.Get(r, "cookie-name")

			// see if we are already authenticated or not
			if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
				http.Error(rw, "You are not authenticated!", http.StatusForbidden)
				log.Println(" : [DEBUG] : ", fmt.Sprintf("%s:%s/%s", r.URL.Hostname(),
					r.URL.Port(), r.URL.Path), " : ", " : [ ", r.Method, " ] :   User Not Authenticated!")

				return
			}

			// call the next middleware | handler in chain
			hf(rw, r)
		}
	}
}

// Middleware ensures that url can only be with specified method, else returns a 400 Bad request
func Method(m string) MiddleWare {
	// create new MiddleWare
	return func(hf http.HandlerFunc) http.HandlerFunc {

		// create new handler
		return func(rw http.ResponseWriter, r *http.Request) {
			if r.Method != m {
				http.Error(rw, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
				return
			}

			// call the next middleware/handler in chain
			hf(rw, r)
		}
	}
}

func chain(f http.HandlerFunc, middlewares ...MiddleWare) http.HandlerFunc {
	for _, m := range middlewares {
		f = m(f)
	}
	return f
}

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// PARSE ENVS WHICH WE GONNA USE INSIDE MAIN SCOPE
	PORT := os.Getenv("PORT")

	tmpl := template.Must(template.ParseFiles("layout.html"))
	tmpl_forms := template.Must(template.ParseFiles("forms.html"))

	r := mux.NewRouter()

	// login handler here
	r.HandleFunc("/login", chain(getlogin, logging())).Methods("GET")
	r.HandleFunc("/login", chain(postlogin, logging())).Methods("POST")

	// secret handler here
	r.HandleFunc("/secret", chain(secret, isAuthenticated(), logging(), Method("GET")))

	// logout handler here
	r.HandleFunc("/logout", chain(logout, isAuthenticated(), logging(), Method("GET")))

	r.HandleFunc("/", chain(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Welcome to the books Home page !")
	}, logging(), Method("GET")))

	r.HandleFunc("/todo", chain(func(w http.ResponseWriter, r *http.Request) {
		data := TodoPageData{
			PageTitle: "My TODO List",
			Todos: []Todo{
				{Title: "Task 1", Done: true},
				{Title: "Task 2", Done: false},
				{Title: "Task 3", Done: true},
			},
		}

		tmpl.Execute(w, data)
	}, isAuthenticated(), logging(), Method("GET")))

	r.HandleFunc("/books/{title}/{page}", chain(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		title := vars["title"]
		page := vars["page"]

		fmt.Fprintf(w, " You 've Requested the book: %s on page %s\n", title, page)

	}, isAuthenticated(), logging(), Method("GET")))

	// handle forms here
	r.HandleFunc("/contact", chain(func(w http.ResponseWriter, r *http.Request) {

		tmpl_forms.Execute(w, nil)

	}, isAuthenticated(), logging())).Methods("GET")

	r.HandleFunc("/contact", chain(func(w http.ResponseWriter, r *http.Request) {

		details := ContactDetails{
			Email:   r.FormValue("email"),
			Subject: r.FormValue("subject"),
			Message: r.FormValue("message"),
		}

		fmt.Println("Received user concerns  oh lah lah ...")
		fmt.Printf("User Email: %s\nSubject of Concern: %s\nMessage: %s\n", details.Email, details.Subject, details.Message)

		tmpl.Execute(w, struct{ Success bool }{true})

	}, isAuthenticated(), logging())).Methods("POST")

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	if ok := http.ListenAndServe(fmt.Sprintf(":%s", PORT), r); ok != nil {
		log.Fatal(ok)
	}
}
