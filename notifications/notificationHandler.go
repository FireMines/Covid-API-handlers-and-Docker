package notifications

import (
	"context"
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"path"
	"strconv"
	"strings"

	consts "covidAss2/variables"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

// Webhook DB
var Webhooks = []consts.WebhookRegistration{}

var ctx context.Context
var client *firestore.Client

const collection = "webhooks"

/*
Entry point handler for Location information
*/
func NotificationHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		notificationPostRequest(w, r)
	case http.MethodGet:
		notificationGetRequest(w, r)
	case http.MethodDelete:
		notificationDeleteRequest(w, r)
	default:
		http.Error(w, "Method not supported. Currently only GET and POST are supported.", http.StatusNotImplemented)
		return
	}
}

func notificationPostRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")

	// Expects incoming body in terms of WebhookRegistration struct
	webhook := consts.WebhookRegistration{}
	err := json.NewDecoder(r.Body).Decode(&webhook)
	if err != nil {
		http.Error(w, "Something went wrong: "+err.Error(), http.StatusBadRequest)
	}
	randString := RandomString(consts.GenLength)
	hash := hmac.New(sha512.New, []byte(randString))
	webhook.Weebhook_ID = hex.EncodeToString(hash.Sum(nil))

	Webhooks = append(Webhooks, webhook)

	ctx = context.Background()

	sa := option.WithCredentialsFile("./firebase-key.json")
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		log.Fatalln(err)
	}

	client, err = app.Firestore(ctx)

	id, _, err := client.Collection(collection).Add(ctx,
		map[string]interface{}{
			"weebhook_id": webhook.Weebhook_ID,
			"url":         webhook.Url,
			"country":     webhook.Country,
			"calls":       webhook.Calls,
		})
	fmt.Println("Webhook_id for POST just done: ", webhook.Weebhook_ID)
	fmt.Println("Webhook " + webhook.Url + " has been registered.")
	http.Error(w, strconv.Itoa(len(Webhooks)-1), http.StatusCreated)

	if err != nil {
		// Error handling
		http.Error(w, "Error when adding message "+webhook.Weebhook_ID+", Error: "+err.Error(), http.StatusBadRequest)
		return
	} else {
		fmt.Println("Entry added to collection. Identifier of returned document: " + id.ID)
		// Returns document ID in body
		http.Error(w, id.ID, http.StatusCreated)
		return
	}

}

func notificationGetRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")

	urlLastVal := strings.ReplaceAll(path.Base(r.URL.Path), " ", "%20")
	//fmt.Println(urlLastVal)
	if urlLastVal == "notifications" {
		err := json.NewEncoder(w).Encode(Webhooks)
		if err != nil {
			http.Error(w, "Something went wrong: "+err.Error(), http.StatusInternalServerError)
		}
		return
	}
	for i := range Webhooks {
		if Webhooks[i].Weebhook_ID == urlLastVal {
			err := json.NewEncoder(w).Encode(Webhooks[i])
			if err != nil {
				http.Error(w, "Something went wrong: "+err.Error(), http.StatusInternalServerError)
			}
			return
		}
	}

}

func notificationDeleteRequest(w http.ResponseWriter, r *http.Request) {
	urlLastVal := strings.ReplaceAll(path.Base(r.URL.Path), " ", "%20")
	fmt.Println(urlLastVal)
	r.Header.Add("content-type", "application/json")
	if urlLastVal == "notifications" {
		http.Error(w, "Looks like you forgot to add a webhook_id! Place do so next time ;)", http.StatusBadRequest)
		return
	}
	for i := range Webhooks {
		if Webhooks[i].Weebhook_ID == urlLastVal {
			Webhooks = append(Webhooks[:i], Webhooks[i+1:]...)
			return
		}
	}
}
