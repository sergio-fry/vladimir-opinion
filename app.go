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
	r.HandleFunc("/", controllers.Welcome().Index)

	//////////////////////////////////////////////////////////////////////////////

	http.Handle("/", r)

	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}
