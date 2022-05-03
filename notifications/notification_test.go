package notifications_test

import (
	"covidAss2/cases"
	consts "covidAss2/variables"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

/**
 *	Tester for the HandlerCases function.
 *	Mocks any outgoing request during testing.
 */
func TestHandlerNotifications(t *testing.T) {
	// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
	// pass 'nil' as the third parameter.
	req, err := http.NewRequest(http.MethodGet, consts.COVIDNOTIFICATIONS, nil)
	if err != nil {
		t.Error("Error when creating new request", err)
	}

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	//handler := http.HandlerFunc(notifications.NotificationHandler)

	cases.CovidInfoHandler(rr, req) // Notifications != cases??
	result := rr.Result()
	defer result.Body.Close()
	data, err := ioutil.ReadAll(result.Body)
	if err != nil {
		t.Error("Error when creating new request", err)
	}
	fmt.Println("hei", string(data))
	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	//handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `{"name":"","date":"","confirmed":0,"recovered":0,"deaths":0,"growthRate":0}`
	fmt.Println(rr.Body.String())
	if rr.Body.String() != expected {
		t.Errorf("Handler returned unexpected body: got:\n%v \nwant:\n%v", rr.Body.String(), expected)
	}

	// e.g. GET /corona/v1/notifications/
	expected = `    {
        "webhook_id": "JLF7HyWtBCU7L4jU7M4L",
        "url": "https://webhook.site/7b87be25-4e67-46e3-9f70-096ed2eaf1b6",
        "country": "Norway",
        "calls": 1
    },
    {
        "webhook_id": "OuyRmaYnAXFbpZrILY3F",
        "url": "https://webhook.site/262a030c-458b-4b8c-b8cf-5a848165ef22",
        "country": "Sweden",
        "calls": 2
    },
    {
        "webhook_id": "UfKDHNGbaDkj5qHZdwBa",
        "url": "https://webhook.site/262a030c-458b-4b8c-b8cf-5a848165ef22",
        "country": "Denmark",
        "calls": 5
    }`
	req, err = http.NewRequest(http.MethodGet, consts.COVIDNOTIFICATIONS,
		nil)
	if err != nil {
		t.Errorf("Handler returned unexpected body: got:\n%v \nwant:\n%v", rr.Body.String(), expected)
		//t.Fatal(err)
	}

	// e.g. GET /corona/v1/notifications/JLF7HyWtBCU7L4jU7M4L
	expected = `    {
			"webhook_id": "JLF7HyWtBCU7L4jU7M4L",
			"url": "https://webhook.site/7b87be25-4e67-46e3-9f70-096ed2eaf1b6",
			"country": "Norway",
			"calls": 1`
	req, err = http.NewRequest(http.MethodGet, consts.COVIDNOTIFICATIONS+"JLF7HyWtBCU7L4jU7M4L",
		nil)
	if err != nil {
		t.Errorf("Handler returned unexpected body: got:\n%v \nwant:\n%v", rr.Body.String(), expected)
		//t.Fatal(err)
	}
	// Our handler might also expect an API key.
	req.Header.Set("Authorization", "Bearer abc123")

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	//handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// ****************************** POST **************************************

	/* Values := map[string][]string{
		"url":     {"https://webhook.site/262a030c-458b-4b8c-b8cf-5a848165ef22"},
		"country": {"Denmark"},
		"calls":   {"5"},
	} */

	// e.g. POST /corona/v1/notifications/
	expected =
		`{"webhook_id": "JLF7HyWtBCU7L4jU7M4L",
		"url": "https://webhook.site/7b87be25-4e67-46e3-9f70-096ed2eaf1b6",
		"country": "Norway",
		"calls": 1}`

	req, err = http.NewRequest(http.MethodPost, consts.COVIDNOTIFICATIONS,
		strings.NewReader("Values"))
	if err != nil {
		t.Errorf("Handler returned unexpected body: got:\n%v \nwant:\n%v", rr.Body.String(), expected)
		//t.Fatal(err)
	}
	// Our handler might also expect an API key.
	//req.Header.Set("Authorization", "Bearer abc123")

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	//handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

}
