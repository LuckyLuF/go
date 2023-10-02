package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	mux := mux.NewRouter()

	//http.HandleFunc()
	mux.HandleFunc("/ciao/{name}",ciao).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe("localhost:9000",mux))
}
func ciao(w http.ResponseWriter, r *http.Request){
	vars:=mux.Vars(r)
	fmt.Fprint(w,"Ciao ",vars["name"])
}
