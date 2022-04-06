/* package main

import (
	"fmt"
	"context"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"

	"google.golang.org/api/option"
  )

  var opt = option.WithCredentialsFile("path/to/serviceAccountKey.json")
  var app, err = firebase.NewApp(context.Background(), nil, opt)
  if err != nil {
	return nil, fmt.Errorf("error initializing app: %v", err)
  }
*/

package main

import (
	"context" // State handling across API boundaries; part of native GoLang API
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"covidAss2"
	consts "covidAss2"
	cases "covidAss2/cases"
	handler "covidAss2/handler"
	notifications "covidAss2/notifications"
	policy "covidAss2/policy"
	status "covidAss2/status"

	"cloud.google.com/go/firestore"   // Firestore-specific support
	firebase "firebase.google.com/go" // Generic firebase support
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

// Firebase context and client used by Firestore functions throughout the program.
var ctx context.Context
var client *firestore.Client

// Collection name in Firestore
const collection = "messages"

// Message counter to produce some variation in content
var ct = 0

// Tasks:
// - Introduce update functionality via PUT and/or PATCH
// - Introduce delete functionality
// - Adapt addMessage and displayMessage function to support custom JSON schema

/*
Lists all the messages in the messages collection to the user.
*/
func displayMessage(w http.ResponseWriter, r *http.Request) {

	// Test for embedded message ID from URL
	elem := strings.Split(r.URL.Path, "/")
	messageId := elem[2]

	if len(messageId) != 0 {
		// Extract individual message

		// Retrieve specific message based on id (Firestore-generated hash)
		res := client.Collection(collection).Doc(messageId)

		// Retrieve reference to document
		doc, err2 := res.Get(ctx)
		if err2 != nil {
			http.Error(w, "Error extracting body of returned document of message "+messageId, http.StatusInternalServerError)
			return
		}

		// A message map with string keys. Each key is one field, like "text" or "timestamp"
		m := doc.Data()
		_, err3 := fmt.Fprintln(w, m["text"])
		if err3 != nil {
			http.Error(w, "Error while writing response body of message "+messageId, http.StatusInternalServerError)
			return
		}
	} else {
		// Collective retrieval of messages
		iter := client.Collection(collection).Documents(ctx) // Loop through all entries in collection "messages"

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
			_, err = fmt.Fprintln(w, m)
			if err != nil {
				http.Error(w, "Error while writing response body.", http.StatusInternalServerError)
			}
		}
	}

}

/*
Reads a string from the body in plain-text and sends it to firestore to be registered as a message.
*/
func addMessage(w http.ResponseWriter, r *http.Request) {
	// very generic way of reading body; should be customized to specific use case
	text, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Reading of payload failed", http.StatusInternalServerError)
		return
	}
	fmt.Println("Received message ", string(text))
	if len(string(text)) == 0 {
		http.Error(w, "Your message appears to be empty. Ensure to terminate URI with /.", http.StatusBadRequest)
	} else {
		// Add element in embedded structure.
		// Note: this structure is defined by the client; but exemplifying a complex one here (including Firestore timestamps).
		id, _, err := client.Collection(collection).Add(ctx,
			map[string]interface{}{
				"text": string(text),
				"ct":   ct,
				"time": firestore.ServerTimestamp,
			})
		ct++
		if err != nil {
			// Error handling
			http.Error(w, "Error when adding message "+string(text)+", Error: "+err.Error(), http.StatusBadRequest)
			return
		} else {
			fmt.Println("Entry added to collection. Identifier of returned document: " + id.ID)
			// Returns document ID in body
			http.Error(w, id.ID, http.StatusCreated)
			return
		}
	}
}

/*
Handler for all message-related operations
*/
func handleMessage(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		addMessage(w, r)
	case http.MethodGet:
		displayMessage(w, r)
	default:
		http.Error(w, "Unsupported request method", http.StatusMethodNotAllowed)
		return
	}
}

func main() {
	consts.Start = time.Now()

	// Firebase initialisation
	ctx = context.Background()

	// We use a service account, load credentials file that you downloaded from your project's settings menu.
	// It should reside in your project directory.
	// Make sure this file is git-ignored, since it is the access token to the database.
	sa := option.WithCredentialsFile("./firebase-key.json")
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		log.Fatalln(err)
	}

	// Instantiate client
	client, err = app.Firestore(ctx)

	// Alternative setup, directly through Firestore (without initial reference to Firebase); but requires Project ID
	// client, err := firestore.NewClient(ctx, projectID)

	if err != nil {
		log.Fatalln(err)
	}

	// Close down client
	defer func() {
		err := client.Close()
		if err != nil {
			log.Fatal("Closing of the firebase client failed. Error:", err)
		}
	}()

	// Make it Heroku-compatible
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	addr := ":" + port

	http.HandleFunc("/messages", handleMessage) // Be forgiving in case people for get the trailing /
	http.HandleFunc("/messages/", handleMessage)
	http.HandleFunc(covidAss2.COVIDCASES, cases.CovidInfoHandler)
	http.HandleFunc(covidAss2.COVIDPOLICY, policy.PolicyHandler)
	http.HandleFunc(covidAss2.COVIDSTATUS, status.StatusHandler)
	http.HandleFunc(covidAss2.COVIDNOTIFICATIONS, notifications.NotificationHandler)
	http.HandleFunc(covidAss2.DEFAULT_PATH, handler.EmptyHandler)

	log.Printf("Listening on %s ...\n", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		panic(err)
	}
}
