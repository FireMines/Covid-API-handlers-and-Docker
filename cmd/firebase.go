package main

import (
	// State handling across API boundaries; part of native GoLang API
	"context"
	"log"
	"net/http"
	"os"
	"time"

	cases "covidAss2/cases"
	handler "covidAss2/handler"
	notifications "covidAss2/notifications"
	policy "covidAss2/policy"
	status "covidAss2/status"
	consts "covidAss2/variables"

	"cloud.google.com/go/firestore"   // Firestore-specific support
	firebase "firebase.google.com/go" // Generic firebase support
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

// Firebase context and client used by Firestore functions throughout the program.
//var consts.Ctx context.Context

// Collection name in Firestore
const collection = "webhooks"

// Message counter to produce some variation in content

func main() {
	consts.Start = time.Now()
	consts.Ctx = context.Background()
	// We use a service account, load credentials file that you downloaded from your project's settings menu.
	// It should reside in your project directory.
	// Make sure this file is git-ignored, since it is the access token to the database.
	sa := option.WithCredentialsFile("./firebase-key.json")
	app, err := firebase.NewApp(consts.Ctx, nil, sa)
	if err != nil {
		log.Fatalln(err)
	}

	// Instantiate client
	consts.Client, err = app.Firestore(consts.Ctx)
	if err != nil {
		log.Fatalln(err)
	}

	// Alternative setup, directly through Firestore (without initial reference to Firebase); but requires Project ID
	// client, err := firestore.NewClient(consts.Ctx, projectID)

	// Collective retrieval of messages
	iter := consts.Client.Collection(collection).Documents(consts.Ctx) // Loop through all entries in collection "messages"

	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate: %v", err)
		}
		// Note: You can access the document ID using "doc.Ref.ID"

		// A message map with string keys. Each key is one field, like "text" or "timestamp"
		m := doc.Data()
		notifications.Webhooks = append(notifications.Webhooks, storeData(m, doc))
	}

	if err != nil {
		log.Fatalln(err)
	}

	// Make it Heroku-compatible
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	addr := ":" + port

	http.HandleFunc(consts.COVIDCASES, cases.CovidInfoHandler)
	http.HandleFunc(consts.COVIDPOLICY, policy.PolicyHandler)
	http.HandleFunc(consts.COVIDSTATUS, status.StatusHandler)
	http.HandleFunc(consts.COVIDNOTIFICATIONS, notifications.NotificationHandler)
	http.HandleFunc(consts.DEFAULT_PATH, handler.EmptyHandler)

	log.Printf("Listening on %s ...\n", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		panic(err)
	}
	// Close down client
	defer func() {
		err := consts.Client.Close()
		if err != nil {
			log.Fatal("Closing of the firebase client failed. Error:", err)
		}
	}()
}
func storeData(data map[string]interface{}, doc *firestore.DocumentSnapshot) consts.WebhookRegistration {
	return consts.WebhookRegistration{
		Weebhook_ID: doc.Ref.ID,
		Url:         doc.Data()["url"].(string),
		Country:     doc.Data()["country"].(string),
		Calls:       doc.Data()["calls"].(int64),
	}
}
