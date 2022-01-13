//package app
//
//import (
//	"github.com/gorilla/mux"
//	"log"
//	"net/http"
//)
//
//func Start() {
//
//	//mux := http.NewServeMux()
//
//	router := mux.NewRouter()
//
//	// define routes
//	router.HandleFunc("/api/time", getTime).Methods(http.MethodGet)
//
//	// starting server
//	log.Fatal(http.ListenAndServe("localhost:3000", router))
//
//}
