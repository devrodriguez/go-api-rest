package main

import (
	"context"
	"log"
	"net/http"

	"./routes"
	firebase "firebase.google.com/go"
	"github.com/gorilla/mux"
	"google.golang.org/api/option"
)

func main() {
	log.Print("Init app")
	router := mux.NewRouter()

	router.HandleFunc("/getTasks", routes.GetTasks).Methods("GET")
	router.HandleFunc("/getCheck", routes.GetCheck).Methods("GET")
	router.HandleFunc("/getChecks", routes.GetChecks).Methods("GET")

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
