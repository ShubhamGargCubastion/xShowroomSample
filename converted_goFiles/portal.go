package main

import (
	"github.com/gorilla/mux"
	_"github.com/go-sql-driver/mysql"
	"net/http"
	"github.com/jinzhu/gorm"
	"log"
	"fmt"
)

var db *gorm.DB
var err error

func main(){

    connect()
    defer db.Close()
	router:=mux.NewRouter()
	router.HandleFunc("/positions",fetch_position).Methods("GET")
	router.HandleFunc("/users",fetch_users).Methods("GET")
	router.HandleFunc("/product_categories",fetch_product_categories).Methods("GET")
	http.ListenAndServe(":8080",router)

}


func connect(){
	db,err = gorm.Open("mysql","root:password@tcp(127.0.0.1:3306)/x_showroom?charset=utf8 &parseTime=True&loc=Local")

	if err!=nil{
		log.Fatal("error connecting database")
	}else {
		fmt.Println("Connected")
	}
	err = db.DB().Ping()
	if err!=nil{
		log.Fatal(err)
	}

}
