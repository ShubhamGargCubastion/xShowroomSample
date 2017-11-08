package main

import (
	"net/http"
	"fmt"
)

func fetch_position(w http.ResponseWriter, req *http.Request){


	fmt.Fprint(w,"hello world")
}
