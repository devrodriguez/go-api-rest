package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"./routes"
	firebase "firebase.google.com/go"
	"github.com/gorilla/mux"
	"google.golang.org/api/option"
)

func main() {
	TestChanel()

	log.Print("Init app")
	// Create router
	router := mux.NewRouter()

	// Define endpoints
	router.HandleFunc("/tasks", routes.GetTasks).Methods("GET")
	router.HandleFunc("/task/{id}", routes.GetTask).Methods("GET")
	router.HandleFunc("/task/{id}", routes.CreateTask).Methods("POST")
	router.HandleFunc("/task/{id}", routes.DeleteTask).Methods("DELETE")
	router.HandleFunc("/check", routes.GetCheck).Methods("GET")
	router.HandleFunc("/checks", routes.GetChecks).Methods("GET")

	// Add listen port
	err := http.ListenAndServe(":3001", router)

	if err != nil {
		log.Fatal(err)
	}

	// Firebase
	log.Print("Init firebase")
	ctx := context.Background()
	sa := option.WithCredentialsFile("./firstclass-d640e-firebase-adminsdk-cne1s-97dbb8cce7.json")
	app, err := firebase.NewApp(ctx, nil, sa)
	client, err := app.Firestore(ctx)

	if err != nil {
		log.Fatalln(err)
	}

	checks := client.Collection("checks")
	check := checks.Doc("04NWBpqlHMnNVHVbV4o1")
	docsnap, _ := check.Get(ctx)

	dataMap := docsnap.Data()

	log.Print(dataMap)

	defer client.Close()

	log.Print("App run..")

}

func TestChanel() {
	chanEl := make(chan int)
	go func() {
		chanEl <- 1
	}()

	fmt.Println(<-chanEl)
}
