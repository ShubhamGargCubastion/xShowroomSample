package main

import (
	"net/http"
	"fmt"
)

func fetch_users(w http.ResponseWriter, req *http.Request)  {
	fmt.Fprint(w,"from user file")
}