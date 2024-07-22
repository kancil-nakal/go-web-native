package main

import (
	"go-web-native/config"
	"go-web-native/controllers/homecontroller"
	"go-web-native/controllers/categorycontroller"
	"net/http"
	"log"
)

func main()  {
	config.ConnetDB()

	// Homepage
	http.HandleFunc("/", homecontroller.Welcome)
	
	// Categories
	http.HandleFunc("/categories", categorycontroller.Index)
	http.HandleFunc("/categories/add", categorycontroller.Add)
	http.HandleFunc("/categories/edit", categorycontroller.Edit)
	http.HandleFunc("/categories/delete", categorycontroller.Delete)

	log.Println("Server running port 8080")
	http.ListenAndServe(":8080", nil)
}