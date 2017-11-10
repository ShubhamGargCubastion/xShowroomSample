package main

import (
	"github.com/gorilla/mux"
	_ "database/sql"
	"net/http"
	"github.com/jinzhu/gorm"
	"log"
	"fmt"
	_ "github.com/go-sql-driver/mysql"

)

var db *gorm.DB
var err error

func main(){

 connect()
 defer db.Close()
	router:=mux.NewRouter()
	router.HandleFunc("/positions",fetch_position)
	router.HandleFunc("/product_categories",fetch_product_categories)

	router.HandleFunc("/users",fetch_users)
	http.ListenAndServe(":8080",router)

}
func connect(){
	db,err=gorm.Open("mysql","root:password@/x_showroom?charset=utf8&parseTime=True&loc=Local")

	if err!=nil {
		log.Fatal(err)
	}

	err = db.DB().Ping()
	if err != nil {
		log.Fatal(err)
	}else{
		fmt.Println("connected")
	}

	db.SingularTable(true)

}