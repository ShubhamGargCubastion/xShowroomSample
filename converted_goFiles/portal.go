package main

import (
	"github.com/gorilla/mux"

	"net/http"
)

func main(){
	router:=mux.NewRouter()

           router.HandleFunc("/positions",fetch_position)
           router.HandleFunc("/users",fetch_users)
           http.ListenAndServe(":8080",router)
}

