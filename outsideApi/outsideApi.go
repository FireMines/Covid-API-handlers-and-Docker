package outsideapi

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/machinebox/graphql"
)

var GetGraphqlResponse = GetGraphqlResponsef
var GetHttpResponse = GetHttpResponsef

/**
 *	Same code as before; gets the graphql response
 *
 *	@param graphqlBody - The body of the graphql request, as a string.
 *	@param url - The url of the graphql API.
 *
 *	@return A map containing the response from the GraphQl API, and an error if something went wrong.
 */
func GetGraphqlResponsef(graphqlBody string, url string) (map[string]interface{}, error) {
	graphqlClient := graphql.NewClient(url)
	graphqlRequest := graphql.NewRequest(graphqlBody)

	var graphqlResponse map[string]interface{}
	if err := graphqlClient.Run(context.Background(), graphqlRequest, &graphqlResponse); err != nil {
		return map[string]interface{}{}, err
	}

	return graphqlResponse, nil
}

/**
 *	Gets a http respons
 *
 *	@param url - The url of the website
 *
 *	@return	A map containing the response from the API, and an error if something went wrong.
 */
func GetHttpResponsef(url string) (map[string]interface{}, error) {
	// Gets covid tracker and country from url
	resp, err := http.Get(url)
	if err != nil {
		return map[string]interface{}{}, err
	}

	defer resp.Body.Close()

	// Decoding
	dataRaw := map[string]interface{}{}
	err = json.NewDecoder(resp.Body).Decode(&dataRaw)
	if err != nil {
		return map[string]interface{}{}, err
	}
	return dataRaw, nil
}
