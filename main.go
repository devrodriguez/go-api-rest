package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"firebase.google.com/go"
	"github.com/gorilla/mux"
	"google.golang.org/api/option"
	
	"github.com/devrodriguez/go-api-rest/routes"
)

func main() {
	var port = ":3001"
	
	log.Printf("Init app on port %s", port)
	// Create router
	router := mux.NewRouter()
	
	TestChanel()
	// Define endpoints
	router.HandleFunc("/tasks", routes.GetTasks).Methods("GET")
	router.HandleFunc("/task/{id}", routes.GetTask).Methods("GET")
	router.HandleFunc("/task/{id}", routes.CreateTask).Methods("POST")
	router.HandleFunc("/task/{id}", routes.DeleteTask).Methods("DELETE")
	router.HandleFunc("/check", routes.GetCheck).Methods("GET")
	router.HandleFunc("/checks", routes.GetChecks).Methods("GET")

	router.HandleFunc("/signin", routes.SignIn).Methods("GET")
	router.HandleFunc("/login", routes.Login).Methods("GET")
	router.HandleFunc("/clientes", routes.GetClientes).Methods("GET")
	router.HandleFunc("/cliente/saludo", routes.SaludoCliente).Methods("GET")
	router.HandleFunc("/cliente", routes.CreateCliente).Methods("POST")
	router.HandleFunc("/parms", routes.GetParms).Methods("GET")

	// Add listen port
	err := http.ListenAndServe(port, router)

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
