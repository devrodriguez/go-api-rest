package routes

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/devrodriguez/go-api-rest/models"

	firebase "firebase.google.com/go"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

func GetCheck(res http.ResponseWriter, req *http.Request) {
	// Firebase
	ctx := context.Background()
	sa := option.WithCredentialsFile("./firstclass-d640e-firebase-adminsdk-cne1s-97dbb8cce7.json")
	app, err := firebase.NewApp(ctx, nil, sa)
	client, err := app.Firestore(ctx)

	if err != nil {
		log.Println(err)
		err := models.Error{Description: err.Error()}
		res.Header().Set("Content-Type", "application/json")
		json.NewEncoder(res).Encode(err)
	}

	checks := client.Collection("checks")
	check := checks.Doc("04NWBpqlHMnNVHVbV4o1")
	docsnap, err := check.Get(ctx)

	if err != nil {
		errRes := models.Error{Description: err.Error()}
		json.NewEncoder(res).Encode(errRes)
	}

	//dataMap := docsnap.Data()

	// Map data
	var checkModel models.Check
	if err := docsnap.DataTo(&checkModel); err != nil {
		errRes := models.Error{Description: err.Error()}
		json.NewEncoder(res).Encode(errRes)
	}

	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(checkModel)

	defer client.Close()
}

func GetChecks(res http.ResponseWriter, req *http.Request) {
	var checksArr []models.Check

	ctx := context.Background()
	sa := option.WithCredentialsFile("./firstclass-d640e-firebase-adminsdk-cne1s-97dbb8cce7.json")
	app, err := firebase.NewApp(ctx, nil, sa)
	client, err := app.Firestore(ctx)

	if err != nil {
		log.Println(err)
		err := models.Error{Description: err.Error()}
		res.Header().Set("Content-Type", "application/json")
		json.NewEncoder(res).Encode(err)
	}

	iter := client.Collection("checks").Documents(ctx)

	for {
		var checkModel models.Check
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}

		if err != nil {
			log.Println(err)
			errRes := models.Error{Description: err.Error()}
			res.Header().Set("Content-Type", "application/json")
			json.NewEncoder(res).Encode(errRes)
		}

		err_ := doc.DataTo(&checkModel)

		if err_ != nil {
			log.Println(err_)
			errRes := models.Error{Description: err_.Error()}
			res.Header().Set("Content-Type", "application/json")
			json.NewEncoder(res).Encode(errRes)
			break
		}

		checksArr = append(checksArr, checkModel)

	}

	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(checksArr)
}
