package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"vladimir-opinion/controllers"
)

func main() {
	r := mux.NewRouter()

	//////////////////////////////////////////////////////////////////////////////
	// Routing

	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets/"))))

	// Controllers
	if os.Getenv("DEVMODE") != "" {
		r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			controllers.Welcome().Index(w, r)
			return
		})
	} else {
		r.HandleFunc("/", controllers.Welcome().Index)
	}

	//////////////////////////////////////////////////////////////////////////////

	http.Handle("/", r)

	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}
