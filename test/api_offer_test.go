/*
Suger API

Testing OfferAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package openapi

import (
	"context"
	"encoding/json"
	"fmt"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_openapi_OfferAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	header := apiClient.GetConfig().DefaultHeader
	header["Authorization"] = "Key b277c95e5e92ff7a8e96e74baf6ee2fb080db3e6507977c0067791abc1f52da4220e866e2081117a1721788aa2e9dc6fe009f2a699f17a7bba23973af6954db4"

	t.Run("Test OfferAPIService CancelOffer", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var orgId string
		var offerId string

		resp, httpRes, err := apiClient.OfferAPI.CancelOffer(context.Background(), orgId, offerId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OfferAPIService CreateOffer", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var orgId string

		resp, httpRes, err := apiClient.OfferAPI.CreateOffer(context.Background(), orgId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OfferAPIService CreateOrUpdateDraftOffer", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var orgId string

		resp, httpRes, err := apiClient.OfferAPI.CreateOrUpdateDraftOffer(context.Background(), orgId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OfferAPIService DeleteOffer", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var orgId string
		var offerId string

		resp, httpRes, err := apiClient.OfferAPI.DeleteOffer(context.Background(), orgId, offerId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OfferAPIService ExtendPrivateOfferExpiryDate", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var orgId string
		var offerId string

		resp, httpRes, err := apiClient.OfferAPI.ExtendPrivateOfferExpiryDate(context.Background(), orgId, offerId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OfferAPIService GetOffer", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var orgId = "w43Vc6UfM"
		var offerId = "1Edc6L49p"
		resp, httpRes, err := apiClient.OfferAPI.GetOffer(context.Background(), orgId, offerId).Execute()

		jsonData, err := json.Marshal(resp)
		fmt.Printf(string(jsonData))
		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OfferAPIService GetOfferByExternalId", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var orgId = "w43Vc6UfM"
		var offerExternalId = "eb2b1be2-42cc-4eb4-a227-57cc8943e8ba"

		resp, httpRes, err := apiClient.OfferAPI.GetOfferByExternalId(context.Background(), orgId, offerExternalId).Execute()
		jsonData, err := json.Marshal(resp)
		fmt.Printf(string(jsonData))

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OfferAPIService GetOfferEula", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var orgId string
		var offerId string

		resp, httpRes, err := apiClient.OfferAPI.GetOfferEula(context.Background(), orgId, offerId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OfferAPIService GetOfferResellerEula", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var orgId string
		var offerId string

		resp, httpRes, err := apiClient.OfferAPI.GetOfferResellerEula(context.Background(), orgId, offerId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OfferAPIService ListOffers", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var orgId string

		resp, httpRes, err := apiClient.OfferAPI.ListOffers(context.Background(), orgId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OfferAPIService SendOfferNotifications", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var orgId string
		var offerId string

		resp, httpRes, err := apiClient.OfferAPI.SendOfferNotifications(context.Background(), orgId, offerId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OfferAPIService UpdateOfferMetaInfo", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var orgId string
		var offerId string

		resp, httpRes, err := apiClient.OfferAPI.UpdateOfferMetaInfo(context.Background(), orgId, offerId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
