package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	var router = mux.NewRouter()
	router.HandleFunc("/flipit", flipIt).Methods("GET")
	
	fmt.Println("Ready to flip somebits")
	log.Fatal(http.ListenAndServe(":3000", router))
}

func flipIt(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Still alive")
}


