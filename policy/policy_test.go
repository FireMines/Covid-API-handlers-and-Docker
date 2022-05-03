package policy_test

import (
	outsideapi "covidAss2/outsideApi"
	"covidAss2/policy"
	consts "covidAss2/variables"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

/**
 *	Tester for the HandlerCases function.
 *	Mocks any outgoing request during testing.
 */
func TestHandlerPolicy(t *testing.T) {
	// Mock outgoing functions so they return hardcoded results
	getHttpResponse_original := outsideapi.GetHttpResponsef
	outsideapi.GetHttpResponse = func(url string) (map[string]interface{}, error) {
		httpResponseMocked, err := outsideapi.ReadJSONToken("httpTestData.json")
		if err != nil {
			return map[string]interface{}{}, err
		}

		return httpResponseMocked, nil
	}

	testurl := consts.COVIDPOLICY + "Norway?scope=2021-09-17"

	// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
	// pass 'nil' as the third parameter.
	req, err := http.NewRequest(http.MethodGet, testurl, nil)
	if err != nil {
		t.Error("Error when creating new request", err)
	}

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	policy.PolicyHandler(rr, req)
	result := rr.Result()
	defer result.Body.Close()

	// Format it into data
	data, err := ioutil.ReadAll(result.Body)
	if err != nil {
		t.Error("Error when creating new request", err)
	}

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := `{"date_value":"2021-09-17","country_code":"NOR","confirmed":181195,"deaths":841,"stringency_actual":38.89,"stringency":38.89,"policy":20}`

	// Check if expected result was the same as the actualy response
	data_str := string(data)
	data_str = data_str[:len(data_str)-2]
	if data_str != expected {
		t.Errorf("Handler did not return expected value: got '%v' want '%v'", data_str, expected)
	}

	// Revert mocking
	outsideapi.GetHttpResponse = getHttpResponse_original

	/* 	// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
		// pass 'nil' as the third parameter.
		req, err := http.NewRequest(http.MethodGet, consts.COVIDPOLICY, nil)
		if err != nil {
			t.Error("Error when creating new request", err)
		}

		// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(policy.PolicyHandler)

		// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
		// directly and pass in our Request and ResponseRecorder.
		handler.ServeHTTP(rr, req)

		// Check the status code is what we expect.
		if status := rr.Code; status != http.StatusOK {
			t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
		}

		/* 	// Check the response body is what we expect.
		   	expected := "Looks like you forgot to add a country! Place do so next time ;)"
		   	if rr.Body.String() != expected {
		   		t.Errorf("Handler returned unexpected body: got:\n%v \nwant:\n%v", rr.Body.String(), expected)
		   	}

		// e.g. GET /corona/v1/policy/nor?scope=2021-09-17
		expected := `{    "date_value": "2021-09-17",
	    "country_code": "NOR",
	    "confirmed": 181195,
	    "deaths": 841,
	    "stringency_actual": 38.89,
	    "stringency": 38.89,
	    "policy": 20}`
		req, err = http.NewRequest(http.MethodGet, consts.COVIDPOLICY+"nor?scope=2021-09-17",
			// Note: url.Values is a map[string][]string
			nil)
		if err != nil {
			t.Errorf("Handler returned unexpected body: got:\n%v \nwant:\n%v", rr.Body.String(), expected)
			//t.Fatal(err)
		}

		// Our handler might also expect an API key.
		req.Header.Set("Authorization", "Bearer abc123")
	*/
	// Then: call handler.ServeHTTP(rr, req) like in our first example.
}
